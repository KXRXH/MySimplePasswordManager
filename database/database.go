package database

import (
	"github.com/kxrxh/password-manager/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PasswordField struct {
	gorm.Model
	Key      string
	Password string
}

var db *gorm.DB

func InitDb(path string) {
	var err error
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&PasswordField{})
}

func AddToDb(key string, value string) {
	passwordField := PasswordField{Key: key, Password: utils.EncodeString(value)}
	db.Create(&passwordField)
}

func GetByKey(key string) string {
	var passField = PasswordField{Key: key}
	db.First(&passField)
	password, err := utils.DecodeString(passField.Password)
	if err != nil {
		panic(err)
	}
	return password
}

func GetAllPassword() []PasswordField {
	var fields []PasswordField
	db.Select("key", "password").Find(&fields)
	return fields
}
