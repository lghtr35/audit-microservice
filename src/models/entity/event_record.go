package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type EventRecord struct {
	Id            primitive.ObjectID `bson:"id,omitempty"`
	Status        *bool              `bson:"status,omitempty"`
	EventId       *int64             `bson:"event_id,omitempty"`
	Action        string             `bson:"action,omitempty"`
	Description   string             `bson:"description,omitempty"`
	VariantFields primitive.M        `bson:"variant_fields,omitempty"`
}
