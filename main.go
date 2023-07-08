package main

import (
	signup "to-do-list/api/user"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 테스트 
	router := gin.Default()
	router.POST("/user/signup", signup.Signup)

	router.Run("localhost:8080")

}