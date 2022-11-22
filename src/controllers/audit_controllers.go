package controllers

import (
	"audit-backend/controllers/audit"
	"github.com/gin-gonic/gin"
)

type AuditController struct {
	handler *audit.Handler
}

func InitializeAuditController(h *audit.Handler) *AuditController {
	return &AuditController{h}
}

func (c AuditController) ConfigureAuditController(group *gin.RouterGroup) *AuditController {
	group.GET("/filter", c.handler.FilterAudits)
	return &c
}
