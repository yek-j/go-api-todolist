package todo

import (
	"net/http"
	"strconv"
	"to-do-list/api/common"
	connector "to-do-list/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type newTodo struct {
	Content string 
	Check bool 
	UserID string 
}

func Addtodo(c *gin.Context) {
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
	
	var todo newTodo 
	todo.Content = c.PostForm("content")
	todo.Check, _ = strconv.ParseBool(c.PostForm("check"))
	todo.UserID = c.PostForm("userid")

	userID, err := uuid.Parse(verifiedClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Addtodo uuid err",
		})
		return
	}

	db, ctx := connector.Connector()
	_ , err = db.Todo.
		Create().
		SetContent(todo.Content).
		SetCheck(todo.Check).
		SetUserID(userID).
		Save(ctx)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Addtodo save err",
		})
		return
	} 	

	defer db.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
