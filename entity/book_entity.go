package entity

type Books struct {
	Id          int    `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	PublishedAt string `json:"published_at" form:"published_at"`
}

type CreateBooksRequest struct {
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	PublishedAt string `json:"published_at" form:"published_at"`
}

type EditBooksRequest struct {
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	PublishedAt string `json:"published_at" form:"published_at"`
}

type BooksResponse struct {
	Id          int    `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	PublishedAt string `json:"published_at" form:"published_at"`
}

func FormatBookResponse(book Books) BooksResponse {
	return BooksResponse{
		Id:          book.Id,
		Title:       book.Title,
		Author:      book.Author,
		PublishedAt: book.PublishedAt,
	}
}
