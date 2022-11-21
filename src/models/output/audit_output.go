package output

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuditOutput struct {
	Id            primitive.ObjectID `json:"id"`
	EventStatus   *bool              `json:"event_status"`
	EventId       *int64             `json:"event_id"`
	Action        string             `json:"action"`
	Description   string             `json:"description"`
	VariantFields primitive.M        `json:"variant_fields"`
	BasicOutput
}
