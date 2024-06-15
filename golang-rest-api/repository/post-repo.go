package repository

import (
	"golang-rest-api/entity"
)

type PostRepository interface {
	Save(post *entity.Post) error
	FindAll() ([]*entity.Post, error)
	Delete(id int) error
}
