package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gym/cmd/domain/class_booking/entity"
	"gym/pkg/database"
)

type ClassBookingRepositoryImpl struct {
	Db *gorm.DB
}

func (r *ClassBookingRepositoryImpl) FindAll(ctx echo.Context, pagination *database.Pagination) (*entity.ClassBookingList, error) {
	var classBookings entity.ClassBookingList

	if e := r.Db.Debug().Preload("ClassBookingDetail").Scopes(database.Paginate(classBookings, pagination, r.Db)).Preload(clause.Associations).Find(&classBookings).Error; e != nil {
		return nil, e
	}
	return &classBookings, nil
}

func (r *ClassBookingRepositoryImpl) FindAllByMember(ctx echo.Context, memberId uint, pagination *database.Pagination) (*entity.ClassBookingList, error) {
	var classBookings entity.ClassBookingList

	if e := r.Db.Debug().Preload("ClassBookingDetail").Where("member_id = ?", memberId).Scopes(database.Paginate(classBookings, pagination, r.Db)).Preload(clause.Associations).Find(&classBookings).Error; e != nil {
		return nil, e
	}
	return &classBookings, nil
}

func (r *ClassBookingRepositoryImpl) FindByInvoice(ctx echo.Context, invoice string) (*entity.ClassBooking, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassBookingRepositoryImpl) Insert(ctx echo.Context, classBooking *entity.ClassBooking) (*entity.ClassBooking, error) {
	if e := r.Db.Debug().Preload("ClassBookingDetail").Create(&classBooking).Error; e != nil {
		return nil, e
	}
	return classBooking, nil
}
