package users

import (
	"daily-tracker-calories/bussiness/users"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Avatar   string
	Gender   string
}

func toDomain(record Users) users.Domain {
	return users.Domain{
		ID:        int(record.ID),
		Name:      record.Name,
		Email:     record.Email,
		Password:  record.Password,
		Avatar:    record.Avatar,
		Gender:    record.Gender,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Avatar:   domain.Avatar,
		Gender:   domain.Gender,
	}
}
