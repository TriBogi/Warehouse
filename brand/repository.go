package brand

import "gorm.io/gorm"

type Repository interface {
	Save(saveBrand Brand) (Brand, error)
	GetBrandList() ([]Brand, error)
	Delete() ([]Brand, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}

}

func (r *repository) Save(saveBrand Brand) (Brand, error) {

	err := r.db.Create(&saveBrand).Error
	if err != nil {
		return saveBrand, err
	}

	return saveBrand, nil

}

func (r *repository) GetBrandList() ([]Brand, error) {
	var getBrand []Brand

	err := r.db.Find(&getBrand).Error
	if err != nil {
		return getBrand, err
	}
	return getBrand, nil
}

func (r *repository) Delete() ([]Brand, error) {
	var deleteBrand []Brand

	err := r.db.Delete(&deleteBrand).Error

	if err != nil {
		return deleteBrand, err
	}
	return deleteBrand, nil

}
