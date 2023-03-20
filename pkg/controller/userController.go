package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"jwt-auth/pkg/config"
	"jwt-auth/pkg/middleware"
	"jwt-auth/pkg/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

var otp int

func GenerateOtp() int {
	rand.Seed(time.Now().UnixNano())
	max_num := 999999
	min_num := 100000
	num := rand.Intn(max_num-min_num+1) + min_num
	return num
}

func SendMail(to string) {
	user := os.Getenv("USER")
	password := os.Getenv("PASS")

	// message template
	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "OTP Verification")
	otp = GenerateOtp()
	body := fmt.Sprintf("OTP for Signin: %v", otp)
	m.SetBody("text/plain", body)

	// Send message
	d := gomail.NewDialer("smtp.gmail.com", 587, user, password)

	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}

func HashPassword(password *string) {
	bytePassword := []byte(*password)
	hPassword, _ := bcrypt.GenerateFromPassword(bytePassword, 10)
	*password = string(hPassword)
}

// func ComparePassword(dbPass, pass string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
// }

func SignUp(ctx *gin.Context) {
	var user models.User
	ctx.ShouldBindJSON(&user)
	var dbUser models.User
	config.DB.Where("email = ?", user.Email).First(&dbUser)
	if dbUser.Id == 0 {
		HashPassword(&user.Password)
		config.DB.Create(&user)
		ctx.JSON(http.StatusOK, gin.H{
			"data":    user,
			"message": "Data added successfully",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user already present",
		})
	}
}

var Id int

func SignIn(ctx *gin.Context) {
	var detail models.Detail
	ctx.ShouldBindJSON(&detail)
	email := detail.Email
	var user models.User
	config.DB.Where("email = ?", email).First(&user)
	// fmt.Println(user)
	ctx.Set("userId", Id)
	if user.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "You have to signup first",
		})
	} else {
		SendMail(email)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OTP sent",
		})
	}
}

func Verify(ctx *gin.Context) {
	var detail models.Detail
	ctx.ShouldBindJSON(&detail)
	user_otp := detail.Otp
	if user_otp == otp {
		token := middleware.GenerateToken(Id)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SignedIn successfully",
			"token":   token,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid OTP",
		})
	}
}
