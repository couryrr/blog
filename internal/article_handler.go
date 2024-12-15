package internal

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

var ( 
	mux = http.NewServeMux()
	repo Repo
)

type Repo interface {
	Create(article *Article) (int, error)
	GetAllArticles() (*[]Article, error)
	FindById(id int) (*Article, error)
}

func GetArticleHandler(articlRepo *ArticleRepo) *http.ServeMux {
	repo = articlRepo
	setup()
	return mux
}

func setup(){
	mux.HandleFunc("POST /article", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("call to POST /article")
		var article Article
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		id, err := repo.Create(&article)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		_, err = w.Write([]byte(strconv.Itoa(id)))
		if err != nil {
			slog.Error("cannot convert:", "id", id)
		}
	})

	mux.HandleFunc("GET /article", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("call to GET /article")
		article, err := repo.GetAllArticles()

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		json.NewEncoder(w).Encode(article)
	})

	mux.HandleFunc("GET /article/{id}", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("call to GET /article/{id}")
		pv := r.PathValue("id")
		if pv == "" {
			http.Error(w, "no id sent", http.StatusBadRequest)
		}

		id, err := strconv.Atoi(pv)

		if err != nil {
			http.Error(w, "bad id sent", http.StatusBadRequest)
		}

		article, err := repo.FindById(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		json.NewEncoder(w).Encode(article)

	})
}

