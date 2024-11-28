package routes

import (
	"github.com/elangreza14/moviefestival/cmd/http/controller"
	"github.com/elangreza14/moviefestival/cmd/http/middleware"
	"github.com/elangreza14/moviefestival/internal/domain"
	"github.com/gin-gonic/gin"
)

func GenreRoute(route *gin.RouterGroup, genreController *controller.GenreController, middleware *middleware.AuthMiddleware) {
	genreRoutes := route.Group("/genres")
	genreRoutes.GET("/popular", middleware.MustAuthMiddleware(), middleware.MustHavePermissionMiddleware(domain.AdminValPermission), genreController.MostViewedGenreList())
}
