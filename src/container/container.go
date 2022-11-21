package container

import (
	"audit-backend/consumer"
	"audit-backend/controllers"
	audit_handler "audit-backend/controllers/audit"
	"audit-backend/producer"
	"audit-backend/repo"
	audit_repository "audit-backend/repo/audit"
	audit_service "audit-backend/services/audit"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var auditController *controllers.AuditController
var auditHandler *audit_handler.Handler
var auditService *audit_service.Service
var auditRepo *audit_repository.Repository
var eventConsumer *consumer.EventConsumer
var eventProducer *producer.EventProducer

var Database *repo.Database

func Initialize(g *gin.RouterGroup) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
	Database = repo.InitializeConnection(uri)

	// Audit Initialize
	auditRepo = audit_repository.Initialize(Database)
	auditService = audit_service.Initialize(auditRepo)
	auditHandler = audit_handler.Initialize(auditService)
	auditController = controllers.InitializeAuditController(auditHandler).ConfigureAuditController(g.Group("/audit"))

	conf := make(kafka.ConfigMap)
	conf["group.id"] = os.Getenv("KAFKA_GROUP_ID")
	conf["auto.offset.reset"] = "earliest"
	conf["bootstrap.servers"] = "localhost:9092"
	var err error
	eventConsumer, err = consumer.Initialize(auditService, &conf)
	if err != nil {
		log.Fatalf("Kafka consumer can not be initialized, Error: %s", err.Error())
	}
	eventProducer, err = producer.Initialize(&conf)
	if err != nil {
		log.Fatalf("Example Kafka producer can not be initialized, Error: %s", err.Error())
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go eventProducer.Produce(sig)
	go eventConsumer.Listen(sig)
}
