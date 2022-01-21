package repository

import (
	"database/sql"
	"designpattern/entity"
	"fmt"
)

type BookRepo interface {
	GetAllBook() ([]entity.Books, error)
	CreateBook(book entity.Books) (entity.Books, error)
	GetBook(idParam int) (entity.Books, error)
	UpdateBook(book entity.Books) (entity.Books, error)
	DeleteBook(book entity.Books) (entity.Books, error)
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) GetAllBook() ([]entity.Books, error) {
	var books []entity.Books
	result, err := r.db.Query("SELECT id,title, author,published_at FROM books")
	if err != nil {
		fmt.Println(err)
		return books, fmt.Errorf("failed to scan")
	}
	defer result.Close()

	for result.Next() {
		var book entity.Books
		err := result.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt)
		if err != nil {
			return books, fmt.Errorf("failed to scan")
		}
		books = append(books, book)
	}
	return books, err
}

func (r *bookRepository) CreateBook(book entity.Books) (entity.Books, error) {
	_, err := r.db.Exec("INSERT INTO books(title, author, published_at) VALUES(?,?,?)", book.Title, book.Author, book.PublishedAt)
	return book, err
}

func (r *bookRepository) GetBook(idParam int) (entity.Books, error) {
	var book entity.Books
	result, err := r.db.Query("SELECT id,title, author,published_at FROM books WHERE id=?", idParam)
	if err != nil {
		return book, fmt.Errorf("failed in query")
	}
	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	errScan := result.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt)
	if errScan != nil {
		return book, fmt.Errorf("failed to read data")
	}
	if idParam == book.Id {
		return book, nil
	}
	return book, fmt.Errorf("book not found")
}

func (r *bookRepository) UpdateBook(book entity.Books) (entity.Books, error) {
	result, err := r.db.Exec("UPDATE books SET title=?, author=?, published_at=? WHERE id=?", book.Title, book.Author, book.PublishedAt, book.Id)
	if err != nil {
		return book, fmt.Errorf("failed to update data")
	}
	NotAffected, _ := result.RowsAffected()
	if NotAffected == 0 {
		return book, fmt.Errorf("failed to find data id")
	}
	return book, nil
}

func (r *bookRepository) DeleteBook(book entity.Books) (entity.Books, error) {
	_, err := r.db.Exec("DELETE FROM books WHERE id=?", book.Id)
	return book, err
}
