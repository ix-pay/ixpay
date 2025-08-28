package models

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Account  string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Name     string
}

type RegisterUser struct {
	Account  string `json:"account" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required,min=2"`
}

type ProfileUser struct {
	Id      string `json:"id" binding:"required"`
	Account string `json:"account" binding:"required,min=4"`
	Name    string `json:"name" binding:"required,min=2"`
}

type LoginCredentials struct {
	Account  string `json:"account" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=8"`
}

type TokenUser struct {
	Token string      `json:"token"`
	User  ProfileUser `json:"user"`
}

func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func AuthenticateUser(account, password string) (*User, error) {
	var user User
	log.Println("account=" + account)
	log.Println("password=" + password)
	if err := DB.Where("account = ?", account).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("记录不存在")
		} else {
			log.Println("查询出错:", err)
		}
		return nil, err
	}
	// 修正点2：添加密码为空的安全检查
	if user.Password == "" {
		return nil, fmt.Errorf("用户密码字段为空")
	}

	log.Println("CompareHashAndPassword=" + user.Password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return &user, nil
	}
	return nil, gorm.ErrRecordNotFound
}
