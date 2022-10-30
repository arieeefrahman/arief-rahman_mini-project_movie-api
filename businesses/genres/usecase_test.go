package genres_test

import (
	"mini-project-movie-api/businesses/genres"
	_genreMocks "mini-project-movie-api/businesses/genres/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	genreRepository	_genreMocks.Repository
	genreService	genres.UseCase

	genreDomain 	genres.Domain
)

func TestMain(m *testing.M) {
	genreService = genres.NewGenreUsecase(&genreRepository)
	genreDomain = genres.Domain{
		Name: "test",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		genreRepository.On("GetAll").Return([]genres.Domain{genreDomain}).Once()

		result := genreService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		genreRepository.On("GetAll").Return([]genres.Domain{}).Once()

		result := genreService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		genreRepository.On("GetByID", "1").Return(genreDomain).Once()
		
		result := genreService.GetByID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		genreRepository.On("GetByID", "-1").Return(genres.Domain{}).Once()
		
		result := genreService.GetByID("-1")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		genreRepository.On("Create", &genreDomain).Return(genreDomain).Once()

		result := genreService.Create(&genreDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		genreRepository.On("Create", &genres.Domain{}).Return(genres.Domain{}).Once()

		result := genreService.Create(&genres.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		genreRepository.On("Update", "1", &genreDomain).Return(genreDomain).Once()

		result := genreService.Update("1", &genreDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		genreRepository.On("Update", "1", &genres.Domain{}).Return(genres.Domain{}).Once()

		result := genreService.Update("1", &genres.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		genreRepository.On("Delete", "1").Return(true).Once()

		result := genreService.Delete("1")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		genreRepository.On("Delete", "-1").Return(false).Once()

		result := genreService.Delete("-1")

		assert.False(t, result)
	})
}

func TestGetByName(t *testing.T) {
	t.Run("Get By Name | Valid", func(t *testing.T) {
		genreRepository.On("GetByName", "test").Return(genreDomain).Once()
		
		result := genreService.GetByName("test")

		assert.NotNil(t, result)
	})

	t.Run("Get By Name | InValid", func(t *testing.T) {
		genreRepository.On("GetByName", ";;;;").Return(genres.Domain{}).Once()
		
		result := genreService.GetByName(";;;;")

		assert.NotNil(t, result)
	})
}