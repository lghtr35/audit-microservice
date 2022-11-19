package audit

import (
	"audit-backend/services/audit"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *audit.Service
}

func Initialize(s *audit.Service) *Handler {
	return &Handler{s}
}

func (h Handler) FilterAudits(c *gin.Context) {

}

func (h Handler) AddEvent(c *gin.Context) {

}
