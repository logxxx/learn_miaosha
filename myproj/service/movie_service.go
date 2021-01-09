package service

import (
	"app/repositories"
	"fmt"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMoviceServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	return fmt.Sprintf("我们获取到的视频名称为:%v", m.repo.GetMovieName())
}