package signup

import (
	connector "to-do-list/db"

	"github.com/gin-gonic/gin"
)

// 새 사용자
type newUser struct {
	Name 	string 	`form:"name"`
	Email 	string 	`form:"email"`
	Password	string	`form:"password"`
}

func Signup(c *gin.Context) {
	var user newUser 
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	db, ctx := connector.Connector()
	_ , err := db.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetMemo("").
		Save(ctx) 

	if err != nil {
		panic(err)
	} 

	defer db.Close();
	if err := db.Schema.Create(ctx); err != nil {
		panic(err)
    }

}