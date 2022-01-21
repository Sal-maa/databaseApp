package service

import (
	"designpattern/entity"
	"designpattern/repository"
	"fmt"
)

type BookService interface {
	GetBooksService() ([]entity.Books, error)
	GetBookByIdService(id int) (entity.Books, error)
	CreateBookService(bookCreate entity.CreateBooksRequest) (entity.Books, error)
	UpdateBookService(id int, bookUpdate entity.EditBooksRequest) (entity.Books, error)
	DeleteBookService(id int) (entity.Books, error)
}

type bookService struct {
	repository repository.BookRepo
}

func NewBookService(repository repository.BookRepo) *bookService {
	return &bookService{repository}
}

func (s *bookService) GetBooksService() ([]entity.Books, error) {
	books, err := s.repository.GetAllBook()
	return books, err
}

func (s *bookService) CreateBookService(bookCreate entity.CreateBooksRequest) (entity.Books, error) {
	book := entity.Books{}
	book.Title = bookCreate.Title
	book.Author = bookCreate.Author
	book.PublishedAt = bookCreate.PublishedAt

	createBook, err := s.repository.CreateBook(book)
	return createBook, err
}

func (s *bookService) GetBookByIdService(id int) (entity.Books, error) {
	book, err := s.repository.GetBook(id)
	if err != nil {
		fmt.Println(err)
		return book, err
	}
	return book, nil
}

func (s *bookService) UpdateBookService(id int, bookUpdate entity.EditBooksRequest) (entity.Books, error) {
	book, err := s.repository.GetBook(id)
	if err != nil {
		return book, err
	}

	book.Title = bookUpdate.Title
	book.Author = bookUpdate.Author
	book.PublishedAt = bookUpdate.PublishedAt

	updateBook, err := s.repository.UpdateBook(book)
	return updateBook, err
}

func (s *bookService) DeleteBookService(id int) (entity.Books, error) {
	bookID, err := s.GetBookByIdService(id)
	if err != nil {
		return bookID, err
	}

	deleteBook, err := s.repository.DeleteBook(bookID)
	return deleteBook, err
}
