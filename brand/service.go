package brand

type Service interface {
	RegisterBrand(input BrandInput) (Brand, error)
	GetBrand() ([]Brand, error)
	DeleteBrand() ([]Brand, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}

}

func (s *service) RegisterBrand(input BrandInput) (Brand, error) {
	inputBrand := Brand{}
	inputBrand.brd_name = input.brd_name

	regBrand, err := s.repository.Save(inputBrand)
	if err != nil {
		return regBrand, err
	}

	return regBrand, nil

}

func (s *service) GetBrand() ([]Brand, error) {
	getBrand, err := s.repository.GetBrandList()
	if err != nil {
		return getBrand, err
	}
	return getBrand, nil

}

func (s *service) DeleteBrand() ([]Brand, error) {
	delBrand, err := s.repository.Delete()
	if err != nil {
		return delBrand, err
	}
	return delBrand, nil

}
