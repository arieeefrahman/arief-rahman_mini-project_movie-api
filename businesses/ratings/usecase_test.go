package ratings_test

import (
	"mini-project-movie-api/businesses/genres"
	"mini-project-movie-api/businesses/movies"
	"mini-project-movie-api/businesses/ratings"
	_ratingMock "mini-project-movie-api/businesses/ratings/mocks"
	"mini-project-movie-api/businesses/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

var (
	ratingRepository _ratingMock.Repository
	ratingService ratings.UseCase

	ratingDomain ratings.Domain
)

func TestMain(m *testing.M) {
	ratingService = ratings.NewRatingUseCase(&ratingRepository)
	
	genreDomain := genres.Domain{
		Name: "test genre",
	}

	date := "2005-01-01"
	releaseDate, _ := time.Parse("2006-01-02", date)

	movieDomain := movies.Domain{
		Title: "title",
		Synopsis: "title synopsis",
		GenreID: genreDomain.ID,
		ReleaseDate: datatypes.Date(releaseDate),
	}

	userDomain := users.Domain{
		Email:    "test@test.com",
		Password: "123123",
	}

	ratingDomain = ratings.Domain{
		Score: 10,
		MovieID: movieDomain.ID,
		UserID: userDomain.ID,
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		ratingRepository.On("GetAll").Return([]ratings.Domain{ratingDomain}).Once()

		result := ratingService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		ratingRepository.On("GetAll").Return([]ratings.Domain{}).Once()

		result := ratingService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		ratingRepository.On("GetByID", "1").Return(ratingDomain).Once()

		result := ratingService.GetByID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		ratingRepository.On("GetByID", "-1").Return(ratings.Domain{}).Once()

		result := ratingService.GetByID("-1")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		ratingRepository.On("Create", &ratingDomain).Return(ratingDomain).Once()

		result := ratingService.Create(&ratingDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		ratingRepository.On("Create", &ratings.Domain{}).Return(ratings.Domain{}).Once()

		result := ratingService.Create(&ratings.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		ratingRepository.On("Update", "1", &ratingDomain).Return(ratingDomain).Once()

		result := ratingService.Update("1", &ratingDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		ratingRepository.On("Update", "1", &ratings.Domain{}).Return(ratings.Domain{}).Once()

		result := ratingService.Update("1", &ratings.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		ratingRepository.On("Delete", "1").Return(true).Once()

		result := ratingService.Delete("1")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		ratingRepository.On("Delete", "-1").Return(false).Once()

		result := ratingService.Delete("-1")

		assert.False(t, result)
	})
}

func TestGetByMovieID(t *testing.T) {
	t.Run("Get By Movie ID | Valid", func(t *testing.T) {
		ratingRepository.On("GetByMovieID", "1").Return([]ratings.Domain{ratingDomain}).Once()

		result := ratingService.GetByMovieID("1")

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get By Movie ID | InValid", func(t *testing.T) {
		ratingRepository.On("GetByMovieID", "1").Return([]ratings.Domain{}).Once()

		result := ratingService.GetByMovieID("1")

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByUserID(t *testing.T) {
	t.Run("Get By User ID | Valid", func(t *testing.T) {
		ratingRepository.On("GetByUserID", "1").Return([]ratings.Domain{ratingDomain}).Once()

		result := ratingService.GetByUserID("1")

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get By User ID | InValid", func(t *testing.T) {
		ratingRepository.On("GetByUserID", "1").Return([]ratings.Domain{}).Once()

		result := ratingService.GetByUserID("1")

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByMovieIdAndUserID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		ratingRepository.On("GetByMovieIdAndUserID", "1", "1").Return(ratingDomain).Once()

		result := ratingService.GetByMovieIdAndUserID("1", "1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		ratingRepository.On("GetByID", "-1", "-1").Return(ratings.Domain{}).Once()

		result := ratingService.GetByMovieIdAndUserID("-1", "-1")

		assert.NotNil(t, result)
	})
}