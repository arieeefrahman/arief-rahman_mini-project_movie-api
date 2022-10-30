package genres

type genreUseCase struct {
	genreRepository Repository
}

func NewGenreUsecase(gr Repository) UseCase {
	return &genreUseCase{
		genreRepository: gr,
	}
}

func (gu *genreUseCase) GetAll() []Domain {
	return gu.genreRepository.GetAll()
}

func (gu *genreUseCase) GetByID(id string) Domain {
	return gu.genreRepository.GetByID(id)
}

func (gu *genreUseCase) Create(genreDomain *Domain) Domain {
	return gu.genreRepository.Create(genreDomain)
}

func (gu *genreUseCase) Update(id string, genreDomain *Domain) Domain {
	return gu.genreRepository.Update(id, genreDomain)
}

func (gu *genreUseCase) Delete(id string) bool {
	return gu.genreRepository.Delete(id)
}

func (gu *genreUseCase) GetByName(name string) Domain {
	return gu.genreRepository.GetByName(name)
}