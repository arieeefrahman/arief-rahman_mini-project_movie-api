package ratings

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Score 	   int
	MovieTitle string
	MovieID    uint
	UserID     uint
}

type UseCase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(ratingDomain *Domain) Domain
	Update(id string, ratingDomain *Domain) Domain
	Delete(id string) bool
	GetByMovieID(movieId string) []Domain
	GetByUserID(userId string) []Domain
	GetByMovieIdAndUserID(movieId string, userId string) Domain
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(ratingDomain *Domain) Domain
	Update(id string, ratingDomain *Domain) Domain
	Delete(id string) bool
	GetByMovieID(movieId string) []Domain
	GetByUserID(userId string) []Domain
	GetByMovieIdAndUserID(movieId string, userId string) Domain
	GetAvgScore(movieId string) float64
}