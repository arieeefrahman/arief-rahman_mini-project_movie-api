package ratings

import (
	"mini-project-movie-api/businesses/ratings"

	"gorm.io/gorm"
)

type ratingRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) ratings.Repository {
	return &ratingRepository{
		conn: conn,
	}
}

func (rr *ratingRepository) GetAll() []ratings.Domain {
	var rec []Rating
	rr.conn.Preload("Movie").Preload("User").Find(&rec)

	ratingDomain := []ratings.Domain{}

	for _, rating := range rec {
		ratingDomain = append(ratingDomain, rating.ToDomain())
	}

	return ratingDomain
}

func (rr *ratingRepository) GetByID(id string) ratings.Domain {
	var rating Rating
	
	rr.conn.First(&rating, "id = ?", id)

	return rating.ToDomain()
}

func (rr *ratingRepository) Create(ratingDomain *ratings.Domain) ratings.Domain {
	rec := FromDomain(ratingDomain)
	
	result := rr.conn.Preload("Movie").Preload("User").Create(&rec)

	result.Last(&rec)
	
	var avg float64
	movieId := rec.MovieID
	
	row := rr.conn.Table("ratings").Where("movie_id = ?", movieId).Select("AVG(score)").Row()
	row.Scan(&avg)

	// update rating_score to table `movies`
	rr.conn.Table("movies").Where("id = ?", movieId).Update("rating_score", &avg)

	return rec.ToDomain()
}

func (rr *ratingRepository) Update(id string, ratingDomain *ratings.Domain) ratings.Domain {
	var rating ratings.Domain = rr.GetByID(id)

	updatedRating := FromDomain(&rating)
	updatedRating.Score = ratingDomain.Score
	
	rr.conn.Save(&updatedRating)

	var avg float64
	movieId := updatedRating.MovieID
	
	row := rr.conn.Table("ratings").Where("movie_id = ?", movieId).Where("deleted_at IS NULL").Select("AVG(score)").Row()
	row.Scan(&avg)

	// update rating_score to table `movies`
	rr.conn.Table("movies").Where("id = ?", movieId).Updates(map[string]interface{}{"rating_score": &avg, "updated_at": updatedRating.UpdatedAt})

	return updatedRating.ToDomain()
}

func (rr *ratingRepository) Delete(id string) bool {
	var rating ratings.Domain = rr.GetByID(id)

	deletedRating := FromDomain(&rating)
	result := rr.conn.Delete(&deletedRating)

	var avg float64
	movieId := deletedRating.MovieID
	
	row := rr.conn.Table("ratings").Where("movie_id = ?", movieId).Where("deleted_at IS NULL").Select("AVG(score)").Row()
	row.Scan(&avg)

	// update rating_score to table `movies`
	rr.conn.Table("movies").Where("id = ?", movieId).Updates(map[string]interface{}{"rating_score": &avg, "updated_at": deletedRating.DeletedAt})

	return result.RowsAffected != 0
}

func (rr *ratingRepository) GetByMovieID(movieId string) []ratings.Domain {
	var rec []Rating

	rr.conn.Preload("Movie").Preload("User").Find(&rec, "movie_id = ?", movieId)

	ratingDomain := []ratings.Domain{}

	for _, rating := range rec {
		ratingDomain = append(ratingDomain, rating.ToDomain())
	}

	return ratingDomain
}

func (rr *ratingRepository) GetByUserID(userId string) []ratings.Domain {
	var rec []Rating

	rr.conn.Preload("User").Find(&rec, "user_id = ?", userId)

	ratingDomain := []ratings.Domain{}

	for _, rating := range rec {
		ratingDomain = append(ratingDomain, rating.ToDomain())
	}

	return ratingDomain
}

func (rr *ratingRepository) GetByMovieIdAndUserID(movieId string, userId string) ratings.Domain {
	var rating Rating

	rr.conn.Preload("Movie").Preload("User").First(&rating, "movie_id = ? AND user_id = ?", movieId, userId)

	return rating.ToDomain()
}

func (rr *ratingRepository) GetAvgScore(movieId string) float64 {
	var avg float64

	row := rr.conn.Table("ratings").Where("movie_id = ?", movieId).Select("AVG(score)").Row()
	
	row.Scan(&avg)

	return avg
}