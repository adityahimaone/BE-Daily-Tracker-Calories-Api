package request

import "daily-tracker-calories/bussiness/users"

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
}

type SaveAvatar struct {
	Avatar string `json:"avatar"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func ToDomain(request User) *users.Domain {
	return &users.Domain{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Gender:   request.Gender,
	}
}
