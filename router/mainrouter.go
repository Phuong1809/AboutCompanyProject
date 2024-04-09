package router

import (
	"fmt"
	"test/controller"
	"test/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRouter(r *gin.Engine,db *gorm.DB){
	r.Use(middleware.AuthSession(db))

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"ping":"ok",
		})
	})

	r.POST("/login",func(c *gin.Context) {
		controller.Login(c,db)
	})

	staffGroup:=r.Group("/staff")
	{
		staffGroup.GET("/all",func(c*gin.Context){
			controller.GetStaffs(c,db)
		})
		staffGroup.GET("/:id",func(c *gin.Context) {
			controller.GetStaffById(c,db)

		})
		staffGroup.POST("/add",func(c *gin.Context) {
			controller.AddStaff(c,db)
		})
		staffGroup.DELETE("/delete/:id",func(c*gin.Context){
			controller.DeleteStaff(c,db)
		})
	}

	addressGroup:= r.Group("/address")
	{
		addressGroup.GET("/all",func(c *gin.Context) {
			controller.GetAllAddress(c,db)
		})
		addressGroup.POST("/add",func(c *gin.Context){
			fmt.Println("add address router")
			controller.AddAddress(c,db)
		})
		addressGroup.DELETE("/delete",func(c *gin.Context) {
			controller.DeleteAddress(c,db)
		})
	}

	companyGroup:= r.Group("/company")
	{

		companyGroup.GET("/all",func(c*gin.Context){
			controller.GetAllCompany(c,db)
		})
		companyGroup.POST("/add",func(c *gin.Context) {
			controller.AddCompany(c,db)
		})
		companyGroup.DELETE("/delete/:id",func(c *gin.Context) {
			controller.DeleteCompany(c,db)
		})
	}
	r.GET("/admin",func(c *gin.Context) {
		c.JSON(200,"HELLO ADMIN")
	})
}