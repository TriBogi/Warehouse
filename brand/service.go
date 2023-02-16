package brand

import (
	"time"
)

type Service interface {
	RegisterBrand(input CreateBrandInput) (Brand, error)
	GetAllBrands() ([]Brand, error)
	Delete(inputID GetBrandDetailInput) error
	UpdateBrand(inputID GetBrandDetailInput, inputData CreateBrandInput) (Brand, error)
	GetBrandByID(input GetBrandDetailInput) (Brand, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterBrand(input CreateBrandInput) (Brand, error) {
	inputBrand := Brand{
		Brd_Name:  input.BrdName,
		CreatedAt: time.Now(),
	}
	res, err := s.repository.Save(inputBrand)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *service) GetAllBrands() ([]Brand, error) {
	res, err := s.repository.GetAll()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *service) Delete(inputID GetBrandDetailInput) error {
	err := s.repository.Delete(inputID.BrdId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateBrand(inputID GetBrandDetailInput, inputData CreateBrandInput) (Brand, error) {
	brand, err := s.repository.FindByID(inputID.BrdId)
	if err != nil {
		return brand, err
	}

	brand.Brd_Name = inputData.BrdName
	brand.UpdatedAt = time.Now()

	updatedBrand, err := s.repository.Update(brand)
	if err != nil {
		return updatedBrand, err
	}
	return updatedBrand, nil
}

func (s *service) GetBrandByID(inputID GetBrandDetailInput) (Brand, error) {
	brand, err := s.repository.FindByID(inputID.BrdId)
	if err != nil {
		return brand, err
	}

	return brand, nil
}
