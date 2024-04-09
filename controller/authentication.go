package controller

import (
	"fmt"
	"net/http"
	"test/assist"
	"test/form"
	"test/models"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login godoc
// @Summary Login
// @Description Login
// @Accept  json
// @Produce  json
// @Tags Authentication
// @Param login body form.FormLogin true "login"
// @Success 200 {string} string	"OK"
// @Router /login [post]
func Login(c *gin.Context, db *gorm.DB){
	var FormLogin form.FormLogin
	if er:=c.BindJSON(&FormLogin);er!=nil{
		fmt.Println("Ko nhan dc du lieu login")
		return
	}
	var Account models.Account
	fmt.Println("formlogin user: ",FormLogin.UserName)
	result:=db.Where("username=? AND password=?",FormLogin.UserName, FormLogin.PassWord).Find(&Account)
	if result.RowsAffected==0{
		fmt.Println("ko tim thay tai khoan")
		return
	}
	expiration := time.Now().Add(24*time.Hour)
	cookie_val:=assist.RandomString(10)
	cookie:= http.Cookie{Name: "key_cookie",Value:cookie_val,Expires:expiration}
	Account.Cookie=cookie_val
	fmt.Println("cookie_val: ",Account.Cookie)
	if er:=db.Save(&Account);er!=nil{
		fmt.Println("er save acc: ",er)
	}
	http.SetCookie(c.Writer,&cookie)
	c.JSON(200,gin.H{"STATUS":"OK"})
}

func Register(c *gin.Context, db *gorm.DB){
	var FormRegister form.FormRegister
	if er:=c.BindJSON(&FormRegister); er!=nil{
		fmt.Println("Cant get json to inject to form register")
		return
	}
	if FormRegister.PassWord!=FormRegister.RePassWord{
		fmt.Println("RePassword is not same Password")
		return
	}
	var acc models.Account
	result:=db.Where("username=?",FormRegister.UserName).Find(&acc)
	if result.RowsAffected>=1 {
		fmt.Println("Username do exist")
		return
	}
	acc1:=models.NewAccount(FormRegister.UserName,FormRegister.PassWord)
	db.Create(acc1)
}