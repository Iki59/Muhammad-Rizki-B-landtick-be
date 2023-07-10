package repository

import (
	"landtick-be/models"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(ticket models.Transaction) (models.Transaction, error)
	// UpdateTicket(ticket models.Ticket) (models.Ticket, error)
	// DeleteTicket(ticket models.Ticket) (models.Ticket, error)
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	// err := r.db.Raw("SELECT * FROM users").Scan(&users).Error
	err := r.db.Preload("User").Preload("Ticket").Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	// err := r.db.Raw("SELECT * FROM users WHERE id = ?", ID).Scan(&user).Error
	err := r.db.First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	// err := r.db.Exec("INSERT INTO users(name, email, password) VALUES(?,?,?)", user.Name, user.Email, user.Password).Error
	err := r.db.Create(&transaction).Error
	return transaction, err
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
