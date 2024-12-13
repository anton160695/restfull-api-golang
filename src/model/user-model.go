package model

import "crud-golang/crud/src/database"

type UserRegisterReqAndRes struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserUpdateReqAndRes struct {
	Name     string `json:"name,omitempty"`
	UserName string `json:"username,omitempty"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRes struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

func ToUserResponse(user database.Users) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Token:    user.Token,
	}
}

func ToUserLoginRes(user database.Users) UserLoginRes {
	return UserLoginRes{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Token:    user.Token,
	}
}

func ToUserRegisterRes(user database.Users) UserRegisterReqAndRes {
	return UserRegisterReqAndRes{
		Username: user.Username,
		Name:     user.Name,
	}
}

func ToUserUpdateRes(user database.Users) UserUpdateReqAndRes {
	return UserUpdateReqAndRes{
		Name:     user.Name,
		UserName: user.Username,
	}
}
