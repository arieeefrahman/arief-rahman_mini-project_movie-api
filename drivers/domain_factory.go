package drivers

import (
	genreDomain "mini-project-movie-api/businesses/genres"
	movieDomain "mini-project-movie-api/businesses/movies"
	ratingDomain "mini-project-movie-api/businesses/ratings"
	userDomain "mini-project-movie-api/businesses/users"

	genreDB "mini-project-movie-api/drivers/mysql/genres"
	movieDB "mini-project-movie-api/drivers/mysql/movies"
	ratingDB "mini-project-movie-api/drivers/mysql/ratings"
	userDB "mini-project-movie-api/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewGenreRepository(conn *gorm.DB) genreDomain.Repository {
	return genreDB.NewMySQLRepository(conn)
}

func NewMovieRepository(conn *gorm.DB) movieDomain.Repository {
	return movieDB.NewMySQLRepository(conn)
}

func NewRatingRepository(conn *gorm.DB) ratingDomain.Repository {
	return ratingDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}