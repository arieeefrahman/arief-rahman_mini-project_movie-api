package genres

import (
	"mini-project-movie-api/businesses/genres"
	ctrl "mini-project-movie-api/controllers"
	"mini-project-movie-api/controllers/genres/request"
	"mini-project-movie-api/controllers/genres/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GenreController struct {
	genreUseCase genres.UseCase
}

func NewGenreController(genreUC genres.UseCase) *GenreController {
	return &GenreController{
		genreUseCase: genreUC,
	}
}

func (gc *GenreController) GetAll(c echo.Context) error {
	genres := []response.Genre{}
	genresData := gc.genreUseCase.GetAll()

	for _, genre := range genresData {
		genres = append(genres, response.FromDomain(&genre))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all genres", genres)
}

func (gc *GenreController) GetByID(c echo.Context) error {	
	var id string = c.Param("id")
	genre := gc.genreUseCase.GetByID(id)

	if genre.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "genre not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "genre found", response.FromDomain(&genre))
}

func (gc *GenreController) Create(c echo.Context) error {
	input := request.Genre{}

	if err := c.Bind(&input); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	check := gc.genreUseCase.GetByName(input.Name)

	// handle duplicate
	if input.Name == check.Name {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "duplicate name", "")
	}

	genre := gc.genreUseCase.Create(input.ToDomain())

	if genre.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "genre already exist", "")
	}

	return ctrl.NewResponse(c, http.StatusCreated, "success", "genre created", response.FromDomain(&genre))
}

func (gc *GenreController) Update(c echo.Context) error {
	input := request.Genre{}
	var id string = c.Param("id")

	if err := c.Bind(&input); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	check := gc.genreUseCase.GetByName(input.Name)

	// handle duplicate
	if input.Name == check.Name {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "duplicate name", "")
	}

	genre := gc.genreUseCase.Update(id, input.ToDomain())

	// handle if id not found
	if genre.ID == 0 {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "id does not exist", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "genre updated", response.FromDomain(&genre))
}

func (gc *GenreController) Delete(c echo.Context) error {
	var id string = c.Param("id")

	isDeleted := gc.genreUseCase.Delete(id)

	if !isDeleted {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "genre not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "genre deleted", "")
}