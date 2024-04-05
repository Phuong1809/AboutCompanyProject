package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"test/assist"
	"test/form"
	"test/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStaffs(c *gin.Context,db *gorm.DB){
	var staff []models.Staff
	db.Preload("QR").Preload("Company").Find(&staff)
	c.JSON(http.StatusOK,gin.H{"staffs":staff})
}

func GetStaffById(c*gin.Context, db *gorm.DB){
	IdStr := c.Param("id")
	id,err:=strconv.ParseUint(IdStr,10,32)
	if err!=nil{
		panic(err)
	}
	var staff models.Staff
	db.Where("id=?",id).Preload("QR").Preload("Company").Find(&staff)
	c.JSON(http.StatusOK, gin.H{"staff":staff})
}

func AddStaff(c *gin.Context, db *gorm.DB){
	file, header, er := c.Request.FormFile("staff")
	if er!=nil{
		fmt.Println("loi k lay dc file")
		return
	}
	fmt.Println("file's name: ", header.Filename)
	FileByte,er:= ioutil.ReadAll(file)
	if er!=nil{
		fmt.Println("ko chuyen ra json ra byte dc")
		return
	}
	var FormStaff form.StaffForm
	json.Unmarshal([]byte(FileByte),&FormStaff)
	fmt.Println("formstaff: ",FormStaff)
	staff:= models.NewStaff(FormStaff.StaffName,FormStaff.Role,FormStaff.NameCompany,
		FormStaff.TelePhone,FormStaff.MobilePhone,FormStaff.Fax,
		FormStaff.Email,FormStaff.AddressID)
	db.Create(staff)
	var Address models.Address
	db.Where("id=?",FormStaff.AddressID).Find(&Address)
	Address.Staff = append(Address.Staff, *staff)
	db.Save(Address)
	fmt.Println(staff.ID)
	URL:= "http://localhost:8080/staff/"+ strconv.FormatUint(uint64(staff.ID),10)
	models.CreateQRCode(URL,db)
	fmt.Println("hello toi dang nhan file anh")
	file,header,er =c.Request.FormFile("image")
	if er!=nil{
		fmt.Println("loi ko nhan dc file ",er)
		return
	}
	dir:= "D:\\Workspace_GO\\test_1\\avatar"
	if _,er:=os.Stat(dir); os.IsNotExist(er){
		fmt.Println("folder k ton tai")
		return
	}
	file_name_image:=assist.RandomString(4)+header.Filename
	out, er:=os.Create(filepath.Join(dir,file_name_image))
	if er!=nil{
		fmt.Println("loi tao file os")
	}
	_, er = io.Copy(out,file)
	if er!=nil{
		fmt.Println("Day file vao folder that bai")
		return
	}
	fmt.Println("ok da thanh cong them anh vao folder")
	staff.FilePathAvatar="D:\\Workspace_GO\\test_1\\avatar\\"+file_name_image
	db.Save(staff)
	c.JSON(http.StatusOK,"add staff ok")
}