package service

import (
	"gym-membership-api/entity"
	"gym-membership-api/models"
)

type Service interface {
	FindAll() ([]entity.User, error)
	FindByID(ID int) (entity.User, error)
	Create(gymRequest models.Users) (entity.User, error)
	Update(ID int, gymRequest models.Users) (entity.User, error)
	Update2(ID int, gymRequest models.Users2) (entity.User, error)
	Delete(ID int) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.User, error) {
	gyms, err := s.repository.FindAll()
	return gyms, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (entity.User, error) {
	gym, err := s.repository.FindByID(ID)
	return gym, err
}

func (s *service) Create(gymRequest models.Users) (entity.User, error) {
	// price, _ := gymRequest.Price.Int64()
	// rating, _ := gymRequest.Rating.Int64()
	// discount, _ := gymRequest.Discount.Int64()
	gym := entity.User{
		Name:     gymRequest.Name,
		Phone:    gymRequest.Phone,
		Email:    gymRequest.Email,
		Password: gymRequest.Password,
	}
	newGym, err := s.repository.Create(gym)
	return newGym, err
}

func (s *service) Update(ID int, gymRequest models.Users) (entity.User, error) {
	gym, _ := s.repository.FindByID(ID)
	// price, _ := gymRequest.Price.Int64()
	// rating, _ := gymRequest.Rating.Int64()
	// discount, _ := gymRequest.Discount.Int64()

	gym.Name = gymRequest.Name
	gym.Phone = gymRequest.Phone
	gym.Email = gymRequest.Email
	gym.Password = gymRequest.Password

	newGym, err := s.repository.Update(gym)
	return newGym, err
}

func (s *service) Update2(ID int, gymRequest models.Users2) (entity.User, error) {
	gym, _ := s.repository.FindByID(ID)
	// price, _ := gymRequest.Price.Int64()
	// rating, _ := gymRequest.Rating.Int64()
	// discount, _ := gymRequest.Discount.Int64()

	gym.Name = gymRequest.Name
	gym.Phone = gymRequest.Phone
	gym.Email = gymRequest.Email
	gym.Password = gymRequest.Password

	newGym, err := s.repository.Update2(gym)
	return newGym, err
}

func (s *service) Delete(ID int) (entity.User, error) {
	gym, _ := s.repository.FindByID(ID)
	newGym, err := s.repository.Delete(gym)
	return newGym, err
}
