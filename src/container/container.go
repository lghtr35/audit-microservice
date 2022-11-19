package container

import (
	"audit-backend/app_config"
	"audit-backend/controllers"
	audit_handler "audit-backend/controllers/audit"
	"audit-backend/repo"
	audit_repository "audit-backend/repo/audit"
	audit_service "audit-backend/services/audit"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var auditController *controllers.AuditController
var auditHandler *audit_handler.Handler
var auditService *audit_service.Service
var auditRepo *audit_repository.Repository

var Database *repo.Database

var AppConfig *app_config.AppConfig

func Initialize(g *gin.RouterGroup) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
	Database = repo.InitializeConnection(uri)

	AppConfig = &app_config.AppConfig{DatabaseName: os.Getenv("DATABASE_NAME")}

	// Audit Initialize
	auditRepo = audit_repository.Initialize(Database)
	auditService = audit_service.Initialize(auditRepo)
	auditHandler = audit_handler.Initialize(auditService)
	auditController = controllers.InitializeAuditController(auditHandler).ConfigureAuditController(g.Group("/audit"))
}
