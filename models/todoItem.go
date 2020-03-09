package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// TodoItem represents an item in the todo list.
type TodoItem struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Date        time.Time
	Description string
	Done        bool
}
