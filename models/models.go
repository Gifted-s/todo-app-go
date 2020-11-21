package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List struct {
	ID  primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`;
	Name string `json:"name,omitempty" bson: "name,omitempty"`;
	DateCreated string  `json:"date_created,omitempty" bson: "date_created,omitempty"`;
	Tasks []Task	`json:"tasks,omitempty" bson: "tasks,omitempty"`;
}

type Task struct {
	ID  primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`;
	Name string `json:"name,omitempty" bson: "name,omitempty"`;
	DateCreated string  `json:"date_created,omitempty" bson: "date_created,omitempty"`;
	Completed bool `json:"completed" bson: "completed"`;
}