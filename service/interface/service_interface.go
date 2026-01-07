package service_interface

import (
	"formative-13/model"
	"formative-13/request"
)

type Service_Interface interface {
	Create(request *request.Bioskop_Request) (*model.Bioskop, error)
	Get() ([]model.Bioskop, error)
	GetById(ID string) (*model.Bioskop, error)
	Update(ID string, request *request.Bioskop_Request) (*model.Bioskop, error)
	Delete(ID string) error
}
