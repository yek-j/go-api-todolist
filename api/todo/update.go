package todo

import (
	"net/http"
	"strconv"
	connector "to-do-list/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateTodoCheck(c *gin.Context) {
	todoId := c.Param("todoid")
	check, _ := strconv.ParseBool(c.PostForm("check"))

	todoUUID, err := uuid.Parse(todoId);

	if todoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "todoid emtpy",
			"message": "UpdateTodoCheck todoid empty",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "UpdateTodoCheck uuid error",
		})
		return
	}

	db, ctx := connector.Connector()
	err = db.Todo.UpdateOneID(todoUUID).SetCheck(check).Exec(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "UpdateTodoCheck update error",
		})
		return
	}

	defer db.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func UpdateTodoContent(c *gin.Context) {
	todoId := c.Param("todoid")
	content := c.PostForm("content")

	todoUUID, err := uuid.Parse(todoId);

	if todoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "todoid emtpy",
			"message": "UpdateTodoContent todoid empty",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "UpdateTodoContent uuid error",
		})
		return
	}

	db, ctx := connector.Connector()
	err = db.Todo.UpdateOneID(todoUUID).SetContent(content).Exec(ctx);

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "UpdateTodoContent update error",
		})
		return
	}

	defer db.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}