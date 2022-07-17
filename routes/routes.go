package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morelmiles/jwt-login-radius/controllers"
)

func Routes() {
	router := gin.Default()

	// Users
	router.POST("/api/v1/users", controllers.CreateUser)
	router.GET("/api/v1/users", controllers.GetUsers)
	router.GET("/api/v1/user/:id", controllers.GetUserById)
	router.DELETE("/api/v1/user/:id", controllers.DeleteUserById)
	router.PUT("/api/v1/user/:id", controllers.UpdateUserById)

	// Posts
	router.POST("/api/v1/posts", controllers.CreatePost)
	router.GET("/api/v1/posts/", controllers.GetPostById)
	router.GET("/api/v1/post/:id", controllers.GetPosts)
	router.DELETE("/api/v1/post/:id", controllers.DeletePostById)
	router.PUT("/api/v1/post/:id", controllers.UpdatePostById)

	// Server port
	router.Run(":8000")
}
