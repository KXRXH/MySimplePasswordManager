package database

import (
	"errors"

	"github.com/kxrxh/password-manager/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PasswordField struct {
	gorm.Model
	Key      string
	Password string
}

var db *gorm.DB

func InitDb(path string) {
	var err error
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&PasswordField{})
}

func AddToDb(key string, value string) {
	passwordField := PasswordField{Key: key, Password: utils.EncodeString(value)}
	db.Create(&passwordField)
}

func GetByKey(key string) (string, error) {
	var passField = PasswordField{}
	db.Where("key = ?", key).Find(&passField)
	if passField.Password == "" {
		return "nil", errors.New("there are no keys with the given value")
	}
	password, _ := utils.DecodeString(passField.Password)
	return password, nil
}

func GetAllPassword() []PasswordField {
	var fields []PasswordField
	db.Select("key", "password").Find(&fields)
	return fields
}
