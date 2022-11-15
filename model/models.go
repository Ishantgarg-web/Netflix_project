package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie_Name string             `json:"movie_name,omitempty"`
	Watched    bool               `json:"watched,omitempty"`
}
