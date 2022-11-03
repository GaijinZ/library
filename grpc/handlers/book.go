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

var Books []*protobuff.Book

func marshallBooks(books []*protobuff.Book) []byte {
	json, err := json.Marshal(Books)

	if err != nil {
		log.Fatalf("Cant marshall books %v", err)
	}

	return json
}

func (b *BookServer) GetAllBooks(in *protobuff.GetAllBooksRequest, stream protobuff.Library_GetAllBooksServer) error {
	for _, book := range Books {
		if err := stream.Send(book); err != nil {
			return err
		}
	}

	return nil
}

func (b *BookServer) GetBook(ctx context.Context, in *protobuff.GetBookRequest) (*protobuff.Book, error) {
	for _, book := range Books {
		if book.GetUid() == in.GetUid() {
			log.Printf("Recived: %v", in.Uid)
			redisdb.GetFromRedisDB(fmt.Sprintf("%v", in.Uid))
			return book, nil
		}
	}

	return nil, errors.New("book not found")
}

func (s *BookServer) AddBook(ctx context.Context, in *protobuff.AddBookRequest) (*protobuff.AddBookResponse, error) {
	res := in.GetBook()
	Books = append(Books, res)

	json := marshallBooks(Books)
	redisdb.AddToRedisDB(fmt.Sprintf("%v", res.Uid), json)

	return &protobuff.AddBookResponse{}, nil
}

func (s *BookServer) EditBook(ctx context.Context, in *protobuff.EditBookRequest) (*protobuff.Book, error) {
	res := in.GetBook()

	for index, book := range Books {
		if book.GetUid() == res.GetUid() {
			Books = append(Books[:index], Books[index+1:]...)
			res.Uid = book.GetUid()
			Books = append(Books, res)

			json := marshallBooks(Books)
			redisdb.AddToRedisDB(fmt.Sprintf("%v", res.Uid), json)
			return res, nil
		}
	}

	return res, nil
}

func (s *BookServer) DeleteBook(ctx context.Context, in *protobuff.DeleteBookRequest) (*protobuff.DeleteBookResponse, error) {
	res := &protobuff.DeleteBookResponse{}

	for index, book := range Books {
		if book.GetUid() == in.GetUid() {
			Books = append(Books[:index], Books[index+1:]...)
			res.Success = true
			break
		}
	}

	redisdb.DeleteFromRedisDB(fmt.Sprintf("%v", in.Uid))

	return res, nil
}
