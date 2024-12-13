package controller

import (
	"crud-golang/crud/src/model"
	"crud-golang/crud/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{bookService: bookService}
}

func (c *BookController) GetAllBook(ctx *gin.Context) {
	books, err := c.bookService.GetAllBook()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (c *BookController) GetBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("bookId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid params format",
		})
		return
	}
	book, err := c.bookService.GetBookById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var req model.CreateBook
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON format",
		})
		return
	}
	id, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	userId := int(id.(float64))
	book, err := c.bookService.CreateBook(userId, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	var req model.UpdateBook
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON format",
		})
		return
	}
	id, err := strconv.Atoi(ctx.Params.ByName("bookId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid params format",
		})
		return
	}
	book, err := c.bookService.UpdateBook(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("bookId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid params format",
		})
		return
	}
	err = c.bookService.DeleteBook(id)
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

