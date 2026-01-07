package service

import (
	"formative-13/helper"
	"formative-13/model"
	repo_interface "formative-13/repository/interface"
	"formative-13/request"
)

type Service struct {
	repository repo_interface.Repository_Interface
}

func NewServiceRegistry(repository repo_interface.Repository_Interface) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(request *request.Bioskop_Request) (*model.Bioskop, error) {

	if request.Nama == nil || *request.Nama == "" {
		return nil, helper.NewBadRequest("Name cannot be empty!")
	}

	if request.Lokasi == nil || *request.Lokasi == "" {
		return nil, helper.NewBadRequest("Location cannot be empty!")
	}

	Bioskop := &model.Bioskop{
		Nama:   request.Nama,
		Lokasi: request.Lokasi,
		Rating: request.Rating,
	}

	errCreate := s.repository.Create(Bioskop)

	if errCreate != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Adding a Cinemas : " + errCreate.Error())
	}

	return Bioskop, nil
}

func (s *Service) Get() ([]model.Bioskop, error) {
	Bioskop, errGet := s.repository.Get()

	if len(Bioskop) == 0 {
		return nil, helper.NewNotFound("Cinema Data Not Found or Still Empty!")
	}

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Get Cinemas :" + errGet.Error())
	}

	return Bioskop, nil
}

func (s *Service) GetById(ID string) (*model.Bioskop, error) {
	Bioskop, errGet := s.repository.GetById(ID)

	if Bioskop == nil {
		return nil, helper.NewNotFound("Cinema Data Not Found or Still Empty!")
	}

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Get Cinemas :" + errGet.Error())
	}

	return Bioskop, nil
}

func (s *Service) Update(ID string, request *request.Bioskop_Request) (*model.Bioskop, error) {

	getBioskop, errGet := s.repository.GetById(ID)

	if errGet != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Get Cinemas :" + errGet.Error())
	}

	if getBioskop == nil {
		return nil, helper.NewNotFound("Cinema Data Not Found or Still Empty!")
	}

	if request.Nama == nil || *request.Nama == "" {
		return nil, helper.NewBadRequest("Name cannot be empty!")
	}

	if request.Lokasi == nil || *request.Lokasi == "" {
		return nil, helper.NewBadRequest("Location cannot be empty!")
	}

	getBioskop.Nama = request.Nama
	getBioskop.Lokasi = request.Lokasi
	getBioskop.Rating = request.Rating

	errUpdate := s.repository.Update(ID, getBioskop)

	if errUpdate != nil {
		return nil, helper.NewInternalServerError("An Error Occurred While Update Cinemas :" + errUpdate.Error())
	}

	return getBioskop, nil
}

func (s *Service) Delete(ID string) error {
	getBioskop, errGet := s.repository.GetById(ID)

	if errGet != nil {
		return helper.NewInternalServerError("An Error Occurred While Get Cinemas :" + errGet.Error())
	}

	if getBioskop == nil {
		return helper.NewNotFound("Cinema Data Not Found or Still Empty!")
	}

	errDelete := s.repository.Delete(ID)

	if errDelete != nil {
		helper.NewInternalServerError("An Error Occurred While Delete Cinemas :" + errDelete.Error())
	}

	return nil
}
