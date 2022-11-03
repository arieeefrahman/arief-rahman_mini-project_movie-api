package movies

import (
	"mini-project-movie-api/businesses/movies"
	ctrl "mini-project-movie-api/controllers"
	"mini-project-movie-api/controllers/movies/request"
	"mini-project-movie-api/controllers/movies/response"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
)

type MovieController struct {
	movieUseCase movies.UseCase
}

func NewMovieController(movieUC movies.UseCase) *MovieController {
	return &MovieController{
		movieUseCase: movieUC,
	}
}

func (mc *MovieController) GetAll(c echo.Context) error {
	movies := []response.Movie{}

	moviesData := mc.movieUseCase.GetAll()

	for _, movie := range moviesData {
		movies = append(movies, response.FromDomain(movie))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all movies", movies)
}

func (mc *MovieController) GetByID(c echo.Context) error {
	var id string = c.Param("id")
	movie := mc.movieUseCase.GetByID(id)

	if movie.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "movie not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "movie found", response.FromDomain(movie))
}

func (mc *MovieController) Create(c echo.Context) error {
	inputTemp := request.MovieHandleDate{}

	if err := c.Bind(&inputTemp); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	date, _ := time.Parse("2006-01-02", inputTemp.ReleaseDate)

	input := request.Movie{
		Title: inputTemp.Title,
		Synopsis: inputTemp.Synopsis,
		GenreID: inputTemp.GenreID,
		ReleaseDate: date,
	}

	err := input.Validate()

	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	movie := mc.movieUseCase.Create(input.ToDomain())

	return ctrl.NewResponse(c, http.StatusCreated, "success", "movie created", response.FromDomain(movie))
}

func (mc *MovieController) Update(c echo.Context) error {
	var id string = c.Param("id")
	movieData := mc.movieUseCase.GetByID(id)

	if movieData.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "movie not found", "")
	}

	inputTemp := request.MovieHandleDate{}

	if err := c.Bind(&inputTemp); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	date, _ := time.Parse("2006-01-02", inputTemp.ReleaseDate)

	input := request.Movie{
		Title: inputTemp.Title,
		Synopsis: inputTemp.Synopsis,
		GenreID: inputTemp.GenreID,
		ReleaseDate: date,
	}

	err := input.Validate()
	
	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	
	movie := mc.movieUseCase.Update(id, input.ToDomain())

	return ctrl.NewResponse(c, http.StatusOK, "success", "movie updated", response.FromDomain(movie))
}

func (mc *MovieController) Delete(c echo.Context) error {
	movieId := c.Param("id")
	isSuccess := mc.movieUseCase.Delete(movieId)

	if !isSuccess {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "movie not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "movie deleted", "")
}

func (mc *MovieController) GetByGenreID(c echo.Context) error {
	var genreId string = c.Param("genre_id")
	
	movies := []response.Movie{}
	moviesData := mc.movieUseCase.GetByGenreID(genreId)

	if genreId == "" {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "genre not found", "")
	}

	if len(moviesData) == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "genre not found", "")
	}

	for _, movie := range moviesData {
		movies = append(movies, response.FromDomain(movie))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "movies found", movies)
}

func (mc *MovieController) GetLatest(c echo.Context) error {
	movies := []response.Movie{}
	moviesData := mc.movieUseCase.GetLatest()

	for _, movie := range moviesData {
		movies = append(movies, response.FromDomain(movie))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "latest movies", movies)
}

func (mc *MovieController) GetByTitle(c echo.Context) error {
	var search string = c.QueryParam("search")

	rgx := regexp.MustCompile(`[+]`)
	title := rgx.ReplaceAllString(search, " ")
	
	movie := mc.movieUseCase.GetByTitle(title)

	if movie.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "movie not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "movie found", response.FromDomain(movie))
}