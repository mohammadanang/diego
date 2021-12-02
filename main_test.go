package main_test

import (
	"testing"

	user "github.com/mohammadanang/diego"
)

type MockRepository struct{}

func NewMockRepository() user.Repository {
	return MockRepository{}
}

func (mock MockRepository) Store(entity user.User) error {
	return nil
}

func TestInsertUser(t *testing.T) {
	userRepo := NewMockRepository()
	userService := user.NewService(userRepo)
	if err := userService.Create("brood", "Brodn Way"); err != nil {
		t.Errorf("Got error %v, expect nil when create new user", err.Error())
	}
}
