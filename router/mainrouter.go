package router

import (
	"test/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRouter(r *gin.Engine,db *gorm.DB){
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"ping":"ok",
		})
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
	}

	addressGroup:= r.Group("/address")
	{
		addressGroup.GET("/all",func(c *gin.Context) {
			controller.GetAllAddress(c,db)
		})
		addressGroup.POST("/add",func(c *gin.Context){
			controller.AddAddress(c,db)
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
	}
}