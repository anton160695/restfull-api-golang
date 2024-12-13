package model

import "crud-golang/crud/src/database"

type CreateBook struct {
	Title      string `json:"title"`
	Excerpt    string `json:"excerpt"`
	Content    string `json:"content"`
	CategoryID int    `json:"category_id"`
}

type UpdateBook struct {
	Title      string `json:"title,omitempty"`
	Excerpt    string `json:"excerpt,omitempty"`
	Content    string `json:"content,omitempty"`
	CategoryID int    `json:"category_id,omitempty"`
}

type CreatBookRes struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Excerpt  string `json:"excerpt"`
	Category string `json:"category"`
}

type AllBooks struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Excerpt  string `json:"excerpt"`
	Category string `json:"category"`
	Creator  string `json:"creator"`
}

type BookDetails struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Excerpt  string `json:"excerpt"`
	Content  string `json:"content"`
	Creator  string `json:"creator"`
	Category string `json:"category"`
}

func ToCreateRes(book database.Book) CreatBookRes {
	return CreatBookRes{
		Id:       book.Id,
		Title:    book.Title,
		Excerpt:  book.Exercpt,
		Category: book.Category.Name,
	}
}

func ToAllBookDetails(books []database.Book) []AllBooks {
	mappedBooks := make([]AllBooks, len(books))
	for i, book := range books {
		mappedBooks[i] = AllBooks{
			Id:       book.Id,
			Title:    book.Title,
			Excerpt:  book.Exercpt,
			Creator:  book.Creator,
			Category: book.Category.Name,
		}
	}

	return mappedBooks
}

func ToBooksDetails(book database.Book) BookDetails {
	return BookDetails{
		Id:       book.Id,
		Title:    book.Title,
		Excerpt:  book.Exercpt,
		Content:  book.Content,
		Creator:  book.Creator,
		Category: book.Category.Name,
	}
}
