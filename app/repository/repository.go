package repository

import (
	"database/sql"
	"errors"

	"github.com/y4h2/golang-error-handling/app/entity"
)

var NotFoundErr = errors.New("article not found")

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetArticleByID(id int64) (*entity.Article, error) {
	query := `SELECT id,title,content, author, updated_at, created_at
  						FROM article WHERE id = ?`

	row := repo.db.QueryRow(query, id)

	article := &entity.Article{}
	err := row.Scan(article.ID, article.Title, article.Author, article.UpdatedAt, article.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFoundErr
		}

		return nil, err
	}

	return article, nil
}
