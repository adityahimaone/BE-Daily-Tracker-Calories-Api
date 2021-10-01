package users

import "time"

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Avatar    string
	Gender    string
	Token     string
	CalorieID int
	Calorie   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	RegisterUser(user *Domain) (*Domain, error)
	Update(id int, user *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Login(email string, password string) (*Domain, error)
}

type Repository interface {
	Insert(user *Domain) (*Domain, error)
	Update(id int, user *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Login(email string, password string) (*Domain, error)
}
