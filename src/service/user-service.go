package service

import (
	"crud-golang/crud/src/database"
	"crud-golang/crud/src/middleware"
	"crud-golang/crud/src/model"
	"crud-golang/crud/src/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(req model.UserLoginReq) (*model.UserLoginRes, error)
	Register(req model.UserRegisterReqAndRes) (*model.UserRegisterReqAndRes, error)
	Update(id int, req model.UserUpdateReqAndRes) (*model.UserUpdateReqAndRes, error)
	Logout(userId int) error
	Me(username string) (*model.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo,
	}
}

func (s *userService) Register(req model.UserRegisterReqAndRes) (*model.UserRegisterReqAndRes, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &database.Users{
		Name:     req.Name,
		Username: req.Username,
		Password: string(hashPassword),
	}

	createUser, err := s.repo.CretaUser(user)
	if err != nil {
		return nil, err
	}
	response := model.ToUserRegisterRes(*createUser)

	return &response, nil
}

func (s *userService) Me(username string) (*model.UserResponse, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	userResponse := model.ToUserResponse(*user)
	return &userResponse, nil
}

func (s *userService) Logout(userId int) error {
	// cek user ada atau tidak
	user, err := s.repo.FindUserByID(userId)
	if err != nil {
		return err
	}
	// update token
	user.Token = ""
	// update token
	err = s.repo.LogOut(user)
	if err != nil {
		return err
	}
	return err
}

func (s *userService) Login(req model.UserLoginReq) (*model.UserLoginRes, error) {
	user, err := s.repo.LoginUSer(req.Username)
	if err != nil {
		return nil, err
	}
	// validasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	acessToken, err := middleware.GenerateToken(*user)
	if err != nil {
		return nil, err
	}
	log.Println("acess token: ", acessToken)
	user.Token = acessToken
	userToken , err := s.repo.UpdateToken(user)
	if err != nil {
		return nil, err
	}
	response := model.ToUserLoginRes(*userToken)
	log.Println("response: ", response)
	return &response, nil
}

func (s *userService) Update(id int, req model.UserUpdateReqAndRes) (*model.UserUpdateReqAndRes, error) {
	_, err := s.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	userUpdate := &database.Users{
		Name:     req.Name,
		Username: req.UserName,
	}
	update, err := s.repo.UpdateUser(id, userUpdate)
	if err != nil {
		return nil, err
	}
	response := model.ToUserUpdateRes(*update)
	return &response, nil
}
