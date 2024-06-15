package repository

import (
	"database/sql"
	"golang-rest-api/entity"
)

type PostRepository interface {
	Save(post *entity.Post) error
	FindAll() ([]*entity.Post, error)
}

type repo struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
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
