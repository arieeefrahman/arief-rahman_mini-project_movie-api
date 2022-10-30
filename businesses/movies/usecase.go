package movies

type movieUseCase struct {
	movieRepository Repository
}

func NewMovieUseCase(mr Repository) UseCase {
	return &movieUseCase{
		movieRepository: mr,
	}
}

func (mu *movieUseCase) GetAll() []Domain {
	return mu.movieRepository.GetAll()
}

func (mu *movieUseCase) GetByID(id string) Domain {
	return mu.movieRepository.GetByID(id)
}

func (mu *movieUseCase) Create(movieDomain *Domain) Domain {
	return mu.movieRepository.Create(movieDomain)
}

func (mu *movieUseCase) Update(id string, movieDomain *Domain) Domain {
	return mu.movieRepository.Update(id, movieDomain)
}

func (mu *movieUseCase) Delete(id string) bool {
	return mu.movieRepository.Delete(id)
}

func (mu *movieUseCase) GetByGenreID(genreId string) []Domain {
	return mu.movieRepository.GetByGenreID(genreId)
}

func (mu *movieUseCase) GetLatest() []Domain {
	return mu.movieRepository.GetLatest()
}

func (mu *movieUseCase) GetByTitle(title string) Domain {
	return mu.movieRepository.GetByTitle(title)
}