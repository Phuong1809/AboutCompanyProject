package controller

import (
	"fmt"
	"strconv"
	"test/form"
	"test/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllCompany godoc
// @Summary return list of all
// @Description Get all company
// @Produce  application/json
// @Tags Company
// @Success 200 {object} models.Company
// @Router /company/all [get]
func GetAllCompany(c *gin.Context, db *gorm.DB){
	var Company []models.Company
	db.Preload("Address").Find(&Company)
	c.JSON(200,gin.H{"companies":Company})
}

// AddCompany godoc
// @Summary Add company
// @Description Add company
// @Accept  json
// @Produce  json
// @Tags Company
// @Param company body form.Company true "company"
// @Success 200 {string} string	"add company ok"
// @Router /company/add [post]
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

// DeleteCompany godoc
// @Summary Delete company
// @Description Delete company
// @Accept  json
// @Produce  json
// @Tags Company
// @Param id path int true "id"
// @Success 200 {string} string	"DELETE OK"
// @Router /company/delete/{id} [delete]
func DeleteCompany(c *gin.Context,db *gorm.DB){
	id,er:= strconv.ParseUint(c.Param("id"),10,32)
		fmt.Println("id: ",id)

	if er!=nil{
		fmt.Println("err parunit: ",er)
		return
	}
	db.Where("id=?",id).Delete(&models.Company{})

	c.JSON(200, "DELETE OK")
}