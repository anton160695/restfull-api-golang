package controller

import (
	"crud-golang/crud/src/model"
	"crud-golang/crud/src/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	catServ service.CategoryBookService
}

func NewCategoryController(catServ service.CategoryBookService) *CategoryController {
	return &CategoryController{catServ}
}

func (c *CategoryController) FindAllCategory(ctx *gin.Context) {
	category, err := c.catServ.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req model.CategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON format",
		})
		return
	}

	category, err := c.catServ.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var req model.CategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON format",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Params.ByName("catId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid params format",
		})
		return
	}

	log.Println("id: ", id)
	category, err := c.catServ.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("catId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid params format",
		})
		return
	}
	err = c.catServ.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted success",
	})
}

func (c *CategoryController) FindCategoryById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("catId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid params format",
		})
		return
	}
	category, err := c.catServ.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}
