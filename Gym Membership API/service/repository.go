package service

import (
	"gym-membership-api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
	FindByID(ID int) (entity.User, error)
	Create(service entity.User) (entity.User, error)
	Update(service entity.User) (entity.User, error)
	Update2(service entity.User) (entity.User, error)
	Delete(service entity.User) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var gyms []entity.User
	err := r.db.Find(&gyms).Error
	return gyms, err
}

func (r *repository) FindByID(ID int) (entity.User, error) {
	var gym entity.User
	err := r.db.Find(&gym, ID).Error
	return gym, err
}

func (r *repository) Create(service entity.User) (entity.User, error) {
	err := r.db.Create(&service).Error
	return service, err
}

func (r *repository) Update(service entity.User) (entity.User, error) {
	err := r.db.Save(&service).Error
	return service, err
}

func (r *repository) Update2(service entity.User) (entity.User, error) {
	err := r.db.Save(&service).Error
	return service, err
}

func (r *repository) Delete(service entity.User) (entity.User, error) {
	err := r.db.Delete(&service).Error
	return service, err
}
