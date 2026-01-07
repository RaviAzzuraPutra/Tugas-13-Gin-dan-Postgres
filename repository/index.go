package repository

import (
	"formative-13/database"
	"formative-13/model"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRpositoryRegistry() *Repository {
	return &Repository{
		DB: database.DB,
	}
}

func (repo *Repository) Create(Bioskop *model.Bioskop) error {
	errCreate := repo.DB.Table("bioskops").Create(Bioskop).Error

	return errCreate
}

func (repo *Repository) Get() ([]model.Bioskop, error) {
	var Bioskop []model.Bioskop

	errGet := repo.DB.Table("bioskops").Find(&Bioskop).Error

	return Bioskop, errGet
}

func (repo *Repository) GetById(ID string) (*model.Bioskop, error) {
	var Bioskop *model.Bioskop

	errGet := repo.DB.Table("bioskops").Where("id = ?", ID).First(&Bioskop).Error

	return Bioskop, errGet
}

func (repo *Repository) Update(ID string, Bioskop *model.Bioskop) error {
	errUpdate := repo.DB.Table("bioskops").Where("id = ?", ID).Updates(Bioskop).Error

	return errUpdate
}

func (repo *Repository) Delete(ID string) error {

	var Bioskop *model.Bioskop

	errDelete := repo.DB.Table("bioskops").Unscoped().Where("id = ?", ID).Delete(&Bioskop).Error

	return errDelete
}
