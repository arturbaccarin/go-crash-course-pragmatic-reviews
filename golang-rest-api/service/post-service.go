package service

import (
	"errors"
	"golang-rest-api/entity"
	"golang-rest-api/repository"
)

var (
	repo repository.PostRepository
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) error
	FindAll() ([]*entity.Post, error)
}

type service struct{}

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("the post is empty")
	}

	if post.Title == "" {
		return errors.New("the title is empty")
	}

	return nil
}

func (*service) Create(post *entity.Post) error {
	return repo.Save(post)
}

func (*service) FindAll() ([]*entity.Post, error) {
	return repo.FindAll()
}
