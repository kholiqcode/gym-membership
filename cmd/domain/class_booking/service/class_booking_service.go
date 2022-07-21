package service

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_booking/dto"
	"gym/pkg/database"
)

type ClassBookingService interface {
	GetAll(ctx echo.Context, pagination *database.Pagination) (*dto.ClassBookingListResponse, error)
	GetAllByMember(ctx echo.Context, pagination *database.Pagination) (*dto.ClassBookingListResponse, error)
	GetByInvoice(ctx echo.Context, invoice string) (*dto.ClassBookingResponse, error)
	Create(ctx echo.Context, request *dto.ClassBookingCreateRequest) (*dto.ClassBookingResponse, error)
}
