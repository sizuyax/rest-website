package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"simple/backend/errors"
	"simple/backend/logger"
	"simple/backend/models"
	"simple/backend/services"
	"time"
)

var (
	user       = &models.User{}
	task       = &models.Task{}
	updateTask = &models.TaskToUpdate{}
)

func GetLoginPage(c echo.Context) error {

	content, err := os.ReadFile("frontend/html/index.html")
	if err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to read HTML login page",
		})
	}

	logger.Logger.Debug("content Login Page successfully loaded\n\n")

	return c.HTML(http.StatusOK, string(content))
}

func GetRegistrationPage(c echo.Context) error {

	content, err := os.ReadFile("frontend/html/registration.html")
	if err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to read HTML registration page",
		})
	}

	logger.Logger.Debug("content Registration Page successfully loaded\n\n")

	return c.HTML(http.StatusOK, string(content))
}
func PostRegistrationPage(c echo.Context) error {

	user.Username = c.FormValue("username")

	isUserExists, err := services.RegisterUserIfNotExists(user)
	if err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to register a new user",
		})
	}

	if isUserExists {
		return c.Redirect(http.StatusMovedPermanently, "/registration")
	}

	if !isUserExists {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	return nil
}

func GetHomePage(c echo.Context) error {

	htmlResponse, err := services.GenerateHTMLTasksList(user)
	if err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to create HTML task list",
		})
	}

	return c.HTML(http.StatusOK, htmlResponse)
}
func PostHomePage(c echo.Context) error {

	user.Username = c.FormValue("username")

	userIDorIncorrectValues, err := services.AuthenticateUser(user)
	if err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to authenticate user",
		})
	}

	alertMessage := "Неверное имя пользователя или пароль."

	if userIDorIncorrectValues == alertMessage {
		return c.HTML(http.StatusOK, fmt.Sprintf(`<script>alert("%s"); window.location="/";</script>`, alertMessage))
	}

	c.SetCookie(&http.Cookie{
		Name:    "session",
		Value:   userIDorIncorrectValues,
		Expires: time.Now().Add(1 * time.Hour),
	})

	return c.Redirect(http.StatusMovedPermanently, "/home")
}

func PostCheckUser(c echo.Context) error {

	var req struct {
		Username       string `json:"username"`
		HashedPassword string `json:"hashedPassword"`
	}

	if err := c.Bind(&req); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.ErrUnmarshalFail)
	}

	user.Password = req.HashedPassword

	existsUser, err := services.VerifyUserExistence(req.Username, req.HashedPassword)
	if err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to check user existence",
		})
	}

	return c.JSON(http.StatusOK, models.UserExistsResponse{Exists: existsUser})
}

func GetLogout(c echo.Context) error {

	sessionID, err := c.Cookie("session")
	if err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to get cookie",
		})
	}

	if err := services.DeleteCookie(user.Username, sessionID.Value); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to delete cookie",
		})
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
		logger.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.ErrUnmarshalFail)
	}

	if err := services.AddTask(user, task); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to add new task",
		})
	}

	return c.JSON(http.StatusOK, errors.ErrOK)
}

func PutUpdateTask(c echo.Context) error {

	if err := c.Bind(&updateTask); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.ErrUnmarshalFail)
	}

	if err := services.UpdateTask(user, updateTask); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to update new task",
		})
	}

	return c.JSON(http.StatusOK, errors.ErrOK)
}

func DeleteTask(c echo.Context) error {

	if err := c.Bind(&task); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, errors.ErrUnmarshalFail)
	}

	if err := services.DeleteTask(user, task); err != nil {
		logger.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    "SERVER_ERROR",
			Message: "failed to delete new task",
		})
	}

	return c.JSON(http.StatusOK, errors.ErrOK)
}
