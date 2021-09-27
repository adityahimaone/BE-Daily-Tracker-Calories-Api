package users

import "time"

type Domain struct {
	ID        int
	Name      string
	Username  string
	Password  string
	Avatar    string
	Gender    string
	CalorieID string
	Calorie   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	RegisterUser(user *Domain) (*Domain, error)
	IsEmailAvailable(email string) (bool, error)
	Update(user *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Login(username string, password string) (*Domain, error)
}

type Repository interface {
	Insert(user *Domain) (*Domain, error)
	Update(user *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	FindByEmail(email string) (*Domain, error)
	Login(username string, password string) (*Domain, error)
}
