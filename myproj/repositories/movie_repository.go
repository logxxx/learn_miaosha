package repositories

import "app/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	movie := &datamodels.Movie{Name:"木渴望视频"}
	return movie.Name
}
