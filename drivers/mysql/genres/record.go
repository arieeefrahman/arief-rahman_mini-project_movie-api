package genres

import (
	"mini-project-movie-api/businesses/genres"
	"time"

	"gorm.io/gorm"
)

type Genre struct {
	ID        uint				`json:"id" gorm:"primaryKey"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
	Name      string			`json:"name" gorm:"unique" faker:"name"`
}

func FromDomain(domain *genres.Domain) *Genre {
	return &Genre{
		ID: domain.ID,
		Name: domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *Genre) ToDomain() genres.Domain {
	return genres.Domain{
		ID: rec.ID,
		Name: rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}