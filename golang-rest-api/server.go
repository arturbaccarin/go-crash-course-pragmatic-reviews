package main

import (
	"fmt"
	"golang-rest-api/controller"
	router "golang-rest-api/http"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository("./db/posts.db")
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVER(port)
}
