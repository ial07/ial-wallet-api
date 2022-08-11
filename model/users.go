package model

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Nohp     int    `json:"nohp" binding:"number"`
	Image    string `json:"image"`
}