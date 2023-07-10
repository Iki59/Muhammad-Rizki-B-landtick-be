package models

type Transaction struct {
	ID       int                       `json:"id" gorm:"primary_key:auto_increment"`
	UserID   int                       `json:"-"`
	User     UserResponseTransaction   `json:"user" gorm:"foreignKey:UserID"`
	TicketID int                       `json:"-"`
	Ticket   TicketResponseTransaction `json:"ticket" gorm:"foreignKey:TicketID"`
	Image    string                    `json:"image"`
}
