package main

import (
	user "to-do-list/api/user"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger()) 
	router.Use(gin.Recovery())

	router.POST("/user/signup", user.Signup)
	router.POST("/user/emailcheck", user.Emailcheck)
	router.POST("/user/signin", user.Signin)

	router.Run("localhost:8080")

}