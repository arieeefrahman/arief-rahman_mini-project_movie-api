package main

import (
	"fmt"
	_middlewares "mini-project-movie-api/app/middlewares"
	_routes "mini-project-movie-api/app/routes"
	_util "mini-project-movie-api/utils"

	_driverFactory "mini-project-movie-api/drivers"
	_dbDriver "mini-project-movie-api/drivers/mysql"

	_genreUseCase "mini-project-movie-api/businesses/genres"
	_genreController "mini-project-movie-api/controllers/genres"

	_movieUseCase "mini-project-movie-api/businesses/movies"
	_movieController "mini-project-movie-api/controllers/movies"

	_ratingUseCase "mini-project-movie-api/businesses/ratings"
	_ratingController "mini-project-movie-api/controllers/ratings"

	_userUseCase "mini-project-movie-api/businesses/users"
	_userController "mini-project-movie-api/controllers/users"

	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "3000"

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: _util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: _util.GetConfig("DB_PASSWORD"),
		DB_HOST: _util.GetConfig("DB_TEST_HOST"),
		DB_PORT: _util.GetConfig("DB_PORT"),
		DB_NAME: _util.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT: _util.GetConfig("JWT_SECRET_KEY"),
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

	var appPort string = fmt.Sprintf(":%s", DEFAULT_PORT)

	app.Logger.Fatal(app.Start(appPort))
}