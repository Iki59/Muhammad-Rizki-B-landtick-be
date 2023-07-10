package models

type Station struct {
	ID   int    `json:"id" gorm:"primaryKey:autoIncrement"`
	Name string `json:"name" gorm:"not null"`
}
