package service

import (
	"golang-rest-api/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) FindAll() ([]*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]*entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "the post is empty")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Id: 1, Title: "", Text: "B"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "the title is empty")
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int = 1

	post := entity.Post{Id: identifier, Title: "A", Text: "B"}

	// Setup Expectations
	mockRepo.On("FindAll").Return([]*entity.Post{&post}, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, identifier, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	post := entity.Post{Id: 0, Title: "A", Text: "B"}

	mockRepo.On("Save").Return(nil)

	testService := NewPostService(mockRepo)

	err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
}
