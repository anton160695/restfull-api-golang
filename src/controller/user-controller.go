package controller

import (
	"crud-golang/crud/src/model"
	"crud-golang/crud/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	// 1. req model jsom
	var req model.UserRegisterReqAndRes
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "invalid JSON format",
		})

		return
	}

	// 2. call service
	res, err := c.userService.Register(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "internal server error",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var req model.UserLoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "invalid JSON format",
		})
		return
	}
	res, err := c.userService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *UserController) Me(ctx *gin.Context) {
	// pick user id from session on jwt generate
	userName, exist := ctx.Get("username")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": false,
			"code":   400,
			"error":  "user not found",
		})
		return
	}
	// call service
	user, err := c.userService.Me(userName.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (c *UserController) Logout(ctx *gin.Context) {
	// pick user id from sesson on jwt generate
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}
	// konversi dari float64 ke int
	userIdInt := int(userId.(float64))
	//cal service
	err := c.userService.Logout(userIdInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "logout success",
	})
}

func (c *UserController) Update(ctx *gin.Context) {
	var req model.UserUpdateReqAndRes
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "invalid JSON format",
		})

		return
	}
	// pick user id
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}
	userIdInt := int(userId.(float64))
	res, err := c.userService.Update(userIdInt, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})

}
