package request

import "daily-tracker-calories/bussiness/users"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type SaveAvatar struct {
	Avatar string `json:"avatar"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(request User) *users.Domain {
	return &users.Domain{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Gender:   request.Gender,
	}
}