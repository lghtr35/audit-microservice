package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventRecord struct {
	Id            primitive.ObjectID `bson:"id,omitempty"`
	EventStatus   *bool              `bson:"event_status,omitempty"`
	EventId       *int64             `bson:"event_id,omitempty"`
	Action        string             `bson:"action,omitempty"`
	Description   string             `bson:"description,omitempty"`
	VariantFields primitive.M        `bson:"variant_fields,omitempty"`
	CreatedAt     *int64             `bson:"created_at"`
}
