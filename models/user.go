package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey:autoIncrement"`
	FullName string `json:"full_name" gorm:"type: varchar(255)"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	NomorHp  string `json:"no_hp" gorm:"type: varchar(255)"`
}

type UserResponseTransaction struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	NomorHp  string `json:"no_hp"`
}

type UserResponseTicket struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func (UserResponseTransaction) TableName() string {
	return "users"
}

func (UserResponseTicket) TableName() string {
	return "users"
}
