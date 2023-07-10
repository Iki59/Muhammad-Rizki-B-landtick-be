package usersdto

type CreateUserRequest struct {
	FullName string `json:"full_name" form:"full_name" validation:"required"`
	Username string `json:"username" form:"username" validation:"required"`
	Email    string `json:"email" form:"email" validation:"required"`
	NomorHp  string `json:"no_hp" form:"no_hp" validation:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" form:"full_name"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	NomorHp  string `json:"no_hp" form:"no_hp"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	NomorHp  string `json:"no_hp" form:"no_hp" validate:"required"`
}
