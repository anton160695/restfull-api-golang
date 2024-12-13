package repository

import (
	"crud-golang/crud/src/database"

	"gorm.io/gorm"
)

type CategoryBookRepo interface {
	FindAll() ([]database.CategoryBook, error)
	FindById(id int) (database.CategoryBook, error)
	FinByIdSimple(id int) error
	Create(category *database.CategoryBook) (*database.CategoryBook, error)
	Update(id int, category *database.CategoryBook) (*database.CategoryBook, error)
	Delete(id int) error
}

type categoryBookRepo struct {
	db *gorm.DB
}

func NewCategoryBookRepo(db *gorm.DB) CategoryBookRepo {
	return &categoryBookRepo{db}
}

func (r *categoryBookRepo) FindAll() ([]database.CategoryBook, error) {
	var categories []database.CategoryBook
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryBookRepo) FindById(id int) (database.CategoryBook, error) {
	var category database.CategoryBook
	err := r.db.Preload("Books", func(db *gorm.DB) *gorm.DB { 
		return db.Select("id, title, exercpt, category_id, creator") 
		}).Where("Id = ?", id).First(&category).Error
	return category, err
}

func (r *categoryBookRepo) Create(category *database.CategoryBook) (*database.CategoryBook, error) {
	err := r.db.Create(category).Error
	return category, err
}

func (r *categoryBookRepo) Update(id int, category *database.CategoryBook) (*database.CategoryBook, error) {
	err := r.db.Where("id = ?", id).Updates(category).Error
	return category, err
}

func (r *categoryBookRepo) Delete(id int) error {
	var category database.CategoryBook
	err := r.db.Where("id = ?", id).Delete(&category).Error
	return err
}
func (r *categoryBookRepo) FinByIdSimple(id int) error {
	var category database.CategoryBook
	err := r.db.First(&category, id).Error
	return err
}
