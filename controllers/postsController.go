package controllers

import (
	"github.com/bimaagung/go-crud/initializers"
	"github.com/bimaagung/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {

	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post) 

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"message": posts,
	})
}

func GetPostById(c *gin.Context) {
	var posts []models.Post

	id := c.Param("id")

	initializers.DB.First(&posts, id)

	c.JSON(200, gin.H{
		"message": posts,
	})
}

func UpdatePost(c *gin.Context) {
	
	id := c.Param("id")
	
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	var posts []models.Post
	initializers.DB.First(&posts, id)
	initializers.DB.Model(&posts).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"message": posts,
	})
}

func DeletePost(c *gin.Context) {
	
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
