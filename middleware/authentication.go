package middleware

import (
	"fmt"
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthSession(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context)  {
		path:= c.Request.URL.Path
		var flag string
		flag="green"
		Cookie,er := c.Cookie("key_cookie")
	
		if er!=nil ||Cookie==""{
			fmt.Println("err auth session: ",er)
			flag="red"
		}
		var acc models.Account
		if row:=db.Where("cookie=?",Cookie).First(&acc).RowsAffected; row==0{
			flag="red"
		}
		fmt.Println("cookie: ", acc.Cookie)
		if flag=="red"{
			fmt.Println("flag red")
			exception:=[]string{"/staff/add","/company/add","/address/add","/admin"}
			for _,endpoint := range exception{
				if path==endpoint{
					c.Redirect(http.StatusSeeOther,"/login")
					c.Abort()
				}
			}
		}
		c.Next()
		
		
	}

}