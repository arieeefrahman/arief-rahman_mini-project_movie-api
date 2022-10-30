package ratings

type ratingUseCase struct {
	ratingRepository Repository
}

func NewRatingUseCase(ru Repository) UseCase {
	return &ratingUseCase{
		ratingRepository: ru,
	}
}

func (ru *ratingUseCase) GetAll() []Domain {
	return ru.ratingRepository.GetAll()
}

func (ru *ratingUseCase) Create(ratingDomain *Domain) Domain {
	return ru.ratingRepository.Create(ratingDomain)
}

func (ru *ratingUseCase) Update(id string, ratingDomain *Domain) Domain {
	return ru.ratingRepository.Update(id, ratingDomain)
}

func (ru *ratingUseCase) Delete(id string) bool {
	return ru.ratingRepository.Delete(id)
}

func (ru *ratingUseCase) GetByID(id string) Domain {
	return ru.ratingRepository.GetByID(id)
}

func (ru *ratingUseCase) GetByMovieID(movieId string) []Domain {
	return ru.ratingRepository.GetByMovieID(movieId)
}

func (ru *ratingUseCase) GetByUserID(userId string) []Domain {
	return ru.ratingRepository.GetByUserID(userId)
}

func (ru *ratingUseCase) GetByMovieIdAndUserID(movieId string, userId string) Domain {
	return ru.ratingRepository.GetByMovieIdAndUserID(movieId, userId)
}

func (ru *ratingUseCase) GetAvgScore(movieId string) float64 {
	return ru.ratingRepository.GetAvgScore(movieId)
}