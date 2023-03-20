package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password"`
}

type Detail struct {
	Email string `json:"email"`
	Otp   int    `json:"otp"`
}
