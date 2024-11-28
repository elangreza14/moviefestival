package routes

import (
	"net/http"

	"github.com/elangreza14/moviefestival/cmd/http/controller"
	"github.com/elangreza14/moviefestival/cmd/http/middleware"
	"github.com/elangreza14/moviefestival/internal/domain"
	"github.com/gin-gonic/gin"
)

func MovieRoute(route *gin.RouterGroup, movieController *controller.MovieController, middleware *middleware.AuthMiddleware) {
	movieRoutes := route.Group("/movies")
	movieRoutes.GET("", movieController.MovieList())
	movieRoutes.StaticFS("/public", http.Dir("public"))
	movieRoutes.POST("/upload", middleware.MustAuthMiddleware(), middleware.MustHavePermissionMiddleware(domain.AdminValPermission), movieController.UploadMovie())
	movieRoutes.POST("", middleware.MustAuthMiddleware(), middleware.MustHavePermissionMiddleware(domain.AdminValPermission), movieController.CreateMovie())
	movieRoutes.PUT("/:movieID", middleware.MustAuthMiddleware(), middleware.MustHavePermissionMiddleware(domain.AdminValPermission), movieController.UpdateMovie())
	movieRoutes.GET("/:movieID", movieController.GetMovieDetail())
	movieRoutes.GET("/popular", middleware.MustAuthMiddleware(), middleware.MustHavePermissionMiddleware(domain.AdminValPermission), movieController.MostViewedMovieList())
}
