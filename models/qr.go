package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"test/assist"

	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

type QR struct {
	gorm.Model
	FilePatch string
	StaffId uint
	Staff Staff `gorm:"foreignKey:StaffId"`
}

func CreateQRCode(URL string,db *gorm.DB){
	fmt.Println(URL)
	dir:= "D:\\Workspace_GO\\test_1\\qr_code"
	qrCode, err := qrcode.Encode(URL,qrcode.Medium, 256)
	if err!=nil{
		fmt.Println("cant create qrcode")
		return
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
        fmt.Println("Folder does not exist")
        return
    }
	nameFile:=assist.RandomString(4)+".png"
	file, err := os.Create(filepath.Join(dir, nameFile))
    if err != nil {
        fmt.Println("Lỗi khi tạo tệp hình ảnh:", err)
        return
    }
	if _, err := file.Write(qrCode); err != nil {
        fmt.Println("Lỗi khi ghi hình ảnh:", err)
        return
    }
	fmt.Println("Tạo và lưu mã QR thành công.")
	parts := strings.Split(URL, "/")
	fmt.Println("parts: ",parts[len(parts)-1])
	StaffID,er:= strconv.ParseUint(parts[len(parts)-1],10,32)
	fmt.Println("parts len ", len(parts))
	fmt.Println("StaffID CONV: ",StaffID)
	if er!=nil{
		fmt.Println("cant parse unit ", er)
		return
	}
	var QR QR
	QR.StaffId=uint(StaffID)
	QR.FilePatch=dir+ "\\" + nameFile
	var Staff Staff
	db.Where("id=?",StaffID).Find(&Staff)
	Staff.QR = append(Staff.QR, QR)
	db.Save(&Staff)
}