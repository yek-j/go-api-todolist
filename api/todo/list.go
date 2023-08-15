package todo

import (
	"net/http"
	"to-do-list/api/common"
	connector "to-do-list/db"
	"to-do-list/ent/todo"
	"to-do-list/ent/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TodoList(c *gin.Context) {
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

	userID, err := uuid.Parse(verifiedClaims.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "TodoList uuid error",
		})
		return
	}

	db, ctx := connector.Connector()
	todos, err := db.Todo.
		Query().
		Where(todo.HasUserWith(user.ID(userID))).
		All(ctx)
	
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "TodoList db error",
		})
		return
	}

	defer db.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todos" : todos,
	})
}