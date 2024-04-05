package controller

import (
	"fmt"
	"test/form"
	"test/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllAddress(c *gin.Context, db *gorm.DB){
	var Address []models.Address
	db.Preload("Staff").Find(&Address)
	c.JSON(200, gin.H{"Addresses":Address})
}

func AddAddress(c *gin.Context, db *gorm.DB){
	var addressForm form.AdressForm
	c.BindJSON(&addressForm)
	fmt.Println("address form: ",addressForm)
	address:= models.NewAddresss(addressForm.AddressCompany,addressForm.AddressSpecific,addressForm.CompanyID)
	db.Create(address)
	var Company models.Company
	db.Where("id=?",addressForm.CompanyID).Find(&Company)
	fmt.Println("company name: ",Company.Name)
	Company.Address= append(Company.Address, *address)
	db.Save(&Company)
	
	
	c.JSON(200,"add address ok")
}