package todolist

import (
	"simple/backend/database/postgres"
	"simple/backend/logger"
	"simple/backend/models"
	"time"
)

func AddTask(task, username string) error {

	if _, err := postgres.DB.Exec("INSERT INTO todo_list (title, username, created_at) VALUES ($1, $2, $3)", task, username, time.Now()); err != nil {
		logger.Logger.Error(err)
		return err
	}

	return nil
}

func GetTasks(username string) ([]string, error) {

	var tasks []string
	if err := postgres.DB.Select(&tasks, "SELECT title FROM todo_list WHERE username = $1", username); err != nil {
		logger.Logger.Error(err)
		return nil, err
	}

	return tasks, nil
}

func UpdateTask(oldTask, newTask, username string) error {

	if _, err := postgres.DB.Exec("UPDATE todo_list SET title = $1, updated_at = $2 WHERE title = $3 AND username = $4", newTask, time.Now(), oldTask, username); err != nil {
		logger.Logger.Error(err)
		return err
	}

	return nil
}

func DeleteTask(task *models.Task) error {

	if _, err := postgres.DB.Exec("DELETE FROM todo_list WHERE title = $1", task.Task); err != nil {
		logger.Logger.Error(err)
		return err
	}

	return nil
}
