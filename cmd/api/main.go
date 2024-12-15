package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/couryrr/blog/internal"
)

var articleRepo *internal.ArticleRepo

func init(){
	repo, err := internal.NewArticleRepo()
	if err != nil {
		slog.Error("article_handler failed to create","err", err) 
		panic(1)
	}	

	articleRepo = repo
}

func main(){
	articleMux := internal.GetArticleHandler(articleRepo)
	slog.Info("server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", articleMux))
}

