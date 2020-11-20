package service

import (
	"errors"

	"github.com/y4h2/golang-error-handling/app/entity"
	"github.com/y4h2/golang-error-handling/app/repository"
)

var NotFoundErr = errors.New("article not found")

type RepositoryLayer interface {
	GetArticleByID(id int64) (*entity.Article, error)
}

type Service struct {
	repository RepositoryLayer
}

func New(repository RepositoryLayer) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetArticle(id int64) (*entity.Article, error) {
	article, err := s.repository.GetArticleByID(id)
	if err != nil {
		if err == repository.NotFoundErr {
			return nil, NotFoundErr
		}
		return nil, err
	}

	return article, nil
}
