package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	_ "net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"awesomeProject/graph/model"
)

// Mock HTTP client for testing
type MockHTTPClient struct{}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if req.Method == "POST" && req.URL.String() == "http://localhost:3000/books" {
		newBook := &model.Book{
			ID:        "12345",
			Title:     "Test Book",
			Author:    "Test Author",
			Published: "2022-01-01",
		}
		jsonData, _ := json.Marshal(newBook)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(jsonData)),
			Header:     make(http.Header),
		}, nil
	} else if req.Method == "GET" && req.URL.String() == "http://localhost:3000/books" {
		books := []*model.Book{
			{
				ID:        "123",
				Title:     "Book 1",
				Author:    "Author 1",
				Published: "2022-02-01",
			},
			{
				ID:        "456",
				Title:     "Book 2",
				Author:    "Author 2",
				Published: "2022-02-02",
			},
		}
		jsonData, _ := json.Marshal(books)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(jsonData)),
			Header:     make(http.Header),
		}, nil
	} else {
		return nil, errors.New("Invalid request")
	}
}

func TestMutationResolver_CreateBook(t *testing.T) {
	// Mock GraphQL context
	ctx := context.TODO()

	// Create a new mock resolver with the mock HTTP client
	resolver := &mutationResolver{&Resolver{}}

	// Call the resolver method
	book, err := resolver.CreateBook(ctx, "Test Book", "Test Author", "2022-01-01")

	// Assert that the result is correct and there are no errors
	assert.Nil(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, "12345", book.ID)
	assert.Equal(t, "Test Book", book.Title)
	assert.Equal(t, "Test Author", book.Author)
	assert.Equal(t, "2022-01-01", book.Published)
}

func TestQueryResolver_Books(t *testing.T) {
	// Mock GraphQL context
	ctx := context.TODO()

	// Create a new mock resolver with the mock HTTP client
	resolver := &queryResolver{&Resolver{}}

	// Call the resolver method
	books, err := resolver.Books(ctx)

	// Assert that the result is correct and there are no errors
	assert.Nil(t, err)
	assert.NotNil(t, books)
	assert.Equal(t, 2, len(books))
	assert.Equal(t, "123", books[0].ID)
	assert.Equal(t, "Book 1", books[0].Title)
	assert.Equal(t, "Author 1", books[0].Author)
	assert.Equal(t, "2022-02-01", books[0].Published)
	assert.Equal(t, "456", books[1].ID)
}
