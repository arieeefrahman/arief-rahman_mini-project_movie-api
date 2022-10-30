package movies

import (
	"mini-project-movie-api/businesses/movies"

	"gorm.io/gorm"

)

type movieRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) movies.Repository {
	return &movieRepository{
		conn: conn,
	}
}

func (mr *movieRepository) GetAll() []movies.Domain {
	var rec []Movie
	mr.conn.Preload("Genre").Order("title").Find(&rec)

	movieDomain := []movies.Domain{}

	for _, movie := range rec {
		movieDomain = append(movieDomain, movie.ToDomain())
	}

	return movieDomain
}

func (mr *movieRepository) GetByID(id string) movies.Domain {
	var movie Movie
	
	mr.conn.Preload("Genre").First(&movie, "id = ?", id)

	return movie.ToDomain()
}

func (mr *movieRepository) Create(movieDomain *movies.Domain) movies.Domain {
	rec := FromDomain(movieDomain)
	
	result := mr.conn.Preload("Genre").Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (mr *movieRepository) Update(id string, movieDomain *movies.Domain) movies.Domain {
	var movie movies.Domain = mr.GetByID(id)

	updatedMovie := FromDomain(&movie)
	updatedMovie.Title = movieDomain.Title
	updatedMovie.Synopsis = movie.Synopsis
	updatedMovie.GenreID = movie.GenreID
	updatedMovie.ReleaseDate = movie.ReleaseDate
	updatedMovie.RatingScore = movie.RatingScore

	mr.conn.Save(&updatedMovie)

	return updatedMovie.ToDomain()
}

func (mr *movieRepository) Delete(id string) bool {
	var movie movies.Domain = mr.GetByID(id)

	deletedMovie := FromDomain(&movie)
	result := mr.conn.Delete(&deletedMovie)

	return result.RowsAffected != 0
}

func (mr *movieRepository) GetByGenreID(genreId string) []movies.Domain {
	var rec []Movie
	mr.conn.Preload("Genre").Where("genre_id = ?", genreId).Find(&rec)

	movieDomain := []movies.Domain{}

	for _, movie := range rec {
		movieDomain = append(movieDomain, movie.ToDomain())
	}

	return movieDomain
}

func (mr *movieRepository) GetLatest() []movies.Domain {
	var rec []Movie
	mr.conn.Limit(10).Preload("Genre").Order("release_date desc, title").Find(&rec)

	movieDomain := []movies.Domain{}

	for _, movie := range rec {
		movieDomain = append(movieDomain, movie.ToDomain())
	}

	return movieDomain
}

func (mr *movieRepository) GetByTitle(title string) movies.Domain {
	var rec Movie
	mr.conn.Preload("Genre").First(&rec, "title = ?", title)

	return rec.ToDomain()
}