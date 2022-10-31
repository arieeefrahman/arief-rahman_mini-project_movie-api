package main

import (
	"encoding/json"
	_middlewares "mini-project-movie-api/app/middlewares"
	_routes "mini-project-movie-api/app/routes"
	_util "mini-project-movie-api/utils"
	"net/http"
	"strconv"
	"testing"

	// "time"

	_driverFactory "mini-project-movie-api/drivers"

	_dbDriver "mini-project-movie-api/drivers/mysql"
	"mini-project-movie-api/drivers/mysql/genres"
	"mini-project-movie-api/drivers/mysql/movies"
	"mini-project-movie-api/drivers/mysql/ratings"
	"mini-project-movie-api/drivers/mysql/users"

	_genreUseCase "mini-project-movie-api/businesses/genres"
	_genreController "mini-project-movie-api/controllers/genres"

	_movieUseCase "mini-project-movie-api/businesses/movies"
	_movieController "mini-project-movie-api/controllers/movies"

	_ratingUseCase "mini-project-movie-api/businesses/ratings"
	_ratingController "mini-project-movie-api/controllers/ratings"

	_userUseCase "mini-project-movie-api/businesses/users"
	_userController "mini-project-movie-api/controllers/users"
	_userRequest "mini-project-movie-api/controllers/users/request"

	echo "github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
	// "gorm.io/datatypes"
)

func newApp() *echo.Echo {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
		DB_HOST: _util.GetConfig("DB_HOST"),
		DB_PORT: _util.GetConfig("DB_PORT"),
		DB_NAME: _util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT:       _util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	app := echo.New()

	genreRepo := _driverFactory.NewGenreRepository(db)
	genreUseCase := _genreUseCase.NewGenreUsecase(genreRepo)
	genreCtrl := _genreController.NewGenreController(genreUseCase)

	movieRepo := _driverFactory.NewMovieRepository(db)
	movieUseCase := _movieUseCase.NewMovieUseCase(movieRepo)
	movieCtrl := _movieController.NewMovieController(movieUseCase)

	ratingRepo := _driverFactory.NewRatingRepository(db)
	ratingUseCase := _ratingUseCase.NewRatingUseCase(ratingRepo)
	ratingCtrl := _ratingController.NewRatingController(ratingUseCase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUseCase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		JWTMiddleware: configJWT.Init(),
		GenreController: *genreCtrl,
		MovieController: *movieCtrl,
		RatingController: *ratingCtrl,
		AuthController: *userCtrl,
	}

	routesInit.RouteRegister(app)

	return app
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		configDB := _dbDriver.ConfigDB{
			DB_USERNAME: _util.GetConfig("DB_USERNAME"),
			DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
			DB_HOST:     _util.GetConfig("DB_HOST"),
			DB_PORT:     _util.GetConfig("DB_PORT"),
			DB_NAME:     _util.GetConfig("DB_TEST_NAME"),
		}

		db := configDB.InitDB()

		_dbDriver.CleanSeeders(db)
	}
}

func getJwtToken(t *testing.T) string {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
		DB_HOST:     _util.GetConfig("DB_HOST"),
		DB_PORT:     _util.GetConfig("DB_PORT"),
		DB_NAME:     _util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	var userRequest *_userRequest.User = &_userRequest.User{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func getUser() users.User {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
		DB_HOST:     _util.GetConfig("DB_HOST"),
		DB_PORT:     _util.GetConfig("DB_PORT"),
		DB_NAME:     _util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	return user
}

func getMovie() movies.Movie {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
		DB_HOST:     _util.GetConfig("DB_HOST"),
		DB_PORT:     _util.GetConfig("DB_PORT"),
		DB_NAME:     _util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	movie := _dbDriver.SeedMovie(db)

	return movie
}

func getGenre() genres.Genre {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
		DB_HOST:     _util.GetConfig("DB_HOST"),
		DB_PORT:     _util.GetConfig("DB_PORT"),
		DB_NAME:     _util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	genre := _dbDriver.SeedGenre(db)

	return genre
}

func getRating() ratings.Rating {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
		DB_HOST:     _util.GetConfig("DB_HOST"),
		DB_PORT:     _util.GetConfig("DB_PORT"),
		DB_NAME:     _util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	rating := _dbDriver.SeedRating(db)

	return rating
}

func TestSignup_Success(t *testing.T) {
	var userRequest *_userRequest.User = &_userRequest.User{
		Email:    "test@mail.com",
		Password: "123123",
	}

	apitest.
		New().
		Handler(newApp()).
		Post("/api/v1/users/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestSignup_ValidationFailed(t *testing.T) {
	var userRequest *_userRequest.User = &_userRequest.User{
		Email: "",
		Password: "",
	}

	apitest.
		New().
		Handler(newApp()).
		Post("/api/v1/users/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	user := getUser()

	var userRequest *_userRequest.User = &_userRequest.User{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *_userRequest.User = &_userRequest.User{
		Email:    "",
		Password: "",
	}

	apitest.
		New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *_userRequest.User = &_userRequest.User{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.
		New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}

func TestCreateGenre_Success(t *testing.T) {
	var token string = getJwtToken(t)

	var genreRequest *genres.Genre = &genres.Genre{
		Name: "test",
	}

	apitest.
		New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/genres").
		Header("Authorization", token).
		JSON(genreRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateGenre_ValidationFailed(t *testing.T) {
	var token string = getJwtToken(t)

	var genreRequest *genres.Genre = &genres.Genre{
		Name: "",
	}

	apitest.
		New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/genres").
		Header("Authorization", token).
		JSON(genreRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestGetGenres_Success(t *testing.T) {
	var token string = getJwtToken(t)

	apitest.
		New().
		Handler(newApp()).
		Get("/api/v1/genres").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetGenreByID_Success(t *testing.T) {
	var token string = getJwtToken(t)
	genre := getGenre()
	genreID := strconv.Itoa(int(genre.ID))

	apitest.
		New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/genres/" + genreID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetGenreByID_Failed(t *testing.T) {
	var token string = getJwtToken(t)

	apitest.
		New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/genres/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}