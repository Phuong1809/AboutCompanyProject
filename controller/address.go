package controller

import (
	"fmt"
	"strconv"
	"test/form"
	"test/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllAddress godoc
// @Summary return list of all
// @Description Get all address
// @Produce  application/json
// @Tags Address
// @Success 200 {object} models.Address
// @Router /address/all [get]
func GetAllAddress(c *gin.Context, db *gorm.DB){
	var Address []models.Address
	db.Preload("Staff").Find(&Address)
	c.JSON(200, gin.H{"Addresses":Address})
}

// AddAddress godoc
// @Summary Add address
// @Description Add address
// @Accept  json
// @Produce  json
// @Tags Address
// @Param address body form.AdressForm true "address"
// @Success 200 {string} string	"add address ok"
// @Router /address/add [post]
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

// DeleteAddress godoc
// @Summary Delete address
// @Description Delete address
// @Accept  json
// @Produce  json
// @Tags Address
// @Param id path int true "id"
// @Success 200 {string} string	"DELETE ADDRESS OK"
// @Router /address/delete/{id} [delete]
func DeleteAddress(c*gin.Context, db *gorm.DB){
	id,_:= strconv.ParseUint(c.Param("id"),10,32)
	db.Where("id=?",id).Delete(&models.Address{})
	c.JSON(200,"DELETE ADDRESS OK")
}