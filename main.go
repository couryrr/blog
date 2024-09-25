package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
    "strconv"
	_ "github.com/mattn/go-sqlite3"
)

var handler *ArticleHandler

func init() {
    db, err := sql.Open("sqlite3", "blog.db")

    if err != nil {
        log.Fatal("database failed to open")
    }
    
    handler, err = NewArticleHandler(db)
    
    if err != nil {
        log.Fatalf("article_handler failed to create: %s", err) 
    }
}

func main(){
    http.HandleFunc("POST /article", func(w http.ResponseWriter, r *http.Request) {
        var article Article
        err := json.NewDecoder(r.Body).Decode(&article)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
        }
        
        id, err := handler.Create(&article)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        
        _, err = w.Write([]byte(strconv.Itoa(id)))
        if err != nil {
            log.Panicln("hi")
        }
    })

    http.HandleFunc("GET /article", func(w http.ResponseWriter, r *http.Request) {
        article, err := handler.GetAllArticles()

        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
        }
        
        json.NewEncoder(w).Encode(article)
    })

    http.HandleFunc("GET /article/{id}", func(w http.ResponseWriter, r *http.Request) {
        pv := r.PathValue("id")
        if pv == "" {
            http.Error(w, "no id sent", http.StatusBadRequest)
        }

        id, err := strconv.Atoi(pv)

        if err != nil {
            http.Error(w, "bad id sent", http.StatusBadRequest)
        }

        article, err := handler.FindById(id)

        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
        }
        
        json.NewEncoder(w).Encode(article)

    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}

