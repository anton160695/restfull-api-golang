package repository

import (
	"crud-golang/crud/src/database"

	"gorm.io/gorm"
)

type UserRepository interface {
	CretaUser(user *database.Users) (*database.Users, error)
	LoginUSer(username string) (*database.Users, error)
	LogOut(user *database.Users) error
	UpdateUser(id int, user *database.Users) (*database.Users, error)
	FindByUsername(username string) (*database.Users, error)
	FindUserByID(userID int) (*database.Users, error)
	UpdateToken(user *database.Users) (*database.Users, error)
	FindUserIdAndToken(id int, token string) (*database.Users, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) CretaUser(user *database.Users) (*database.Users, error) {
	err := r.db.Create(user).Error
	return user, err
}

func (r *userRepo) LoginUSer(username string) (*database.Users, error) {
	var user database.Users
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepo) LogOut(user *database.Users) error {
	err := r.db.Save(&user).Error
	return err
}

func (r *userRepo) UpdateUser(id int, user *database.Users) (*database.Users, error) {
	err := r.db.Where("id = ?", id).Updates(user).Error
	return user, err
}

func (r *userRepo) FindByUsername(username string) (*database.Users, error) {
	var user database.Users
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepo) UpdateToken(user *database.Users) (*database.Users, error) {
	err := r.db.Save(&user).Select("id, name, token, username").Error
	return user, err
}

func (r *userRepo) FindUserByID(userID int) (*database.Users, error) {
	var user database.Users
	err := r.db.Where("id = ?", userID).First(&user).Error
	return &user, err
}

func (r *userRepo) FindUserIdAndToken(id int, token string) (*database.Users, error) {
	var user database.Users
	err := r.db.Where("id = ? AND token = ?", id, token).First(&user).Error
	return &user, err
}
