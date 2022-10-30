package genres

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string 
}

type UseCase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(genreDomain *Domain) Domain
	Update(id string, genreDomain *Domain) Domain
	Delete(id string) bool
	GetByName(name string) Domain
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(genreDomain *Domain) Domain
	Update(id string, genreDomain *Domain) Domain
	Delete(id string) bool
	GetByName(name string) Domain
}