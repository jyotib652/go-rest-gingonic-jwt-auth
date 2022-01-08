package routes

import (
	controller "github.com/jyotib652/go-rest-gingonic-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("user/login", controller.Login())

}
