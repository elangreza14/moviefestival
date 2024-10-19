package routes

import (
	"github.com/elangreza14/moviefestival/controller"
	"github.com/gin-gonic/gin"
)

func MovieRoute(route *gin.RouterGroup, movieController *controller.MovieController) {
	movieRoutes := route.Group("/movies")
	movieRoutes.GET("", movieController.MovieList())
}
