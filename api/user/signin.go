package user

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"to-do-list/api/common"
	connector "to-do-list/db"
	"to-do-list/ent/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func Signin(c *gin.Context) {
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

	defer db.Close()

	// email check
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"signin": "faile",
			"message": "로그인 실패",
		})
		return
	}

	// password check 
	if !HashCheck(password, user[0].Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"signin": "faile",
			"message": "로그인 실패",
		})
	} else {
		
		// 토큰 발생 
		token, err := GenerateToken(user[0].Name, user[0].ID.String())
		fmt.Println(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"signin": "faile",
				"error": err,
				"message": "로그인 실패",
			})	
		} else {
			c.JSON(http.StatusOK, gin.H{
				"signin": "success",
				"token": token,
				"message": "로그인 성공",
			})
		}
	}
	
}

func HashCheck(password, hash string) bool {
	check := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return check == nil
}

func GenerateToken(name string, id string) (string, error) {
		
	claims := &common.JwtClaims{
		UserNM: name,
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()

	if err != nil {
		return "", err
	}

	secretkey := viper.GetString("secretkey")

	result, err := token.SignedString([]byte(secretkey))

	if err != nil {
		return "", err
	}

	return result, nil
}