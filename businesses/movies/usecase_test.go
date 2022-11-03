package movies_test

import (
	"mini-project-movie-api/businesses/genres"
	"mini-project-movie-api/businesses/movies"
	_movieMock "mini-project-movie-api/businesses/movies/mocks"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	movieRepository _movieMock.Repository
	movieService movies.UseCase

	movieDomain movies.Domain
)

func TestMain(m *testing.M) {
	movieService = movies.NewMovieUseCase(&movieRepository)

	genreDomain := genres.Domain{
		Name: "test genre",
	}

	date := "2005-01-01"
	releaseDate, _ := time.Parse("2006-01-02", date)

	movieDomain = movies.Domain{
		Title: "title",
		Synopsis: "title synopsis",
		GenreID: genreDomain.ID,
		ReleaseDate: releaseDate,
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		movieRepository.On("GetAll").Return([]movies.Domain{movieDomain}).Once()

		result := movieService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		movieRepository.On("GetAll").Return([]movies.Domain{}).Once()

		result := movieService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		movieRepository.On("GetByID", "1").Return(movieDomain).Once()

		result := movieService.GetByID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		movieRepository.On("GetByID", "-1").Return(movies.Domain{}).Once()

		result := movieService.GetByID("-1")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		movieRepository.On("Create", &movieDomain).Return(movieDomain).Once()

		result := movieService.Create(&movieDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		movieRepository.On("Create", &movies.Domain{}).Return(movies.Domain{}).Once()

		result := movieService.Create(&movies.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		movieRepository.On("Update", "1", &movieDomain).Return(movieDomain).Once()

		result := movieService.Update("1", &movieDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		movieRepository.On("Update", "1", &movies.Domain{}).Return(movies.Domain{}).Once()

		result := movieService.Update("1", &movies.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		movieRepository.On("Delete", "1").Return(true).Once()

		result := movieService.Delete("1")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		movieRepository.On("Delete", "-1").Return(false).Once()

		result := movieService.Delete("-1")

		assert.False(t, result)
	})
}

func TestGetByGenreID(t *testing.T) {
	t.Run("Get By Genre ID | Valid", func(t *testing.T) {
		movieRepository.On("GetByGenreID", "1").Return([]movies.Domain{movieDomain}).Once()

		result := movieService.GetByGenreID("1")

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get By Genre ID | InValid", func(t *testing.T) {
		movieRepository.On("GetByGenreID", "1").Return([]movies.Domain{}).Once()

		result := movieService.GetByGenreID("1")

		assert.Equal(t, 0, len(result))
	})
}

func TestGetLatest(t *testing.T) {
	t.Run("Get Latest | Valid", func(t *testing.T) {
		movieRepository.On("GetLatest").Return([]movies.Domain{movieDomain}).Once()

		result := movieService.GetLatest()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get Latest | InValid", func(t *testing.T) {
		movieRepository.On("GetLatest").Return([]movies.Domain{}).Once()

		result := movieService.GetLatest()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByTitle(t *testing.T) {
	t.Run("Get By Title | Valid", func(t *testing.T) {
		movieRepository.On("GetByTitle", "title").Return(movieDomain).Once()

		result := movieService.GetByTitle("title")

		assert.NotNil(t, result)
	})

	t.Run("Get By Title | InValid", func(t *testing.T) {
		movieRepository.On("GetByTitle", ";;;;").Return(movies.Domain{}).Once()

		result := movieService.GetByTitle(";;;;")

		assert.NotNil(t, result)
	})
}