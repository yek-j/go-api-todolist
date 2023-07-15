package user

import (
	"fmt"
	"log"
	"net/http"
	connector "to-do-list/db"
	"to-do-list/ent/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signin(c *gin.Context) {
	fmt.Println("signin test")

	email := c.PostForm("email")
	password := c.PostForm("password")

	// user 정보
	db, ctx := connector.Connector()
	user, err := db.User.Query().
			Where(user.Email(email)).
			All(ctx)
	
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"signin": "error",
			"message": "Signin Error",
		})
	} 

	fmt.Println(user)

	// email check
	if len(user) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"signin": "faile",
			"message": "로그인 실패",
		})
		return
	}

	// password check 
	if !HashCheck(password, user[0].Password) {
		c.JSON(http.StatusOK, gin.H{
			"signin": "faile",
			"message": "로그인 실패",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"signin": "success",
			"message": "로그인 성공",
		})
	}

}

func HashCheck(password, hash string) bool {
	check := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return check == nil
}