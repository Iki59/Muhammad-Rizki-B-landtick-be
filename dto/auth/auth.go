package authdto

type AuthRequest struct {
	FullName string `json:"full_name" form:"full_name" validation:"required"`
	Username string `json:"username" form:"username" validation:"required"`
	Email    string `json:"email" form:"email" validation:"required"`
	Password string `json:"password" form:"password" validation:"required"`
	NomorHp  string `json:"no_hp" form:"no_hp" validation:"required"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponse struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
