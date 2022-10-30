package ratings

import (
	ratingUseCase "mini-project-movie-api/businesses/ratings"
	"mini-project-movie-api/drivers/mysql/movies"
	"mini-project-movie-api/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Score 	  int			 `json:"score"`
	Movie     movies.Movie	 `json:"movie"`
	MovieID   uint			 `json:"movie_id"`
	UserID    uint			 `json:"user_id"`
	User      users.User	 `json:"user"`
}

func FromDomain(domain *ratingUseCase.Domain) *Rating {
	return &Rating{
		ID: domain.ID,
		Score: domain.Score,
		MovieID: domain.MovieID,
		UserID: domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *Rating) ToDomain() ratingUseCase.Domain {
	return ratingUseCase.Domain{
		ID: rec.ID,
		Score: rec.Score,
		MovieTitle: rec.Movie.Title,
		MovieID: rec.MovieID,
		UserID: rec.UserID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}