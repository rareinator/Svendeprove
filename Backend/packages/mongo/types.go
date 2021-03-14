package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	SensorID int32              `bson:"sensorID,omitempty"`
	Data     string             `bson:"data,omitempty"`
	Date     time.Time          `bson:"date,omitempty"`
}
