package ratings

import (
	"mini-project-movie-api/app/middlewares"

	"mini-project-movie-api/businesses/ratings"
	ctrl "mini-project-movie-api/controllers"
	"mini-project-movie-api/controllers/ratings/request"
	"mini-project-movie-api/controllers/ratings/response"

	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RatingController struct {
	ratingUseCase ratings.UseCase
}

func NewRatingController(RatingUC ratings.UseCase) *RatingController {
	return &RatingController{
		ratingUseCase: RatingUC,
	}
}

func (rc *RatingController) GetAll(c echo.Context) error {
	ratings := []response.Rating{}

	ratingsData := rc.ratingUseCase.GetAll()

	for _, rating := range ratingsData {
		ratings = append(ratings, response.FromDomain(rating))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all ratings", ratings)
}

func (rc *RatingController) GetByID(c echo.Context) error {
	var id string = c.Param("id")
	
	rating := rc.ratingUseCase.GetByID(id)

	if rating.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "rating not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "rating found", response.FromDomain(rating))
}

func (rc *RatingController) Create(c echo.Context) error {
	input := request.Rating{}

	if err := c.Bind(&input); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	
	user := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "invalid token",
		})
	}

	getUserId := middlewares.GetUser(user).ID
	input.UserID = uint(getUserId)
	
	err := input.Validate()

	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	// handle if score input out of range [1-10]
	if input.Score < 1 || input.Score > 10 {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "score must be within 1 to 10 ", "")
	}
	
	movieId := strconv.FormatUint(uint64(input.MovieID), 10)
	userId := strconv.FormatUint(uint64(getUserId), 10)
	check := rc.ratingUseCase.GetByMovieIdAndUserID(movieId, userId)

	if check.ID != 0 {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "user already rated", "")
	}

	rating := rc.ratingUseCase.Create(input.ToDomain())

	if response.FromDomain(rating).ID == 0 {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	return ctrl.NewResponse(c, http.StatusCreated, "success", "rating created", response.FromDomain(rating))
}

func (rc *RatingController) Update(c echo.Context) error {
	var id string = c.Param("id")

	// handle if id does not exist
	ratingData := rc.ratingUseCase.GetByID(id) 
	
	if ratingData.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "rating not found", "")
	}
	
	input := request.Rating{}

	if err := c.Bind(&input); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	user := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "invalid token",
		})
	}

	getUserId := middlewares.GetUser(user).ID
	input.UserID = uint(getUserId)

	err := input.Validate()
	
	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	
	rating := rc.ratingUseCase.Update(id, input.ToDomain())

	return ctrl.NewResponse(c, http.StatusOK, "success", "rating updated", response.FromDomain(rating))
}

func (rc *RatingController) Delete(c echo.Context) error {
	ratingId := c.Param("id")

	// handle if id does not exist
	ratingData := rc.ratingUseCase.GetByID(ratingId) 
	
	if ratingData.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "rating not found", "")
	}

	user := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "invalid token",
		})
	}
	
	isSuccess := rc.ratingUseCase.Delete(ratingId)

	if !isSuccess {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "rating not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "rating deleted", "")
}

func (rc *RatingController) GetByMovieID(c echo.Context) error {
	var movieId string = c.QueryParam("movie_id")

	if movieId == "" {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "query parameters input do not match", "")
	}

	ratings := []response.Rating{}
	ratingsData := rc.ratingUseCase.GetByMovieID(movieId)

	for _, rating := range ratingsData {
		ratings = append(ratings, response.FromDomain(rating))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all ratings with movie id", ratings)
}

func (rc *RatingController) GetByUserID(c echo.Context) error {
	ratings := []response.Rating{}
	var userId string = c.QueryParam("user_id")

	if userId == "" {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "query parameters input do not match", "")
	}

	ratingsData := rc.ratingUseCase.GetByUserID(userId)

	for _, rating := range ratingsData {
		ratings = append(ratings, response.FromDomain(rating))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all ratings with user id", ratings)
}