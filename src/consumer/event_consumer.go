package consumer

import (
	"audit-backend/models/input"
	audit_service "audit-backend/services/audit"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"sync"
	"time"
)

type EventConsumer struct {
	service  *audit_service.Service
	consumer *kafka.Consumer
}

// TODO termination of consumer has problems ...

func Initialize(s *audit_service.Service, conf *kafka.ConfigMap, topic string) (*EventConsumer, error) {
	c, err := kafka.NewConsumer(conf)
	if err != nil {
		return nil, err
	}
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("Consumer Initialized\n")
	return &EventConsumer{s, c}, nil
}

func (e *EventConsumer) Listen(terminate chan os.Signal, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	log.Printf("Consumer Started to listen\n")
	ok := true
	defer wg.Done()
	for ok {
		select {
		case sig := <-terminate:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			ok = false
		default:
			ev, err := e.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}
			var input input.AuditInput
			err = json.Unmarshal(ev.Value, &input)
			if err != nil {
				log.Printf("An error occured while parsing the kafka message, Error: %s\n", err.Error())
				continue
			}
			log.Printf("Event received with %d event id and %s sequence id. Event action is %s\n", input.EventId, string(ev.Key), input.Action)
			e.service.Create(&input)
		}
	}
	log.Println("Out of listen")
}

func (e *EventConsumer) Close() {
	err := e.consumer.Close()
	if err != nil {
		log.Println(fmt.Sprintf("An error %v occurred while closing kafka consumer.", err))
	}
}
