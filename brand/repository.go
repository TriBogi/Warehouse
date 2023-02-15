package brand

import (
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	Save(saveBrand Brand) (Brand, error)
	GetAll() ([]Brand, error)
	Delete(ID int) error
	FindByID(ID int) (Brand, error)
	Update(brand Brand) (Brand, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}

}

func (r *repository) Save(saveBrand Brand) (Brand, error) {

	err := r.db.Table("brand").Create(&saveBrand).Error
	if err != nil {
		return saveBrand, err
	}

	return saveBrand, nil

}

func (r *repository) GetAll() ([]Brand, error) {
	var brands []Brand

	err := r.db.Table("brand").Find(&brands).Error
	if err != nil {
		log.Fatalln(err)
	}
	return brands, nil
}

func (r *repository) Delete(ID int) error {
	var brand Brand

	err := r.db.Table("brand").Delete(&brand, ID).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindByID(ID int) (Brand, error) {
	var brand Brand

	err := r.db.Table("brand").Where("brd_id = ?", ID).Find(&brand).Error

	if err != nil {
		return brand, err
	}
	return brand, nil
}

func (r *repository) Update(brand Brand) (Brand, error) {
	err := r.db.Table("brand").Save(&brand).Error

	if err != nil {
		return brand, err
	}
	return brand, nil
}
