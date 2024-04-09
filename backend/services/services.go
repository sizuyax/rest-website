package services

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"simple/backend/logger"
	"simple/backend/models"
	"simple/backend/repository/postgres/todolist"
	"simple/backend/repository/postgres/users"
	"simple/backend/repository/redis/cookie"
	"simple/backend/services/middleware"
)

func RegisterUserIfNotExists(user *models.User) (bool, error) {

	logger.Logger.Debugf("Registration Page form values [username]: %s, [hashed password from js]: %s", user.Username, user.Password)

	isUserExists, err := users.IsUserExists(user.Username)
	if err != nil {
		logger.Logger.Error(err)
		return false, err
	}

	if isUserExists {
		logger.Logger.Debugf("[username]: %s, exists in the registration page, redirect back to registration page\n\n", user.Username)
		return true, nil
	}

	if !isUserExists {

		logger.Logger.Debug("change hashed password from js to hashed password from go..")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			logger.Logger.Error(err)
			return false, err
		}

		logger.Logger.Debug("password hash generated successfully")

		if err := users.AddUser(user.Username, string(hashedPassword)); err != nil {
			logger.Logger.Error(err)
			return false, err
		}

		logger.Logger.Debugf("new [username]: %s was added successfully to database, [new hashed password from go]: %s\n\n", user.Username, hashedPassword)
	}

	logger.Logger.Info("new user was added successfully to database\n\n")

	return false, err
}

func GenerateHTMLTasksList(user *models.User) (string, error) {

	logger.Logger.Debug("getting tasks...")

	tasks, err := todolist.GetTasks(user.Username)
	if err != nil {
		logger.Logger.Error(err)
		return "", err
	}

	tmpl, err := template.ParseFiles("frontend/html/todo.html")
	if err != nil {
		logger.Logger.Error(err)
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, tasks); err != nil {
		logger.Logger.Error(err)
		return "", err
	}

	logger.Logger.Debug("tasks received!\n\n")

	return string(buf.Bytes()), nil
}

func AuthenticateUser(user *models.User) (string, error) {

	isUserExists, err := users.IsUserExists(user.Username)
	if err != nil {
		logger.Logger.Error(err)
		return "", err
	}

	hashedPassword, err := users.GetHashedPassword(user.Username)
	if err != nil {
		logger.Logger.Error(err)
		return "", err
	}

	logger.Logger.Debugf("Home Page form values [username]: %s, [hashed password from go]: %s", user.Username, hashedPassword)

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err == nil && isUserExists == true {

		logger.Logger.Debug("username and password was compared")

		userID, err := middleware.GenerateSessionID()
		if err != nil {
			logger.Logger.Error(err)
			return "", err
		}

		if err := cookie.SetCookieSession(userID, 1); err != nil {
			logger.Logger.Error(err)
			return "", err
		}

		logger.Logger.Debug("session and cookie was created, redirect to home page\n\n")

		return userID, nil

	} else {

		logger.Logger.Debugf("incorrect username or password from [user]: %s\n\n", user.Username)

		return "Неверное имя пользователя или пароль.", nil
	}
}

func VerifyUserExistence(username, hashedPassword string) (bool, error) {

	logger.Logger.Debugf("request body post check user, [username]: %v, [hashed password from js]: %s", username, hashedPassword)

	existsUser, err := users.IsUserExists(username)
	if err != nil {
		logger.Logger.Error(err)
		return false, err
	}

	logger.Logger.Debugf("return username exists: %v\n\n", existsUser)

	return existsUser, nil
}

func DeleteCookie(username, sessionID string) error {

	logger.Logger.Debugf("deleting [cookie]: %v, from db and website for [user]: %s..\n\n", sessionID, username)

	if err := cookie.DeleteCookieFromRedis(sessionID); err != nil {
		logger.Logger.Error(err)
		return err
	}

	return nil
}

func AddTask(user *models.User, task *models.Task) error {

	logger.Logger.Debugf("adding [task]: %s, for [username]: %s..", task.Task, user.Username)

	if err := todolist.AddTask(task.Task, user.Username); err != nil {
		logger.Logger.Error(err)
		return err
	}

	logger.Logger.Debug("task was added\n\n")

	return nil
}

func UpdateTask(user *models.User, updateTask *models.TaskToUpdate) error {

	logger.Logger.Debugf("updating task [old task]: %s, [new task]: %s, for [username]: %s..", updateTask.OldTask, updateTask.NewTask, user.Username)

	if err := todolist.UpdateTask(updateTask.OldTask, updateTask.NewTask, user.Username); err != nil {
		logger.Logger.Error(err)
		return err
	}

	logger.Logger.Debug("task was updated\n\n")

	return nil
}

func DeleteTask(user *models.User, task *models.Task) error {

	logger.Logger.Debugf("deleting [task]: %s, for [username]: %s..", task.Task, user.Username)

	if err := todolist.DeleteTask(task); err != nil {
		logrus.Error(err)
		return err
	}

	logger.Logger.Debug("task was deleted\n\n")

	return nil
}
