package request

import "daily-tracker-calories/bussiness/users"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type SaveAvatar struct {
	Avatar string `json:"avatar"`
}

func ToDomain(request UserRegister) *users.Domain {
	return &users.Domain{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Gender:   request.Gender,
	}
}
