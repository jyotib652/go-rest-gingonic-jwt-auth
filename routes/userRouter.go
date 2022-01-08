package routes

import (
	controller "github.com/jyotib652/go-rest-gingonic-jwt/controllers"

	"github.com/jyotib652/go-rest-gingonic-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}
