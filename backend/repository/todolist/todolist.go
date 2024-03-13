package todolist

import (
	"bytes"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"html/template"
	"simple/backend/config"
	"simple/backend/database/postgres"
	"time"
)

func AddTask(task, username string, db *sqlx.DB) error {

	logrus.Info("Adding new task...")

	if _, err := db.Exec("INSERT INTO todo_list (title, username, created_at) VALUES ($1, $2, $3)", task, username, time.Now()); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("New task!\n\n")

	return nil
}

func GetTasks(username string, db *sqlx.DB) ([]string, error) {

	logrus.Info("Getting tasks...")

	var tasks []string
	if err := db.Select(&tasks, "SELECT title FROM todo_list WHERE username = $1", username); err != nil {
		logrus.Error(err)
		return nil, err
	}

	logrus.Info("Tasks received!\n\n")

	return tasks, nil
}

func PutTasksInHTML(titles []string) ([]byte, error) {

	if err := postgres.DB.Select(&titles, "SELECT title FROM todo_list"); err != nil {
		logrus.Error(err)
		return nil, err
	}

	tmpl, err := template.ParseFiles("frontend/html/todo.html")
	if err != nil {
		logrus.Fatal(err)
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, titles); err != nil {
		logrus.Error(err)
		return nil, err
	}

	logrus.Info("Tasks in HTML!\n\n")

	return buf.Bytes(), nil
}

func UpdateTask(oldTask, newTask, username string, db *sqlx.DB) error {

	logrus.Info("Updating task...")

	if _, err := db.Exec("UPDATE todo_list SET title = $1, updated_at = $2 WHERE title = $3 AND username = $4", newTask, time.Now(), oldTask, username); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Task updated!\n\n")

	return nil
}

func DeleteTask(task *config.Task, db *sqlx.DB) error {

	logrus.Info("Deleting task...")

	if _, err := db.Exec("DELETE FROM todo_list WHERE title = $1", task.Task); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Task deleted!\n\n")

	return nil
}
