package handlers

import (
	"github.com/labstack/echo/v4"
	"simple/backend/services"
)

func Routes(e *echo.Echo) {

	e.GET("/", services.GetLoginPage)

	e.GET("/registration", services.GetRegisterPage)
	e.POST("/registration", services.PostRegistrationPage)

	e.GET("/home", services.GetHomePage, services.IsAuthenticated)
	e.POST("/home", services.PostHomePage)

	e.POST("/check-user", services.PostCheckUser)

	e.GET("/logout", services.GetLogout)

	e.POST("/add-task", services.PostAddTask)

	e.PUT("/update-task", services.PutUpdateTask)

	e.DELETE("/delete-task", services.DeleteTask)

}
