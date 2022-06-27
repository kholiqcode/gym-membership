package service

import (
	"errors"
	"fmt"
	"gym-membership-api/entity"
)

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) FindAll() ([]entity.User, error) {
	var gyms []entity.User
	fmt.Println("Test 12345")
	return gyms, errors.New("dummy")
}

func (r *fileRepository) FindByID(ID int) (entity.User, error) {
	var gym entity.User
	fmt.Println("Test 12345")
	return gym, nil
}

func (r *fileRepository) Create(service entity.User) (entity.User, error) {
	fmt.Println("Test 12345")
	return service, nil
}

