package todo

import (
	"net/http"
	"to-do-list/api/common"
	connector "to-do-list/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteTodo(c *gin.Context) {
	
	// token 확인
	token := c.GetHeader("Authorization")
	verifiedClaims, err :=common.ValidateToken(token)

	if verifiedClaims == nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"add": "faile",
			"error": err,
			"message": "로그인이 필요합니다.",
		})	
		return 
	}

	todoId := c.Param("todoid")
	todoUUID, err := uuid.Parse(todoId);


	if todoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "todoid emtpy",
			"message": "deleteTodo todoid empty",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "deleteTodo uuid error",
		})
		return
	}

	db, ctx := connector.Connector()
	err = db.Todo.
			DeleteOneID(todoUUID).
			Exec(ctx);

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "deleteTodo delete error",
		})
		return
	}
	
	defer db.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}