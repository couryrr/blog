package internal

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type ArticleRepo struct {
	repo *sqliteRepo 
}

func NewArticleRepo() (*ArticleRepo, error){
	db, err := sql.Open("sqlite3", "blog.db")

	db.Exec("s", "test", "some")

	if err != nil {
		log.Fatal("database failed to open")
	}

	repo := &sqliteRepo{
		db: db,
	}

	err = repo.Migrate()

	if err != nil {
		return nil, err
	}

	articleRepo := &ArticleRepo{
		repo: repo, 
	}

	return articleRepo, nil

}

func (r *sqliteRepo) Migrate() error {
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

func (h *ArticleRepo) Create(article *Article) (int, error) {
	return h.repo.Create(article)
}

func (h *ArticleRepo) GetAllArticles() (*[]Article, error){
	// TODO: Add pagination
	return h.repo.GetAll()
}

func (h *ArticleRepo) FindById(id int) (*Article, error){
	return h.repo.Read(id)
}

func (r *sqliteRepo) Create(article *Article) (int, error){ 
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

func (r *sqliteRepo) GetAll() (*[]Article, error) {
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

func (r *sqliteRepo) Read(id int) (*Article, error) {
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
func (r *sqliteRepo) Update(article *Article) error { return nil } 
func (r *sqliteRepo) Delete(id int) error { return nil }

type sqliteRepo struct {
	db *sql.DB
}

func newSQLiteRepository(db *sql.DB) *sqliteRepo{
	return &sqliteRepo{
		db: db,
	}
}

