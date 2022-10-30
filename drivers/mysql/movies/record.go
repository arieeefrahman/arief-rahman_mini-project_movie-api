package movies

import (
	"time"

	movieUseCase "mini-project-movie-api/businesses/movies"
	"mini-project-movie-api/drivers/mysql/genres"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Movie struct {
	ID        	uint 			`json:"id" gorm:"primaryKey"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at"`
	Title 	  	string 			`json:"title"`
	Synopsis  	string 			`json:"synopsis"`
	Genre 		genres.Genre 	`json:"genre"`
	GenreID 	uint			`json:"genre_id"`
	ReleaseDate	datatypes.Date	`json:"release_date"`
	RatingScore float64			`json:"rating_score"`
}

func FromDomain(domain *movieUseCase.Domain) *Movie {
	return &Movie{
		ID: domain.ID,
		Title: domain.Title,
		Synopsis: domain.Synopsis,
		GenreID: domain.GenreID,
		ReleaseDate: domain.ReleaseDate,
		RatingScore: domain.RatingScore,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *Movie) ToDomain() movieUseCase.Domain {
	return movieUseCase.Domain{
		ID: rec.ID,
		Title: rec.Title,
		Synopsis: rec.Synopsis,
		GenreName: rec.Genre.Name,
		GenreID: rec.GenreID,
		ReleaseDate: rec.ReleaseDate,
		RatingScore: rec.RatingScore,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}