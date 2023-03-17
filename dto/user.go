package dto

type UserRegister struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UserUpdate struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name,omitempy"`
	Email       string `json:"email,omitempy"`
	PhoneNumber string `json:"phone_number,omitempy"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
