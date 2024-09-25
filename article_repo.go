package main

import (
	"database/sql"
	"fmt"
)

type Repo interface {
    Migrate() error
    Create(article *Article) (int, error)
    Read(id int) (*Article, error)
    Update(article *Article) error
    Delete(id int) error
}

type SQLiteRepo struct {
    db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepo{
    return &SQLiteRepo{
        db: db,
    }
}

func (r *SQLiteRepo) Migrate() error {
    // Read files from directory. 
    // Apply changes.
    query := `
    CREATE TABLE IF NOT EXISTS articles(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL UNIQUE,
        slug TEXT NOT NULL UNIQUE,
        file_path TEXT NOT NULL UNIQUE,
        date_created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        date_updated DATETIME
    );
    `

    _, err := r.db.Exec(query)
    return err
}

func (r *SQLiteRepo) Create(article *Article) (int, error){ 
    query := `
        INSERT INTO articles(title, slug, file_path) VALUES(?, ?, ?);
    `

    res, err := r.db.Exec(query, article.Title, article.Slug, article.FilePath)
    if err != nil {
        return -1, err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return -1, err
    }

    return int(id), nil

}

func (r *SQLiteRepo) GetAll() (*[]Article, error) {
    query := `
        SELECT id, title, slug, date_created, date_updated
        FROM articles;
    `
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    
    defer rows.Close()
    var all []Article
    for rows.Next() {
        var article Article
        if err := rows.Scan(&article.Id, &article.Title, &article.Slug, &article.DateCreated, &article.DateUpdate); err != nil {
            return nil, err
        }
        all = append(all, article)
    }

    return &all, nil
   }
func (r *SQLiteRepo) Read(id int) (*Article, error) {
    query := `
        SELECT id, title, slug, date_created, date_updated
        FROM articles
        WHERE id = ?;
    `
    row := r.db.QueryRow(query, id)
    
    var article Article

    if err := row.Scan(&article.Id, &article.Title, &article.Slug, &article.DateCreated, &article.DateUpdate); err != nil {
        return nil, fmt.Errorf("something did not go to plan: %w", err)
    }
  
    return &article, nil
}
func (r *SQLiteRepo) Update(article *Article) error { return nil } 
func (r *SQLiteRepo) Delete(id int) error { return nil }


