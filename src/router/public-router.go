package router

import (
	"crud-golang/crud/src/controller"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(router *gin.Engine, userCont *controller.UserController, catCont *controller.CategoryController, bookCont *controller.BookController) {
	public := router.Group("/v1/public")
	{
		//user
		public.POST("/user/register", userCont.Register)
		public.POST("/user/login", userCont.Login)
		// category
		public.GET("/category", catCont.FindAllCategory)
		public.GET("/category/:catId", catCont.FindCategoryById)
		// book
		public.GET("/books", bookCont.GetAllBook)
		public.GET("/books/:bookId", bookCont.GetBookById)

	}
}
