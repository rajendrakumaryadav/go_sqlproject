package main

import (
	"github.com/rajendrakumaryadav/go_sqlproject/initializers"
	"github.com/rajendrakumaryadav/go_sqlproject/models"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
