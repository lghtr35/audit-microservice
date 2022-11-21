package audit

import (
	"audit-backend/models/input"
	"audit-backend/models/output"
	"audit-backend/repo/audit"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

type Service struct {
	repo *audit.Repository
}

func Initialize(r *audit.Repository) *Service {
	s := new(Service)
	s.repo = r
	return s
}

func (s *Service) getTotalRecordsForQuery(parameters map[string]interface{}) (int64, error) {
	res, err := s.repo.Collection.CountDocuments(context.Background(), parameters)
	if err != nil {
		return -1, err
	}
	return res, nil
}

func (s *Service) getMongoParametersFromQueryParameters(qs map[string][]string) (map[string]interface{}, int64, int64, error) {
	limit, err := strconv.ParseInt(qs["size"][0], 10, 64)
	if err != nil {
		return nil, 0, 0, err
	}
	page, err := strconv.ParseInt(qs["page"][0], 10, 64)
	if err != nil {
		return nil, 0, 0, err
	}
	offset := limit * page
	delete(qs, "size")
	delete(qs, "page")
	params := make(map[string]interface{})
	for key, val := range qs {
		if len(val) < 2 {
			params[key] = val[0]
		} else {
			params[key] = map[string]interface{}{val[0]: val[1]}
		}
	}

	return params, limit, offset, nil
}

func (s *Service) GetOne(id int) *output.AuditOutput {
	result := new(output.AuditOutput)
	err := s.repo.Collection.FindOne(context.Background(), bson.D{{"id", id}}).Decode(&result)
	if err != nil {
		result.Message = fmt.Sprintf("No record with given Id can not be retrieved, ID: %d\nError: %s", id, err.Error())
		result.Status = 500
		log.Printf("No record with given Id can not be retrieved, ID: %d\nError: %s", id, err.Error())
		return result
	}
	result.Status = 200
	result.Message = ""

	return result
}

func (s *Service) GetAll(queryParameters map[string][]string) *output.Page {
	result := new(output.Page)
	params, limit, offset, err := s.getMongoParametersFromQueryParameters(queryParameters)
	if err != nil {
		result.Message = fmt.Sprintf("An error occured while parsing page and size, Error: %s\n", err.Error())
		result.Status = 500
		log.Printf("An error occured while parsing page and size, Error: %s\n", err.Error())
		return result
	}
	cursor, err := s.repo.Collection.Find(context.Background(), params, options.Find().SetLimit(limit).SetSkip(offset))
	if err != nil {
		result.Message = fmt.Sprintf("An error occured while getting the records, Error: %s\n", err.Error())
		result.Status = 500
		log.Printf("An error occured while getting the records, Error: %s\n", err.Error())
		return result
	}
	cursor.All(context.Background(), &result.Content)
	result.Status = 200
	result.Message = ""
	result.PageNumber = offset / limit
	result.PageSize = limit
	result.TotalRecords, err = s.getTotalRecordsForQuery(params)
	if err != nil {
		result.Message = fmt.Sprintf("An error occured while getting the records, Error: %s\n", err.Error())
		result.Status = 500
		log.Printf("An error occured while getting the records, Error: %s\n", err.Error())
		return result
	}
	return result
}

func (s *Service) Create(input *input.AuditInput) *output.BasicOutput {
	result := new(output.BasicOutput)
	object := bson.D{
		{"event_id", input.EventId},
		{"event_status", input.EventStatus},
		{"action", input.Action},
		{"description", input.Description},
		{"variant_fields", input.VariantFields},
		{"created_at", time.Now().Unix()},
	}

	_, err := s.repo.Collection.InsertOne(context.Background(), object)
	if err != nil {
		result.Message = fmt.Sprintf("An error occured while creating the record, Error: %s\n", err.Error())
		result.Status = 500
		log.Printf("An error occured while creating the record, Error: %s\n", err.Error())
		return result
	}
	result.Message = ""
	result.Status = 201
	return result
}

func (s *Service) Delete(id int) *output.BasicOutput {
	result := new(output.BasicOutput)
	_, err := s.repo.Collection.DeleteOne(context.Background(), bson.D{{"id", id}})
	if err != nil {
		result.Message = fmt.Sprintf("An error occured while deleting the record with id %d, Error: %s\n", id, err.Error())
		result.Status = 500
		log.Printf("An error occured while deleting the record with id %d, Error: %s\n", id, err.Error())
		return result
	}
	result.Message = ""
	result.Status = 200
	return result
}
