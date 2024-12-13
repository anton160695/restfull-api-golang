package repository

import (
	"crud-golang/crud/src/database"

	"gorm.io/gorm"
)

type BooksRepository interface {
	FindAll() ([]database.Book, error)
	FindByID(id int) (*database.Book, error)
	Create(book *database.Book) (*database.Book, error)
	Update(id int, book *database.Book) (*database.Book, error)
	Delete(id int) error
}

type booksRepo struct {
	db *gorm.DB
}

func NewBooksRepository(db *gorm.DB) BooksRepository {
	return &booksRepo{db}
}

func (r *booksRepo) FindAll() ([]database.Book, error) {
	var books []database.Book
	err := r.db.Preload("Category").Select("id, title, exercpt, category_id, creator").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *booksRepo) FindByID(id int) (*database.Book, error) {
	var book database.Book
	err := r.db.Preload("Category").Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *booksRepo) Create(book *database.Book) (*database.Book, error) {
	err := r.db.Create(book).Error
	if err != nil {
		return nil, err
	}
	// preload category after create books
	if err = r.db.Preload("Category").First(book, book.Id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *booksRepo) Update(id int, book *database.Book) (*database.Book, error) {
	err := r.db.Where("id = ?", id).Updates(book).Error
	if err != nil {
		return nil, err

	}
	return book, nil
}

func (r *booksRepo) Delete(id int) error {
	var book database.Book
	err := r.db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		return err
	}
	return err
}
