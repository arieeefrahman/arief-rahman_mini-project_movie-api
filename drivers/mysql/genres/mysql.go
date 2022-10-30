package genres

import (
	"mini-project-movie-api/businesses/genres"

	"gorm.io/gorm"
)

type genreRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) genres.Repository {
	return &genreRepository{
		conn: conn,
	}
}

func (gr *genreRepository) GetAll() []genres.Domain {
	var rec []Genre
	gr.conn.Find(&rec)

	genreDomain := []genres.Domain{}

	for _, genre := range rec {
		genreDomain = append(genreDomain, genre.ToDomain())
	}

	return genreDomain
}

func (gr *genreRepository) GetByID(id string) genres.Domain {
	var genre Genre
	gr.conn.First(&genre, "id = ?", id)

	return genre.ToDomain()
}

func (gr *genreRepository) Create(genreDomain *genres.Domain) genres.Domain {
	rec := FromDomain(genreDomain)
	result := gr.conn.Create(&rec)
	result.Last(&rec)

	return rec.ToDomain()
}

func (gr *genreRepository) Update(id string, genreDomain *genres.Domain) genres.Domain {
	var genre genres.Domain = gr.GetByID(id)

	updatedGenre := FromDomain(&genre)
	updatedGenre.Name = genreDomain.Name

	gr.conn.Save(&updatedGenre)

	return updatedGenre.ToDomain()
}

func (gr *genreRepository) Delete(id string) bool {
	var genre genres.Domain = gr.GetByID(id)
	
	deletedCategory := FromDomain(&genre)
	result := gr.conn.Delete(&deletedCategory)

	return result.RowsAffected != 0
}

func (gr *genreRepository) GetByName(name string) genres.Domain {
	var rec Genre
	gr.conn.First(&rec, "name = ?", name)

	return rec.ToDomain()
} 