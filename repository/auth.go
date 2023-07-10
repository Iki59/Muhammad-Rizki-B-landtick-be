package repository

import "landtick-be/models"

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(username string) (models.User, error)
}

func (r *repository) Register(users models.User) (models.User, error) {
	err := r.db.Create(&users).Error

	return users, err
}

func (r *repository) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=?", username).Error

	return user, err
}
