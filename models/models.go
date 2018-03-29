package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// The User struct describes a single User
	User struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpasword,omitempty"`
	}

	// The Task struct describes a single Task
	Task struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}

	// The TaskNote struct describes a single Note. A Note blongs to a Task.
	TaskNote struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskID      bson.ObjectId `json:"taskid"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
)
