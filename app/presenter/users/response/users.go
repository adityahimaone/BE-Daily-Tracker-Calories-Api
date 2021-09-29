package response

import (
	"daily-tracker-calories/bussiness/users"
	"time"
)

type UserRegister struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain users.Domain) UserRegister {
	return UserRegister{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Gender:    domain.Gender,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type UserLogin struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FromDomainLogin(domain users.Domain) UserLogin {
	return UserLogin{
		ID:    domain.ID,
		Email: domain.Email,
		Token: domain.Token,
	}
}
