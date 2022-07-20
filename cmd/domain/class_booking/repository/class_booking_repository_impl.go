package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gym/cmd/domain/class_booking/entity"
)

type ClassBookingRepositoryImpl struct {
	Db *gorm.DB
}

func (r *ClassBookingRepositoryImpl) FindAll(ctx echo.Context) (*entity.ClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassBookingRepositoryImpl) FindByInvoice(ctx echo.Context, trainerID uint) (*entity.ClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassBookingRepositoryImpl) Insert(ctx echo.Context, classBooking *entity.ClassBooking) (*entity.ClassBooking, error) {
	if e := r.Db.Debug().Preload("ClassBookingDetail").Create(&classBooking).Error; e != nil {
		return nil, e
	}
	return classBooking, nil
}
