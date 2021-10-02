package response

import (
	"daily-tracker-calories/bussiness/users"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
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
