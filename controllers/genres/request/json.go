package request

import (
	"mini-project-movie-api/businesses/genres"

	"github.com/go-playground/validator/v10"
)

type Genre struct {
	Name string `json:"name" validate:"required"`
}

func (req *Genre) ToDomain() *genres.Domain {
	return &genres.Domain{
		Name: req.Name,
	}
}

func (req *Genre) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}