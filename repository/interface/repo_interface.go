package repo_interface

import "formative-13/model"

type Repository_Interface interface {
	Create(Bioskop *model.Bioskop) error
	Get() ([]model.Bioskop, error)
	GetById(ID string) (*model.Bioskop, error)
	Update(ID string, Bioskop *model.Bioskop) error
	Delete(ID string) error
}
