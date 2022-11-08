package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/GaijinZ/grpc/pkg/protobuff"
	"github.com/GaijinZ/grpc/pkg/redisdb"
)

type BookServer struct {
	protobuff.UnimplementedLibraryServer
}

var books []*protobuff.Book
var book redisdb.Books
var redis redisdb.Redis = book

// helper used to return the JSON from protobuff
func marshallBooks(books []*protobuff.Book) []byte {
	json, err := json.Marshal(books)

	if err != nil {
		log.Fatalf("Cant marshall books %v", err)
	}

	return json
}

// stream all books
func (b *BookServer) GetAllBooks(in *protobuff.GetAllBooksRequest, stream protobuff.Library_GetAllBooksServer) error {
	for _, book := range books {
		if err := stream.Send(book); err != nil {
			return err
		}
	}

	return nil
}

// get single book by uid from protobuff and redis db
func (b *BookServer) GetBook(ctx context.Context, in *protobuff.GetBookRequest) (*protobuff.Book, error) {
	for _, book := range books {
		if book.GetUid() == in.GetUid() {
			get, err := redis.Get(fmt.Sprintf("%v", in.Uid))
			if err != nil {
				log.Println("could not get data", get)
			}

			return book, nil
		}
	}

	return nil, errors.New("book not found")
}

// add a book to protobuff and redis db
func (b *BookServer) AddBook(ctx context.Context, in *protobuff.AddBookRequest) (*protobuff.AddBookResponse, error) {
	res := in.GetBook()
	books = append(books, res)

	json := marshallBooks(books)
	set, err := redis.Set(fmt.Sprintf("%v", res.Uid), json)
	if err != nil {
		log.Println("could not set data", set)
	}

	return &protobuff.AddBookResponse{}, nil
}

// update a book by uid in protobuff and redis db
func (b *BookServer) EditBook(ctx context.Context, in *protobuff.EditBookRequest) (*protobuff.Book, error) {
	res := in.GetBook()

	for index, book := range books {
		if book.GetUid() == res.GetUid() {
			books = append(books[:index], books[index+1:]...)
			res.Uid = book.GetUid()
			books = append(books, res)

			json := marshallBooks(books)
			update, err := redis.Set(fmt.Sprintf("%v", res.Uid), json)
			if err != nil {
				log.Println("could not update data", update)
			}

			return res, nil
		}
	}

	return res, nil
}

// delete a book from protobuff and redis db
func (b *BookServer) DeleteBook(ctx context.Context, in *protobuff.DeleteBookRequest) (*protobuff.DeleteBookResponse, error) {
	res := &protobuff.DeleteBookResponse{}

	for index, book := range books {
		if book.GetUid() == in.GetUid() {
			books = append(books[:index], books[index+1:]...)
			res.Success = true
			break
		}
	}

	del, err := redis.Delete(fmt.Sprintf("%v", in.Uid))
	if err != nil {
		log.Println("could not del data", del)
	}

	return res, nil
}
