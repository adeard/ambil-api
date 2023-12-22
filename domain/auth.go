package domain

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthData struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Fullname        string `json:"fullname" binding:"required"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
}
