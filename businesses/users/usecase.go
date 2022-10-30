package users

import "mini-project-movie-api/app/middlewares"

type UserUseCase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) UseCase {
	return &UserUseCase{
		userRepository: ur,
		jwtAuth: jwtAuth,
	}	
}

func (uu *UserUseCase) Signup(userDomain *Domain) Domain {
	return uu.userRepository.Signup(userDomain)
}

func (uu *UserUseCase) Login(userDomain *Domain) string {
	user := uu.userRepository.GetByEmail(userDomain)

	if user.ID == 0 {
		return ""
	}

	token := uu.jwtAuth.GenerateToken(int(user.ID))

	return token
}