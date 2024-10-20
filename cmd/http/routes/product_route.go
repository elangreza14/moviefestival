package routes

import (
	"net/http"

	"github.com/elangreza14/moviefestival/controller"
	"github.com/elangreza14/moviefestival/middleware"
	"github.com/elangreza14/moviefestival/model"
	"github.com/gin-gonic/gin"
)

func MovieRoute(route *gin.RouterGroup, movieController *controller.MovieController, middleware *middleware.AuthMiddleware) {
	movieRoutes := route.Group("/movies")
	movieRoutes.GET("", movieController.MovieList())
	movieRoutes.StaticFS("/public", http.Dir("public"))
	movieRoutes.POST("/upload", middleware.MustAuthMiddleware(), middleware.MustHavePermissionMiddleware(model.AdminValPermission), movieController.UploadMovie())
	movieRoutes.POST("", middleware.MustAuthMiddleware(), middleware.MustHavePermissionMiddleware(model.AdminValPermission), movieController.CreateMovie())
}
