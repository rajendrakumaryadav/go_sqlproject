package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rajendrakumaryadav/go_sqlproject/controller"
	"github.com/rajendrakumaryadav/go_sqlproject/initializers"
	"log"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {
	// Initialize gin Application
	r := gin.Default()

	// Adding Request handler
	r.POST("/post", controller.CreatePost)
	r.GET("/posts", controller.PostList)
	r.DELETE("/post/:id", controller.DeletePost)
	r.GET("/post/:id", controller.GetPost)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
