package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 테스트 
	router := gin.Default()
	router.Run("localhost:8080")
}