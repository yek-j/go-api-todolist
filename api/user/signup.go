package signup

import (
	"fmt"
	"net/http"
	connector "to-do-list/db"
	"to-do-list/ent/user"

	"github.com/gin-gonic/gin"
)

// 새 사용자
type newUser struct {
	Name 	string 	`form:"name"`
	Email 	string 	`form:"email"`
	Password	string	`form:"password"`
}

// 회원가입
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

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// 이메일 중복 확인
func Emailcheck(c *gin.Context) {
	var email = c.PostForm("email")

	db, ctx := connector.Connector()
	check, err := db.User.
		Query().
		Where(user.EmailIn(email)).
		All(ctx)

	fmt.Println(len(check));

	if err != nil {
		panic(err)
	} 

	defer db.Close();
	
	if len(check) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "using",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}