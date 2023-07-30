package user

import (
	"net/http"
	connector "to-do-list/db"
	"to-do-list/ent/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 새 사용자
type newUser struct {
	Name 	string
	Email 	string
	Password	string
}

// 회원가입
func Signup(c *gin.Context) {
	var user newUser 
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")

	hashPwd, hashErr := HashPassword(user.Password)

	if hashErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "hash err",
		})
		return
	}

	db, ctx := connector.Connector()
	_ , err := db.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(hashPwd).
		SetMemo("").
		Save(ctx) 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "signup err",
		})
		return
	} 

	defer db.Close()
	if err := db.Schema.Create(ctx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "signup err",
		})
		return
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

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Emailcheck DB err",
		})
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

// 비밀번호 hash
func HashPassword(password string)(string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(hashPwd), err
}