package audit

import (
	"audit-backend/models/entity"
	"audit-backend/models/output"
	"audit-backend/repo/audit"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Service struct {
	repo *audit.Repository
}

func Initialize(r *audit.Repository) *Service {
	s := new(Service)
	s.repo = r
	return s
}

func (s *Service) GetOne(id int) *entity.EventRecord {
	var result entity.EventRecord
	err := s.repo.Collection.FindOne(context.Background(), bson.D{{"id", id}}).Decode(&result)
	if err != nil {
		log.Panicf("No record with given Id has been found, ID: %s\n", id)
		return nil
	}

	return &result
}

func (s *Service) GetAll(queryParameters map[string]interface{}) *output.Page {
	return new(output.Page)
}
