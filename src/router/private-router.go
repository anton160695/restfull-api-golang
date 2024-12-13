package router

import (
	"crud-golang/crud/src/controller"
	"crud-golang/crud/src/middleware"
	"crud-golang/crud/src/repository"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(router *gin.Engine, userCont *controller.UserController, catCont *controller.CategoryController, bookCont *controller.BookController, userRepo repository.UserRepository) {
	private := router.Group("/v1/private")
	private.Use(middleware.AuthMiddleware(userRepo))
	{
		private.GET("/user/me", userCont.Me)
		private.PATCH("/user/me", userCont.Update)
		private.DELETE("/user/logout", userCont.Logout)
		//category
		private.POST("/category", catCont.CreateCategory)
		private.PATCH("/category/:catId", catCont.UpdateCategory)
		private.DELETE("/category/:catId", catCont.DeleteCategory)
		//book
		private.POST("/book", bookCont.CreateBook)
		private.PATCH("/book/:bookId", bookCont.UpdateBook)
		private.DELETE("/book/:bookId", bookCont.DeleteBook)
	}
}
