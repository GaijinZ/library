package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/GaijinZ/grpc/protobuff"
	"github.com/GaijinZ/grpc/redisdb"
)

type BookServer struct {
	protobuff.UnimplementedLibraryServer
}

var books []*protobuff.Book
var book redisdb.Books
var redis redisdb.Redis = book

func marshallBooks(books []*protobuff.Book) []byte {
	json, err := json.Marshal(books)

	if err != nil {
		log.Fatalf("Cant marshall books %v", err)
	}

	return json
}

func (b *BookServer) GetAllBooks(in *protobuff.GetAllBooksRequest, stream protobuff.Library_GetAllBooksServer) error {
	for _, book := range books {
		if err := stream.Send(book); err != nil {
			return err
		}
	}

	return nil
}

func (b *BookServer) GetBook(ctx context.Context, in *protobuff.GetBookRequest) (*protobuff.Book, error) {
	for _, book := range books {
		if book.GetUid() == in.GetUid() {
			log.Printf("Recived: %v", in.Uid)
			redis.Get(fmt.Sprintf("%v", in.Uid))
			return book, nil
		}
	}

	return nil, errors.New("book not found")
}

func (b *BookServer) AddBook(ctx context.Context, in *protobuff.AddBookRequest) (*protobuff.AddBookResponse, error) {
	res := in.GetBook()
	books = append(books, res)

	json := marshallBooks(books)
	redis.Set(fmt.Sprintf("%v", res.Uid), json)

	return &protobuff.AddBookResponse{}, nil
}

func (b *BookServer) EditBook(ctx context.Context, in *protobuff.EditBookRequest) (*protobuff.Book, error) {
	res := in.GetBook()

	for index, book := range books {
		if book.GetUid() == res.GetUid() {
			books = append(books[:index], books[index+1:]...)
			res.Uid = book.GetUid()
			books = append(books, res)

			json := marshallBooks(books)
			redis.Set(fmt.Sprintf("%v", res.Uid), json)
			return res, nil
		}
	}

	return res, nil
}

func (b *BookServer) DeleteBook(ctx context.Context, in *protobuff.DeleteBookRequest) (*protobuff.DeleteBookResponse, error) {
	res := &protobuff.DeleteBookResponse{}

	for index, book := range books {
		if book.GetUid() == in.GetUid() {
			books = append(books[:index], books[index+1:]...)
			res.Success = true
			break
		}
	}

	redis.Delete(fmt.Sprintf("%v", in.Uid))

	return res, nil
}
