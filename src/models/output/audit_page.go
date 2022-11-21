package output

import "audit-backend/models/entity"

type AuditPage struct {
	Content []entity.EventRecord `json:"content"`
	Page
}
