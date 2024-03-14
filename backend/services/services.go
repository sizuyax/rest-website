package services

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"simple/backend/config"
	"simple/backend/database/postgres"
	"simple/backend/database/redis"
	"simple/backend/repository/todolist"
	"simple/backend/repository/users"
	"time"
)

var (
	username   string
	task       config.Task
	updateTask config.TaskToUpdate
)

func GetLoginPage(c echo.Context) error {

	content, err := os.ReadFile("frontend/html/index.html")
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.HTML(http.StatusOK, string(content))
}

func GetRegisterPage(c echo.Context) error {

	content, err := os.ReadFile("frontend/html/registration.html")
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.HTML(http.StatusOK, string(content))
}
func PostRegistrationPage(c echo.Context) error {

	username = c.FormValue("username")
	password := c.FormValue("password")

	isUserExists, err := users.IsUserExists(postgres.DB, username)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if isUserExists {
		return c.Redirect(http.StatusMovedPermanently, "/registration")
	}

	if !isUserExists {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			logrus.Error(err)
			return err
		}

		if err = users.AddUser(postgres.DB, username, string(hashedPassword)); err != nil {
			logrus.Error(err)
			return err
		}
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func GetHomePage(c echo.Context) error {

	tasks, err := todolist.GetTasks(username, postgres.DB)
	if err != nil {
		logrus.Error(err)
		return err
	}

	htmlResponse, err := todolist.PutTasksInHTML(tasks)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.HTML(http.StatusOK, string(htmlResponse))
}
func PostHomePage(c echo.Context) error {

	username = c.FormValue("username")
	password := c.FormValue("password")

	isUserExists, err := users.IsUserExists(postgres.DB, username)
	if err != nil {
		logrus.Error(err)
		return err
	}

	hashedPassword, err := users.GetHashedPassword(postgres.DB, username)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err == nil && isUserExists == true {

		userID, err := GenerateSessionID()
		if err != nil {
			logrus.Error(err)
			return err
		}

		userID = "session: " + userID

		c.SetCookie(&http.Cookie{
			Name:    "session",
			Value:   userID,
			Expires: time.Now().Add(1 * time.Hour),
		})

		if err = StartSession(userID, 1); err != nil {
			logrus.Error(err)
			return err
		}

		return c.Redirect(http.StatusMovedPermanently, "/home")

	} else {

		alertMessage := "Неверное имя пользователя или пароль."

		return c.HTML(http.StatusOK, fmt.Sprintf(`<script>alert("%s"); window.location="/";</script>`, alertMessage))
	}
}

func PostCheckUser(c echo.Context) error {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		return err
	}

	existsUser, err := users.IsUserExists(postgres.DB, req.Username)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, config.UserExistsResponse{Exists: existsUser})
}

func GetLogout(c echo.Context) error {

	sessionID, err := c.Cookie("session")
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	if err = redis.Client.Del(sessionID.Value).Err(); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	c.SetCookie(&http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	})

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func PostAddTask(c echo.Context) error {

	if err := c.Bind(&task); err != nil {
		logrus.Error(err)
		return err
	}

	if err := todolist.AddTask(task.Task, username, postgres.DB); err != nil {
		logrus.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Task added successfully"})
}

func PutUpdateTask(c echo.Context) error {

	if err := c.Bind(&updateTask); err != nil {
		logrus.Error(err)
		return err
	}

	if err := todolist.UpdateTask(updateTask.OldTask, updateTask.NewTask, username, postgres.DB); err != nil {
		logrus.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Task updated successfully"})
}

func DeleteTask(c echo.Context) error {

	if err := c.Bind(&task); err != nil {
		logrus.Error(err)
		return err
	}

	if err := todolist.DeleteTask(&task, postgres.DB); err != nil {
		logrus.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Task deleted successfully"})
}
