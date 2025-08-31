package models

type User struct {
	BaseModel
	Account  string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Name     string
}

type CurrentUser struct {
	Id      string `json:"id" binding:"required"`
	Account string `json:"account" binding:"required,min=4"`
	Name    string `json:"name" binding:"required,min=2"`
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
