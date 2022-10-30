package request

import (
	"mini-project-movie-api/businesses/movies"

	"github.com/go-playground/validator/v10"
	"gorm.io/datatypes"
)

type Movie struct {
	Title       string 			`json:"title" validate:"required"`
	Synopsis    string 			`json:"synopsis" validate:"required"`
	GenreID     uint   			`json:"genre_id" validate:"required"`
	ReleaseDate datatypes.Date	`json:"release_date" validate:"required"`
	RatingScore float64 			`json:"rating_score"`
}

type MovieHandleDate struct {
	Title       string 			`json:"title" validate:"required"`
	Synopsis    string 			`json:"synopsis" validate:"required"`
	GenreID     uint   			`json:"genre_id" validate:"required"`
	ReleaseDate string			`json:"release_date"`
	RatingScore float64 		`json:"rating_score"`
}

func (req *Movie) ToDomain() *movies.Domain {
	return &movies.Domain{
		Title: req.Title,
		Synopsis: req.Synopsis,
		GenreID: req.GenreID,
		ReleaseDate: req.ReleaseDate,
		RatingScore: req.RatingScore,
	}
}

func (req *Movie) Validate() error {
	validate := validator.New()
	
	err := validate.Struct(req)

	return err
}