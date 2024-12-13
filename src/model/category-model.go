package model

import "crud-golang/crud/src/database"

type CategoryReq struct {
	Name string `json:"name"`
}

type CategoryRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BookDetail struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Excerpt  string `json:"excerpt"`
}

type CategoryBookById struct {
	Id    int          `json:"id"`
	Name  string       `json:"name"`
	Books []BookDetail `json:"books"`
}

func mapBookDetails(books []database.Book) []BookDetail {
	mappedBooks := make([]BookDetail, len(books))
	for i, book := range books {
		mappedBooks[i] = BookDetail{
			Id:       book.Id,
			Title:    book.Title,
			Excerpt:  book.Exercpt,
		}
	}

	return mappedBooks
}

func ToCategoryBookById(category database.CategoryBook) CategoryBookById {
	return CategoryBookById{
		Id:    category.Id,
		Name:  category.Name,
		Books: mapBookDetails(category.Books),
	}
}

func ToAllCategory(categorys []database.CategoryBook) []CategoryRes {
	books := make([]CategoryRes, len(categorys))
	for i, category := range categorys {
		books[i] = CategoryRes{
			Id:   category.Id,
			Name: category.Name,
		}
	}
	return books
}
