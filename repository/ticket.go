package repository

import (
	"landtick-be/models"
)

type TicketRepository interface {
	FindTickets() ([]models.Ticket, error)
	GetTicket(ID int) (models.Ticket, error)
	CreateTicket(ticket models.Ticket) (models.Ticket, error)
	// UpdateTicket(ticket models.Ticket) (models.Ticket, error)
	// DeleteTicket(ticket models.Ticket) (models.Ticket, error)
}

func (r *repository) FindTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket
	// err := r.db.Raw("SELECT * FROM users").Scan(&users).Error
	err := r.db.Preload("StartStation").Preload("DestinationStation").Find(&tickets).Error

	return tickets, err
}

func (r *repository) GetTicket(ID int) (models.Ticket, error) {
	var ticket models.Ticket
	// err := r.db.Raw("SELECT * FROM users WHERE id = ?", ID).Scan(&user).Error
	err := r.db.First(&ticket, ID).Error

	return ticket, err
}

func (r *repository) CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	// err := r.db.Exec("INSERT INTO users(name, email, password) VALUES(?,?,?)", user.Name, user.Email, user.Password).Error
	err := r.db.Create(&ticket).Error
	return ticket, err
}

// func (r *repository) UpdateTicket(ticket models.Ticket) (models.Ticket, error) {
// 	// err := r.db.Raw("UPDATE users SET name=?, email=?, password=?, updated_at=? WHERE id = ?", user.Name, user.Email, user.Password, user.UpdatedAt, ID).Scan(&user).Error
// 	err := r.db.Save(&ticket).Error
// 	return ticket, err
// }

// func (r *repository) DeleteTicket(ticket models.Ticket) (models.Ticket, error) {
// 	// err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error
// 	err := r.db.Delete(&ticket).Error
// 	return ticket, err
// }
