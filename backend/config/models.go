package config

import "time"

type Config struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresName     string `env:"POSTGRES_DB"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPorts    string `env:"POSTGRES_PORTS"`
	RedisHost        string `env:"REDIS_HOST"`
	RedisPorts       string `env:"REDIS_PORTS"`
	Reload           bool   `env:"RELOAD"`
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
