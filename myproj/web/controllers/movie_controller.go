package controllers

import (
	"app/repositories"
	"app/service"
	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View {
	movieRepo := repositories.NewMovieManager()
	movieService := service.NewMoviceServiceManager(movieRepo)
	movieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieResult,
	}
}