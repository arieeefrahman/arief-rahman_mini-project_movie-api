package routes

import (
	"mini-project-movie-api/controllers/genres"
	"mini-project-movie-api/controllers/movies"
	"mini-project-movie-api/controllers/ratings"
	"mini-project-movie-api/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware	echo.MiddlewareFunc
	JWTMiddleware		middleware.JWTConfig
	AuthController		users.AuthController
	GenreController		genres.GenreController
	MovieController		movies.MovieController
	RatingController	ratings.RatingController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/api/v1/users")
	users.POST("/signup", cl.AuthController.Signup).Name = "user-signup"
	users.POST("/login", cl.AuthController.Login).Name = "user-login"

	genres := e.Group("/api/v1/genres", middleware.JWTWithConfig(cl.JWTMiddleware))
	genres.GET("", cl.GenreController.GetAll).Name = "get-all-genres"
	genres.GET("/:id", cl.GenreController.GetByID).Name = "get-genre-by-id"
	genres.POST("", cl.GenreController.Create).Name = "create-genre"
	genres.PUT("/:id", cl.GenreController.Update).Name = "update-genre"
	genres.DELETE("/:id", cl.GenreController.Delete).Name = "delete-genre"

	movies := e.Group("/api/v1/movies", middleware.JWTWithConfig(cl.JWTMiddleware))
	movies.GET("", cl.MovieController.GetAll).Name = "get-all-movies"
	movies.GET("/:id", cl.MovieController.GetByID).Name = "get-movie-by-id"
	movies.POST("", cl.MovieController.Create).Name = "create-movie"
	movies.PUT("/:id", cl.MovieController.Update).Name = "update-movie"
	movies.DELETE("/:id", cl.MovieController.Delete).Name = "delete-movie"
	movies.GET("/genre/:genre_id", cl.MovieController.GetByGenreID).Name = "get-movies-by-genre-id"
	movies.GET("/latest", cl.MovieController.GetLatest).Name = "get-latest-movie"
	movies.GET("/title", cl.MovieController.GetByTitle).Name = "get-movie-by-title"

	ratings := e.Group("/api/v1/ratings", middleware.JWTWithConfig(cl.JWTMiddleware))
	ratings.GET("", cl.RatingController.GetAll).Name = "get-all-ratings"
	ratings.GET("/:id", cl.RatingController.GetByID).Name = "get-rating-by-id"
	ratings.POST("", cl.RatingController.Create).Name = "create-rating"
	ratings.PUT("/:id", cl.RatingController.Update).Name = "update-rating"
	ratings.DELETE("/:id", cl.RatingController.Delete).Name = "delete-rating"
	ratings.GET("/movie", cl.RatingController.GetByMovieID).Name = "get-ratings-by-movie-id"
	ratings.GET("/user", cl.RatingController.GetByUserID).Name = "get-ratings-by-user-id"

	auth := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	auth.POST("/logout", cl.AuthController.Logout).Name = "user-logout"
}