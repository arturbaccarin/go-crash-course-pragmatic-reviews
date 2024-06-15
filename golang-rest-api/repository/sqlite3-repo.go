package repository

import (
	"database/sql"
	"golang-rest-api/entity"

	_ "github.com/mattn/go-sqlite3"
)

type repo struct {
	db *sql.DB
}

func NewSQLite3Repository() PostRepository {
	var err error
	db, err := sql.Open("sqlite3", "./db/postDB.db")
	if err != nil {
		panic("Failed to open the database: " + err.Error())
	}

	if err = db.Ping(); err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	return &repo{
		db: db,
	}
}

func (r *repo) Save(post *entity.Post) error {
	stmt, err := r.db.Prepare("INSERT INTO post (title, text) VALUES ($1, $2);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Text)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) FindAll() ([]*entity.Post, error) {
	stmt, err := r.db.Prepare("SELECT id, title, text FROM post;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*entity.Post
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Text)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
