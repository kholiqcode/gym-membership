package repository

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_booking/entity"
	"gym/pkg/database"
)

type ClassBookingRepository interface {
	FindAll(ctx echo.Context, pagination *database.Pagination) (*entity.ClassBookingList, error)
	FindAllByMember(ctx echo.Context, memberId uint, pagination *database.Pagination) (*entity.ClassBookingList, error)
	FindByInvoice(ctx echo.Context, invoice string) (*entity.ClassBooking, error)
	Insert(ctx echo.Context, entity *entity.ClassBooking) (*entity.ClassBooking, error)
}
