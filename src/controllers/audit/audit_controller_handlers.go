package audit

import (
	"audit-backend/models/output"
	"audit-backend/services/audit"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *audit.Service
}

func Initialize(s *audit.Service) *Handler {
	return &Handler{s}
}

// FilterAudits godoc
// @Summary Filter audit records with pagination
// @Description Query Parameters will give the filtering abilities you can write mongo queries for comparisons
// @Tags audits
// @Accept json
// @Produce json
// @Param page query int true "Page Number"
// @Param size query int true "Page Size"
// @Success 200 {object} output.AuditPage
// @Failure 500 {object} output.BasicOutput
// @Failure 400 {object} output.BasicOutput
// @Router /audit/filter [get]
func (h Handler) FilterAudits(c *gin.Context) {
	qs := make(map[string][]string)
	for key, val := range c.Request.URL.Query() {
		qs[key] = val
	}
	if _, ok := qs["page"]; !ok {
		failure := generateErrorResponse("Page number is a mandatory parameter", 400)
		c.JSON(failure.Status, failure)
	}
	if _, ok := qs["size"]; !ok {
		failure := generateErrorResponse("Page size is a mandatory parameter", 400)
		c.JSON(failure.Status, failure)
	}
	res := h.service.GetAll(qs)
	c.JSON(res.Status, res)
}

func generateErrorResponse(msg string, status int) *output.BasicOutput {
	failure := new(output.BasicOutput)
	failure.Status = status
	failure.Message = msg
	return failure
}
