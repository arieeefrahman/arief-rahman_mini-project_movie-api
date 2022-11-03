package movies

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        		uint
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt
	Title 	  		string
	Synopsis  		string
	GenreName 		string
	GenreID 		uint
	ReleaseDate		time.Time
	RatingScore 	float64
}

type UseCase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(movieDomain *Domain) Domain
	Update(id string, movieDomain *Domain) Domain
	Delete(id string) bool
	GetByGenreID(genreId string) []Domain
	GetLatest() []Domain
	GetByTitle(title string) Domain
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(movieDomain *Domain) Domain
	Update(id string, movieDomain *Domain) Domain
	Delete(id string) bool
	GetByGenreID(genreId string) []Domain
	GetLatest() []Domain
	GetByTitle(title string) Domain
}