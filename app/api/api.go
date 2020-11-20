package api

import (
	"encoding/json"
	"net/http"

	"github.com/y4h2/golang-error-handling/app/entity"
	"github.com/y4h2/golang-error-handling/app/service"
)

type ServiceLayer interface {
	GetArticle(id int64) (*entity.Article, error)
}

type API struct {
	service ServiceLayer
}

func New(service ServiceLayer) *API {
	return &API{
		service: service,
	}
}

func (api *API) Get(w http.ResponseWriter, r *http.Request) {
	id := int64(1)

	article, err := api.service.GetArticle(id)
	if err != nil {
		if err == service.NotFoundErr {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("article not found"))
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
	}

	b, err := json.Marshal(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
