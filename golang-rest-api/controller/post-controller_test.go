package controller

import (
	"bytes"
	"encoding/json"
	"golang-rest-api/entity"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository("../db/postsTest.db")
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postController PostController            = NewPostController(postSrv)
)

func TestAddPost(t *testing.T) {
	// Create a new HTTP POST request
	var jsonReq = []byte(`{"id": 1, "title": "title 4", "text": "text 4"}`)
	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Error(err.Error())
		return
	}

	// Assign HTTP Handler function (controller AddPost function)
	handler := http.HandlerFunc(postController.AddPosts)

	// Record HTTP Reponse (httptest)
	response := httptest.NewRecorder()

	// Disátch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP Status Code and the response
	status := response.Code

	if status != http.StatusCreated {
		t.Errorf("handler returned a wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Clean Up database
	cleanUp(1)

	/*
		// Decode the HTTP Response
		var post entity.Post
		json.NewDecoder(io.Reader(response.Body)).Decode(&post)

		// Assert HTTP response
		assert.NotNil(t, post.Id)
		assert.Equal(t, TITLE, post.Title)
		assert.Equal(t, TEXT, post.Text)
	*/
}

func TestGetPosts(t *testing.T) {
	setup()

	req, _ := http.NewRequest("GET", "/posts", nil)

	handler := http.HandlerFunc(postController.GetPosts)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("handler returned a worng status code: got %v want %v", status, http.StatusOK)
	}

	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	assert.NotNil(t, posts[0].Id)
	assert.Equal(t, "title", posts[0].Title)
	assert.Equal(t, "text", posts[0].Text)

	cleanUp(posts[0].Id)
}

func setup() {
	var post entity.Post = entity.Post{
		Id:    15,
		Title: "title",
		Text:  "text",
	}

	postRepo.Save(&post)
}

func cleanUp(id int) {
	postRepo.Delete(id)
}
