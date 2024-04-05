package controller

import (
	"fmt"
	"test/form"
	"test/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllCompany(c *gin.Context, db *gorm.DB){
	var Company []models.Company
	db.Preload("Address").Find(&Company)
	c.JSON(200,gin.H{"companies":Company})
}

func AddCompany(c *gin.Context, db *gorm.DB){
	var CompanyForm form.Company
	if er:= c.BindJSON(&CompanyForm); er!=nil{
		fmt.Println("er: ",er)
		return
	}
	fmt.Println("company form: ",CompanyForm)
	company := models.NewCompany(CompanyForm.Name)
	db.Create(company)
	c.JSON(200, "add company ok")
}