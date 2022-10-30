package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email  	  string
	Password  string
}

type UseCase interface {
	Signup(userDomain *Domain) Domain
	Login(userDomain *Domain) string
}

type Repository interface {
	Signup(userDomain *Domain) Domain
	GetByEmail(userDomain *Domain) Domain
}