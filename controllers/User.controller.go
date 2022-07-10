package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morelmiles/jwt-login-radius/config"
	"github.com/morelmiles/jwt-login-radius/models"
)

// CreateUser creates
func CreateUser(c *gin.Context) {

	var user models.User

	db := config.GetDB()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{Name: user.Name, Username: user.Username,
		Password: user.Password, Email: user.Email}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newUser)

}

func GetUserById(c *gin.Context) {
	var user models.User

	db := config.GetDB()

	if err := db.Where("id= ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {

	var users []models.User

	db := config.GetDB()

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

func DeleteUserById(c *gin.Context) {

	var user models.User

	db := config.GetDB()

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found!"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})

}
