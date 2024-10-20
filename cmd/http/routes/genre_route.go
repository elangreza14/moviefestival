package routes

import (
	"github.com/elangreza14/moviefestival/controller"
	"github.com/elangreza14/moviefestival/middleware"
	"github.com/elangreza14/moviefestival/model"
	"github.com/gin-gonic/gin"
)

func GenreRoute(route *gin.RouterGroup, genreController *controller.GenreController, middleware *middleware.AuthMiddleware) {
	genreRoutes := route.Group("/genres")
	genreRoutes.GET("/popular", middleware.MustAuthMiddleware(), middleware.MustHavePermissionMiddleware(model.AdminValPermission), genreController.MostViewedGenreList())
}
