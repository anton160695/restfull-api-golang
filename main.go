package main

import (
	"crud-golang/crud/config"
	"crud-golang/crud/migrations"
	"crud-golang/crud/src/controller"
	"crud-golang/crud/src/repository"
	"crud-golang/crud/src/router"
	"crud-golang/crud/src/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		log.Println("running migrations...")
		migrations.RunMigration()
		return
	}
	db := config.NewDB()

	// add repository
	userRepo := repository.NewUserRepository(db)
	catRepo := repository.NewCategoryBookRepo(db)
	bookRepo := repository.NewBooksRepository(db)

	// add service
	userServ := service.NewUserService(userRepo)
	catServ := service.NewCategoryBookService(catRepo)
	bookServ := service.NewBookService(bookRepo, userRepo, catRepo)

	// add controller
	userCont := controller.NewUserController(userServ)
	catCont := controller.NewCategoryController(catServ)
	bookCont := controller.NewBookController(bookServ)

	r := gin.Default()
	// add public router
	router.PublicRoutes(r, userCont, catCont, bookCont)
	// add private router
	router.PrivateRoutes(r,  userCont, catCont, bookCont, userRepo)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "connect",
		})
	})

	if err := r.Run("lohalhost:4100"); err != nil {
		log.Fatalf("server tidak dapat berjalan: %v", err)
	}
}
