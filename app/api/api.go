package api

import (
	"encoding/json"
	"net/http"

	"github.com/y4h2/golang-error-handling/app/entity"
	"github.com/y4h2/golang-error-handling/pkg/myerr"
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
		if userErr, ok := myerr.ToUserError(err); ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(userErr.Error()))
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
