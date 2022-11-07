package main

import (
	"encoding/json"
	_middlewares "mini-project-movie-api/app/middlewares"
	_routes "mini-project-movie-api/app/routes"
	_util "mini-project-movie-api/utils"
	"net/http"
	"regexp"
	"strconv"
	"testing"
	"time"

	_driverFactory "mini-project-movie-api/drivers"

	_genreUseCase "mini-project-movie-api/businesses/genres"
	_genreController "mini-project-movie-api/controllers/genres"
	_dbDriver "mini-project-movie-api/drivers/mysql"
	"mini-project-movie-api/drivers/mysql/genres"
	"mini-project-movie-api/drivers/mysql/movies"
	"mini-project-movie-api/drivers/mysql/ratings"
	"mini-project-movie-api/drivers/mysql/users"

	_movieUseCase "mini-project-movie-api/businesses/movies"
	_movieController "mini-project-movie-api/controllers/movies"

	_ratingUseCase "mini-project-movie-api/businesses/ratings"
	_ratingController "mini-project-movie-api/controllers/ratings"

	_userUseCase "mini-project-movie-api/businesses/users"
	_userController "mini-project-movie-api/controllers/users"
	_userRequest "mini-project-movie-api/controllers/users/request"

	echo "github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func newApp() *echo.Echo {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
		DB_HOST: _util.GetConfig("DB_TEST_HOST"),
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
			DB_HOST:     _util.GetConfig("DB_TEST_HOST"),
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
		DB_HOST:     _util.GetConfig("DB_TEST_HOST"),
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
		DB_HOST:     _util.GetConfig("DB_TEST_HOST"),
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
		DB_HOST:     _util.GetConfig("DB_TEST_HOST"),
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
		DB_HOST:     _util.GetConfig("DB_TEST_HOST"),
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
		DB_HOST:     _util.GetConfig("DB_TEST_HOST"),
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

	apitest.New().
		Observe(cleanup).
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

	apitest.New().
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

	apitest.New().
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

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}

func TestGetGenres_Success(t *testing.T) {
	var token string = getJwtToken(t)

	apitest.New().
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
	id := strconv.Itoa(int(genre.ID))

	apitest.New().
		Handler(newApp()).
		Get("/api/v1/genres/" + id).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetGenreByID_Failed(t *testing.T) {
	var token string = getJwtToken(t)

	apitest.New().
		Handler(newApp()).
		Get("/api/v1/genres/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateGenre_Success(t *testing.T) {
	var token string = getJwtToken(t)

	var genreRequest *genres.Genre = &genres.Genre{
		Name: "test",
	}

	apitest.New().
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

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/genres").
		Header("Authorization", token).
		JSON(genreRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateGenre_Success(t *testing.T) {
	var token string = getJwtToken(t)
	
	genre := getGenre()
	genreID := strconv.Itoa(int(genre.ID))

	var genreRequest *genres.Genre = &genres.Genre{
		Name: "testupdate",
	}

	apitest.
		New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/api/v1/genres/" + genreID).
		Header("Authorization", token).
		JSON(genreRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateGenre_ValidationFailed(t *testing.T) {
	var token string = getJwtToken(t)

	genre := getGenre()
	genreID := strconv.Itoa(int(genre.ID))

	var genreRequest *genres.Genre = &genres.Genre{}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/api/v1/genres/" + genreID).
		Header("Authorization", token).
		JSON(genreRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeleteGenre_Success(t *testing.T) {
	var token string = getJwtToken(t)
	
	genre := getGenre()
	genreID := strconv.Itoa(int(genre.ID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Delete("/api/v1/genres/" + genreID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()

}

func TestDeleteGenre_Failed(t *testing.T) {
	var token string = getJwtToken(t)
	
	apitest.New().
		Handler(newApp()).
		Delete("/api/v1/genres/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestGetMovies_Success(t *testing.T) {
	token := getJwtToken(t)
	
	apitest.New().
		Handler(newApp()).
		Get("/api/v1/movies").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetMovieByID_Success(t *testing.T) {
	token := getJwtToken(t)

	movie := getMovie()
	id := strconv.Itoa(int(movie.ID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/movies/" + id).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetMovieByID_NotFound(t *testing.T) {
	token := getJwtToken(t)

	apitest.New().
		Handler(newApp()).
		Get("/api/v1/movies/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateMovie_Success(t *testing.T) {
	genre := getGenre()

	releaseDate := "2020-01-01"
	date, _ := time.Parse("2006-01-02", releaseDate)
	

	var movieRequest *movies.Movie = &movies.Movie{
		Title: "test",
		Synopsis: "test",
		GenreID: genre.ID,
		ReleaseDate: date,
	}

	var token string = getJwtToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/movies").
		Header("Authorization", token).
		JSON(movieRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateMovie_ValidationFailed(t *testing.T) {
	token := getJwtToken(t)

	var movieRequest *movies.Movie = &movies.Movie{}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/movies").
		Header("Authorization", token).
		JSON(movieRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateMovie_Success(t *testing.T) {
	token := getJwtToken(t)

	movie := getMovie()
	id := strconv.Itoa(int(movie.ID))

	var movieRequest *movies.Movie = &movies.Movie{
		Title: "abcd",
		Synopsis: "abcd",
		GenreID: movie.GenreID,
		ReleaseDate: movie.ReleaseDate,
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/api/v1/movies/" + id).
		Header("Authorization", token).
		JSON(movieRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateMovie_ValidationFailed(t *testing.T) {
	token := getJwtToken(t)

	movie := getMovie()
	id := strconv.Itoa(int(movie.ID))

	var movieRequest *movies.Movie = &movies.Movie{}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/api/v1/movies/" + id).
		Header("Authorization", token).
		JSON(movieRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeleteMovie_Success(t *testing.T) {
	token := getJwtToken(t)

	movie := getMovie()
	id := strconv.Itoa(int(movie.ID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Delete("/api/v1/movies/" + id).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteMovie_Failed(t *testing.T) {
	token := getJwtToken(t)

	apitest.New().
		Handler(newApp()).
		Delete("/api/v1/movies/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestGetMoviesByGenreID_Success(t *testing.T) {
	token := getJwtToken(t)
	
	movie := getMovie()
	genreID := strconv.Itoa(int(movie.GenreID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/movies/genre/" + genreID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetMoviesByGenreID_NotFound(t *testing.T) {
	token := getJwtToken(t)
	
	apitest.New().
		Handler(newApp()).
		Get("/api/v1/movies/genre/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestGetLatestMovies_Success(t *testing.T) {
	token := getJwtToken(t)

	apitest.New().
		Handler(newApp()).
		Get("/api/v1/movies/latest").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetMovieByTitle_Success(t *testing.T) {
	token := getJwtToken(t)

	movie := getMovie()

	movieTitle := movie.Title

	rgx := regexp.MustCompile(`[+]`)
	title := rgx.ReplaceAllString(movieTitle, " ")

	apitest.New().Observe(cleanup).
		Handler(newApp()).Get("/api/v1/movies/title?").Query("search", title).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetMovieByTitle_NotFound(t *testing.T) {
	token := getJwtToken(t)

	apitest.New().Handler(newApp()).Get("/api/v1/movies/title?").Query("search", "").Header("Authorization", token).Expect(t).Status(http.StatusNotFound).End()
}

func TestGetRatings_Success(t *testing.T) {
	token := getJwtToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/api/v1/ratings").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetRatingByID_Success(t *testing.T) {
	token := getJwtToken(t)

	rating := getRating()
	ratingID := strconv.Itoa(int(rating.ID))

	apitest.New().Observe(cleanup).Handler(newApp()).Get("/api/v1/ratings/" + ratingID).Header("Authorization", token).Expect(t).Status(http.StatusOK).End()
}

func TestGetRatingByID_NotFound(t *testing.T) {
	token := getJwtToken(t)

	apitest.New().Handler(newApp()).Get("/api/v1/ratings/0").Header("Authorization", token).Expect(t).Status(http.StatusNotFound).End()
}


func TestCreateRating_Success(t *testing.T) {
	token := getJwtToken(t)
	movie := getMovie()
	user := getUser()
	
	var ratingRequest *ratings.Rating = &ratings.Rating{
		Score: 10,
		MovieID: movie.ID,
		UserID: user.ID,
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/v1/ratings").
		Header("Authorization", token).
		JSON(ratingRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateRating_ValidationFailed(t *testing.T) {
	var token string = getJwtToken(t)
	
	var ratingRequest *ratings.Rating = &ratings.Rating{}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/ratings").
		Header("Authorization", token).
		JSON(ratingRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateRating_Success(t *testing.T) {
	token := getJwtToken(t)

	rating := getRating()
	ratingId := strconv.Itoa(int(rating.ID))
	
	var ratingRequest *ratings.Rating = &ratings.Rating{
		Score: 10,
		MovieID: rating.MovieID,
		UserID: rating.UserID,
	}

	apitest.New().Observe(cleanup).Handler(newApp()).Put("/api/v1/ratings/" + ratingId).Header("Authorization", token).JSON(ratingRequest).Expect(t).Status(http.StatusOK).End()
}

func TestUpdateRating_ValidationFailed(t *testing.T) {
	token := getJwtToken(t)

	rating := getRating()
	ratingId := strconv.Itoa(int(rating.ID))
	
	var ratingRequest *ratings.Rating = &ratings.Rating{}

	apitest.New().Observe(cleanup).Handler(newApp()).Put("/api/v1/ratings/" + ratingId).Header("Authorization", token).JSON(ratingRequest).Expect(t).Status(http.StatusBadRequest).End()
}

func TestDeleteRating_Success(t *testing.T) {
	token := getJwtToken(t)

	rating := getRating()
	ratingId := strconv.Itoa(int(rating.ID))

	apitest.New().Observe(cleanup).Handler(newApp()).Delete("/api/v1/ratings/" + ratingId).Header("Authorization", token).Expect(t).Status(http.StatusOK).End()
}

func TestDeleteRating_Failed(t *testing.T) {
	token := getJwtToken(t)

	apitest.New().Handler(newApp()).Delete("/api/v1/ratings/0").Header("Authorization", token).Expect(t).Status(http.StatusNotFound).End()
}

func TestGetRatingByMovieID_Success(t *testing.T) {
	token := getJwtToken(t)

	rating := getRating()
	movieID := strconv.Itoa(int(rating.MovieID))

	apitest.New().Observe(cleanup).Handler(newApp()).Get("/api/v1/ratings/movie?").Query("movie_id", movieID).Header("Authorization", token).Expect(t).Status(http.StatusOK).End()
}

func TestGetRatingByUserID_Success(t *testing.T) {
	token := getJwtToken(t)

	rating := getRating()
	userID := strconv.Itoa(int(rating.UserID))

	apitest.New().Observe(cleanup).Handler(newApp()).Get("/api/v1/ratings/user?").Query("user_id", userID).Header("Authorization", token).Expect(t).Status(http.StatusOK).End()
}

func TestLogout_Success(t *testing.T) {
	var token string = getJwtToken(t)

	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Post("/api/v1/users/logout").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogout_Failed(t *testing.T) {
	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Post("/api/v1/users/logout").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}