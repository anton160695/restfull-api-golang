package service

import (
	"crud-golang/crud/src/database"
	"crud-golang/crud/src/model"
	"crud-golang/crud/src/repository"
	"errors"
)

type BookService interface {
	GetAllBook() ([]model.AllBooks, error)
	GetBookById(id int) (*model.BookDetails, error)
	CreateBook(userId int, req model.CreateBook) (*model.CreatBookRes, error)
	UpdateBook(id int, req model.UpdateBook) (*model.CreatBookRes, error)
	DeleteBook(id int) error
}

type bookService struct {
	userRepo repository.UserRepository
	bookRepo repository.BooksRepository
	catRepo  repository.CategoryBookRepo
}

func NewBookService(bookRepo repository.BooksRepository, userRepo repository.UserRepository, catRepo repository.CategoryBookRepo) BookService {
	return &bookService{
		bookRepo: bookRepo,
		userRepo: userRepo,
		catRepo:  catRepo,
	}
}

func (s *bookService) GetAllBook() ([]model.AllBooks, error) {
	books, err := s.bookRepo.FindAll()
	if err != nil {
		return nil, err
	}
	res := model.ToAllBookDetails(books)
	return res, nil
}

func (s *bookService) GetBookById(id int) (*model.BookDetails, error) {
	book, err := s.bookRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("book not found")
	}
	res := model.ToBooksDetails(*book)
	return &res, nil
}

func (s *bookService) CreateBook(userId int, req model.CreateBook) (*model.CreatBookRes, error) {
	cekUser, err := s.userRepo.FindUserByID(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}
	err = s.catRepo.FinByIdSimple(req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	book := &database.Book{
		Title:      req.Title,
		Exercpt:    req.Excerpt,
		Content:    req.Content,
		Creator:    cekUser.Name,
		CategoryID: req.CategoryID,
	}
	book, err = s.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}
	res := model.ToCreateRes(*book)
	return &res, nil
}

func (s *bookService) UpdateBook(id int, req model.UpdateBook) (*model.CreatBookRes, error) {
	_, err := s.bookRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("book not found")
	}
	update := &database.Book{
		Title:      req.Title,
		Exercpt:    req.Excerpt,
		Content:    req.Content,
		CategoryID: req.CategoryID,
	}
	book, err := s.bookRepo.Update(id, update)
	if err != nil {
		return nil, errors.New("failed to update book")
	}
	res := model.ToCreateRes(*book)
	return &res, nil
}

func (s *bookService) DeleteBook(id int) error {
	err := s.bookRepo.Delete(id)
	if err != nil {
		return errors.New("failed to delete book")
	}
	return nil
}
