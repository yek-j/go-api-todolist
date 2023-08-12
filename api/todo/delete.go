package todo

import (
	"fmt"
	"net/http"
	connector "to-do-list/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteTodo(c *gin.Context) {
	todoId := c.Param("todoid")
	todoUUID, err := uuid.Parse(todoId);
	fmt.Print(todoId)
	if todoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "todoid emtpy",
			"message": "delete Todo todoid empty",
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