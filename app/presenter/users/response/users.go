package response

import (
	"daily-tracker-calories/bussiness/users"
	"time"
)

type Users struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Gender:    domain.Gender,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
