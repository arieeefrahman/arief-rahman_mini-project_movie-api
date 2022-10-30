package response

import (
	"mini-project-movie-api/businesses/ratings"
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	ID			uint			`json:"id" gorm:"primaryKey"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`json:"deleted_at"`
	Score		int				`json:"score"`
	MovieTitle	string			`json:"movie_title"`
	MovieID		uint			`json:"movie_id"`
	UserID		uint			`json:"user_id"`
}

func FromDomain(domain ratings.Domain) Rating {
	return Rating{
		ID: domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Score: domain.Score,
		MovieTitle: domain.MovieTitle,
		MovieID: domain.MovieID,
		UserID: domain.UserID,
	}
}