package repository

import (
	"database/sql"
	"golang-rest-api/entity"
	"log"
	"os"
)

type sqliteRepo struct{}

var (
	dsn string
)

func NewSQLiteRepository(connectionString string) PostRepository {
	dsn = connectionString

	os.Remove(dsn)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS posts (id integer not null primary key, title text, txt text);
	DELETE FROM posts;
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	return &sqliteRepo{}
}

func (repo *sqliteRepo) Save(post *entity.Post) error {
	// Open database connection
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close() // Ensure the database connection is closed

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback() // Rollback transaction on panic
			panic(p)
		} else if err != nil {
			tx.Rollback() // Rollback transaction on error
		} else {
			err = tx.Commit() // Commit transaction on success
		}
	}()

	// Prepare statement
	stmt, err := tx.Prepare("insert into posts(id, title, txt) values(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	// Execute statement
	_, err = stmt.Exec(post.Id, post.Title, post.Text)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo *sqliteRepo) FindAll() ([]*entity.Post, error) {
	// Open database connection
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close() // Ensure the database connection is closed

	// Query the database
	rows, err := db.Query("SELECT id, title, txt FROM posts")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close() // Ensure the rows are closed

	// Create a slice to hold the posts
	var posts []*entity.Post

	// Iterate over the rows
	for rows.Next() {
		var post entity.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Text); err != nil {
			log.Println(err)
			return nil, err
		}
		posts = append(posts, &post)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return posts, nil
}

func (repo *sqliteRepo) Delete(id int) error {
	// Open database connection
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close() // Ensure the database connection is closed

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback() // Rollback transaction on panic
			panic(p)
		} else if err != nil {
			tx.Rollback() // Rollback transaction on error
		} else {
			err = tx.Commit() // Commit transaction on success
		}
	}()

	// Prepare statement
	stmt, err := tx.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	// Execute statement
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
