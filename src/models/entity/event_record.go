package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventRecord struct {
	Id            primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	EventStatus   *bool                  `bson:"event_status,omitempty" json:"event_status"`
	EventId       *int64                 `bson:"event_id,omitempty" json:"event_id"`
	Action        string                 `bson:"action,omitempty" json:"action"`
	Description   string                 `bson:"description,omitempty" json:"description"`
	VariantFields map[string]interface{} `bson:"variant_fields,omitempty" json:"variant_fields"`
	CreatedAt     *int64                 `bson:"created_at" json:"created_at"`
}
