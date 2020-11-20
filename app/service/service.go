package service

import (
	"github.com/pkg/errors"
	"github.com/y4h2/golang-error-handling/app/entity"
)

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
		return nil, errors.Wrap(err, "failed to GetArticle")
	}

	return article, nil
}
