package service

import (
	"crud-golang/crud/src/database"
	"crud-golang/crud/src/model"
	"crud-golang/crud/src/repository"
	"errors"
	"log"
)

type CategoryBookService interface {
	GetAll() ([]model.CategoryRes, error)
	GetById(id int) (*model.CategoryBookById, error)
	Create(req model.CategoryReq) (*model.CategoryRes, error)
	Update(id int, req model.CategoryReq) (*model.CategoryRes, error)
	Delete(id int) error
}

type categoryBookService struct {
	repo repository.CategoryBookRepo
}

func NewCategoryBookService(repo repository.CategoryBookRepo) CategoryBookService {
	return &categoryBookService{repo}
}

func (s *categoryBookService) GetAll() ([]model.CategoryRes, error) {
	cat, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	res := model.ToAllCategory(cat)

	return res, err
}

func (s *categoryBookService) GetById(id int) (*model.CategoryBookById, error) {
	cat, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("category not found")
	}
	res := model.ToCategoryBookById(cat)
	return &res, err
}

func (s *categoryBookService) Create(req model.CategoryReq) (*model.CategoryRes, error) {
	category := &database.CategoryBook{
		Name: req.Name,
	}
	cat, err := s.repo.Create(category)
	if err != nil {
		return nil, err
	}
	res := model.CategoryRes{
		Id:   cat.Id,
		Name: cat.Name,
	}
	return &res, nil
}

func (s *categoryBookService) Update(id int, req model.CategoryReq) (*model.CategoryRes, error) {
	// cek category
	_, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("category not found")
	}
	category := &database.CategoryBook{
		Name: req.Name,
	}
	cat, err := s.repo.Update(id, category)
	if err != nil {
		return nil, err
	}
	log.Println("cat: ", cat)
	res := model.CategoryRes{
		Id:   id,
		Name: cat.Name,
	}
	return &res, nil
}

func (s *categoryBookService) Delete(id int) error {
	// cek category
	_, err := s.repo.FindById(id)
	if err != nil {
		return errors.New("category not found")
	}
	err = s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
