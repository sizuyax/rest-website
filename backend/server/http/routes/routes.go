package routes

import (
	"github.com/labstack/echo/v4"
	"simple/backend/server/http/handlers"
	"simple/backend/services/middleware"
)

func SetupRoutes(e *echo.Echo) {

	e.GET("/", handlers.GetLoginPage)

	e.GET("/registration", handlers.GetRegistrationPage)
	e.POST("/registration", handlers.PostRegistrationPage)

	e.GET("/home", handlers.GetHomePage, middleware.IsAuthenticated)
	e.POST("/home", handlers.PostHomePage)

	e.POST("/check-user", handlers.PostCheckUser)

	e.POST("/logout", handlers.GetLogout)

	e.POST("/add-task", handlers.PostAddTask, middleware.IsAuthenticated)

	e.PUT("/update-task", handlers.PutUpdateTask, middleware.IsAuthenticated)

	e.DELETE("/delete-task", handlers.DeleteTask, middleware.IsAuthenticated)
}
