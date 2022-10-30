package users

import (
	"mini-project-movie-api/app/middlewares"
	"mini-project-movie-api/businesses/users"
	"mini-project-movie-api/controllers/users/request"
	"mini-project-movie-api/controllers/users/response"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase users.UseCase
}

func NewAuthController(authUC users.UseCase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ac *AuthController) Signup(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message" : "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	user := ac.authUseCase.Signup(userInput.ToDomain())

	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "email already taken",
		})
	}

	return c.JSON(http.StatusCreated, response.FromDomain(user))
}

func (ac *AuthController) Login(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	token := ac.authUseCase.Login(userInput.ToDomain())

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (ac *AuthController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	isListed := middlewares.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "invalid token",
		})
	}

	middlewares.Logout(user.Raw)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "logout success",
	})
}