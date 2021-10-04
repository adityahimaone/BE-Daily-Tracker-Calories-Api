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
	Token string `json:"token"`
}
type Validate struct {
	Error string `json:"error"`
}

