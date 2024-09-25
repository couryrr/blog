package main

import "database/sql"

type ArticleHandler struct {
   repo *SQLiteRepo 
}

func NewArticleHandler(db *sql.DB) (*ArticleHandler, error){
    repo := &SQLiteRepo{
        db: db,
    }

    err := repo.Migrate()

    if err != nil {
        return nil, err
    }

    handler := &ArticleHandler{
        repo: repo, 
    }
    
    return handler, nil

}


func (h *ArticleHandler) Create(article *Article) (int, error) {
    return h.repo.Create(article)
}

func (h *ArticleHandler) GetAllArticles() (*[]Article, error){
    // TODO: Add pagination
    return h.repo.GetAll()
}
func (h *ArticleHandler) FindById(id int) (*Article, error){
    return h.repo.Read(id)
}
