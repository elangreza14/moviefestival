package routes

import (
	"github.com/elangreza14/moviefestival/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoute(route *gin.RouterGroup, AuthController *controller.AuthController) {
	AuthRoutes := route.Group("/auth")
	AuthRoutes.POST("/register", AuthController.RegisterUser())
	AuthRoutes.POST("/login", AuthController.LoginUser())
}
