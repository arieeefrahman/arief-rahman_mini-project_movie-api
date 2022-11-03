package response

import (
	"mini-project-movie-api/businesses/movies"
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID          uint			`json:"id" gorm:"primaryKey"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`json:"deleted_at"`
	Title       string			`json:"title"`
	Synopsis    string			`json:"synopsis"`
	GenreName   string			`json:"genre_name"`
	GenreID     uint			`json:"genre_id"`
	ReleaseDate time.Time		`json:"release_date"`
	RatingScore float64			`json:"rating_score"`
}

func FromDomain(domain movies.Domain) Movie {
	return Movie{
		ID:			domain.ID,
		Title:		domain.Title,
		Synopsis:	domain.Synopsis,
		GenreName:	domain.GenreName,
		GenreID:	domain.GenreID,
		ReleaseDate: domain.ReleaseDate,
		RatingScore: domain.RatingScore,
		CreatedAt:	domain.CreatedAt,
		UpdatedAt:	domain.UpdatedAt,
		DeletedAt:	domain.DeletedAt,
	}
}