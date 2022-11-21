package consumer

import (
	"audit-backend/models/input"
	audit_service "audit-backend/services/audit"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

type EventConsumer struct {
	service  *audit_service.Service
	consumer *kafka.Consumer
}

func Initialize(s *audit_service.Service, conf *kafka.ConfigMap) (*EventConsumer, error) {
	c, err := kafka.NewConsumer(conf)
	if err != nil {
		return nil, err
	}
	err = c.SubscribeTopics([]string{"events"}, nil)
	if err != nil {
		return nil, err
	}
	return &EventConsumer{s, c}, nil
}

func (e *EventConsumer) Listen(terminate chan os.Signal) {
	ok := true
	for ok {
		select {
		case sig := <-terminate:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			ok = false
			break
		default:
			ev, err := e.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			var input input.AuditInput
			err = json.Unmarshal(ev.Value, &input)
			if err != nil {
				log.Printf("An error occured while parsing the kafka message, Error: %s\n", err.Error())
				continue
			}
			e.service.Create(&input)
		}
	}
	err := e.consumer.Close()
	if err != nil {
		return
	}
}
