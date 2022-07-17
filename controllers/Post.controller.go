package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morelmiles/jwt-login-radius/config"
	"github.com/morelmiles/jwt-login-radius/models"
)

// CreateUser creates
func CreatePost(c *gin.Context) {

	var post models.Post

	db := config.GetDB()

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost := models.Post{Title: post.Title, Content: post.Content, CoverImage: post.CoverImage, Author: post.Author}

	if err := db.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newPost)

}

func GetPostById(c *gin.Context) {
	var post models.Post

	db := config.GetDB()

	if err := db.Where("id= ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func GetPosts(c *gin.Context) {

	var posts []models.Post

	db := config.GetDB()

	if err := db.Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)

}

func DeletePostById(c *gin.Context) {

	var post models.Post

	db := config.GetDB()

	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found!"})
		return
	}

	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})

}

func UpdatePostById(c *gin.Context) {
	var post models.Post
	db := config.GetDB()

	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found!"})
		return
	}
	c.BindJSON(&post)

	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}
