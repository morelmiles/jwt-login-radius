package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morelmiles/jwt-login-radius/controllers"
)

func Routes() {
	router := gin.Default()

	// Users
	router.POST("/api/v1/user", controllers.CreateUser)
	router.GET("/api/v1/users", controllers.GetUsers)
	router.GET("/api/v1/user/:id", controllers.GetUserById)
	router.DELETE("/api/v1/user/:id", controllers.DeleteUserById)

	// Posts
	router.POST("/api/v1/users", controllers.CreatePost)
	router.GET("/api/v1/user/:id", controllers.GetPostById)
	router.GET("/api/v1/user/:id", controllers.GetPosts)
	router.DELETE("/api/v1/user/:id", controllers.DeletePostById)

	router.Run(":8000")
}
