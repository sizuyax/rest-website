package models

import (
	"time"
)

type User struct {
	Username string
	Password string
}

type Task struct {
	ID        int       `json:"id" primary_key:"false"`
	Task      string    `json:"task" primary_key:"false"`
	Username  string    `json:"username" primary_key:"false"`
	CreatedAt time.Time `json:"created_at" primary_key:"false"`
	UpdatedAt time.Time `json:"updated_at" primary_key:"false"`
}

type TaskToUpdate struct {
	OldTask  string `json:"oldTask"`
	NewTask  string `json:"newTask"`
	Username string `json:"username"`
}

type UserExistsResponse struct {
	Exists bool `json:"exists"`
}
