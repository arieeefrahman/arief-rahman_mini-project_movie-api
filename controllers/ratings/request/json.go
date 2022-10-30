package request

import (
	"mini-project-movie-api/businesses/ratings"

	"github.com/go-playground/validator/v10"
)

type Rating struct {
	Score   int  `json:"score" validate:"required"`
	MovieID uint `json:"movie_id" validate:"required"`
	UserID  uint `json:"user_id"`
}

func (req *Rating) ToDomain() *ratings.Domain {
	return &ratings.Domain{
		Score: req.Score,
		MovieID: req.MovieID,
		UserID: req.UserID,
	}
}

func (req *Rating) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}