package main

import (
	"fmt"
	"test/database"
	"test/router"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123123@tcp(127.0.0.1:3306)/inforstaff?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)

	db, err:= gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		fmt.Println("er: ",err)
		return
	}
	if er:= database.ORM(db); er!=nil{
		panic("cant ORM")
	}
	r:= gin.Default()
	router.MainRouter(r,db)
	r.Run(":8080")
}