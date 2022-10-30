package response

import (
	"mini-project-movie-api/businesses/genres"
	"time"

	"gorm.io/gorm"
)

type Genre struct {
	ID        uint			 `json:"id" gorm:"primaryKey;unique"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time		 `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string		 `json:"name" gorm:"unique"`
}

func FromDomain(domain *genres.Domain) Genre {
	return Genre{
		ID: domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name: domain.Name,
	}
}