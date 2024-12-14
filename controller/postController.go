package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rajendrakumaryadav/go_sqlproject/initializers"
	"github.com/rajendrakumaryadav/go_sqlproject/models"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	post := models.Post{
		Title:   body.Title,
		Content: body.Content,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func PostList(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{"posts": posts})
}

func DeletePost(c *gin.Context) {
	var post models.Post

	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Post not found"})
		} else {
			c.JSON(500, gin.H{"error": "Database error"})
		}
		return

	}

	initializers.DB.Delete(&post)
	c.JSON(200, gin.H{"post": post})
}

func GetPost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Post not found"})
		} else {
			c.JSON(500, gin.H{"error": "Database error"})
		}
		return
	}
	c.JSON(200, gin.H{"post": post})
}
