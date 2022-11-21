package producer

import (
	"audit-backend/models/input"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type EventProducer struct {
	producer *kafka.Producer
}

func Initialize(conf *kafka.ConfigMap) (*EventProducer, error) {
	p, err := kafka.NewProducer(conf)
	if err != nil {
		return nil, err
	}
	return &EventProducer{p}, nil
}

func (e *EventProducer) Produce(terminate chan os.Signal) {
	select {
	case sig := <-terminate:
		fmt.Printf("Caught signal %v: terminating\n", sig)
		e.producer.Flush(20 * 1000)
		e.producer.Close()
		return
	default:
		var inputs []input.AuditInput
		file, err := os.Open("./producer/events.json")
		if err != nil {
			log.Fatalf("An error occured while reading example events file, Error: %s", err.Error())
		}
		byteValue, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalf("An error occured while reading bytes of events file, Error: %s", err.Error())
		}
		err = json.Unmarshal(byteValue, &inputs)
		if err != nil {
			log.Fatalf("An error occured while parsing example events file, Error: %s", err.Error())
		}

		go func() {
			for event := range e.producer.Events() {
				switch ev := event.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
							*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
					}
				}
			}
		}()

		topic := "events"
		for i, v := range inputs {
			singleMessage, er := json.Marshal(v)
			if er != nil {
				continue
			}
			er = e.producer.Produce(&kafka.Message{
				Key:            []byte(strconv.Itoa(i)),
				Value:          singleMessage,
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}}, nil)
			if er != nil {
				continue
			}
		}
		left := e.producer.Flush(20 * 1000)
		log.Printf("Flushed messages, still left: %d", left)
		e.producer.Close()
	}

}
