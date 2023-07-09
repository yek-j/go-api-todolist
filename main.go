package main

import (
	signup "to-do-list/api/user"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger()) 
	router.Use(gin.Recovery())

	router.POST("/user/signup", signup.Signup)

	router.Run("localhost:8080")

}