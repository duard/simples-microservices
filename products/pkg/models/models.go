package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product is used to represent product profile data
type Product struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"title,omitempty"`
	CreatedOn time.Time          `bson:"createdon,omitempty"`
}
