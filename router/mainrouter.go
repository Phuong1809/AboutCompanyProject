package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRouter(r *gin.Engine,db *gorm.DB){
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"ping":"ok",
		})
	})
}