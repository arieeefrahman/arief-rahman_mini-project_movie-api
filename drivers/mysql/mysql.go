package mysql_driver

import (
	"errors"
	"fmt"
	"log"

	"mini-project-movie-api/drivers/mysql/genres"
	"mini-project-movie-api/drivers/mysql/movies"
	"mini-project-movie-api/drivers/mysql/ratings"
	"mini-project-movie-api/drivers/mysql/users"

	_utils "mini-project-movie-api/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error
	loc := "Asia%2FJakarta"

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to database: %s", err)
	}

	log.Println("connected to database")

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&genres.Genre{}, &movies.Movie{}, &ratings.Rating{}, &users.User{})
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}

func SeedGenre(db *gorm.DB) genres.Genre {
	genre, _ := _utils.CreateFaker[genres.Genre]()

	if err := db.Create(&genre).Error; err != nil {
		panic(err)
	}

	db.Last(&genre)

	return genre
}

func SeedMovie(db *gorm.DB) movies.Movie {
	genre := SeedGenre(db)

	movie, _ := _utils.CreateFaker[movies.Movie]()

	movie.GenreID = genre.ID

	if err := db.Create(&movie).Error; err != nil {
		panic(err)
	}

	db.Last(&movie)

	return movie
}

func SeedRating(db *gorm.DB) ratings.Rating {
	movie := SeedMovie(db)

	rating, _ := _utils.CreateFaker[ratings.Rating]()

	rating.MovieID = movie.ID

	if err := db.Create(&rating).Error; err != nil {
		panic(err)
	}

	db.Last(&rating)

	return rating
}

func SeedUser(db *gorm.DB) users.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	fakeUser, _ := _utils.CreateFaker[users.User]()
	fakeEmail := fmt.Sprintf("%s@gmail.com", fakeUser.Email)
	userRecord := users.User{
		Email:    fakeEmail,
		Password: string(password),
	}

	if err := db.Create(&userRecord).Error; err != nil {
		panic(err)
	}

	var foundUser users.User

	db.Last(&foundUser)

	foundUser.Password = "123123"

	return foundUser
}

func CleanSeeders(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	genreResult := db.Exec("DELETE FROM genres")
	movieResult := db.Exec("DELETE FROM movies")
	ratingResult := db.Exec("DELETE FROM ratings")
	userResult := db.Exec("DELETE FROM users")

	var isFailed bool = movieResult.Error != nil || userResult.Error != nil || genreResult.Error != nil || ratingResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	log.Println("Seeders are cleaned up successfully")
}