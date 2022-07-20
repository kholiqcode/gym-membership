package class_booking

import "github.com/labstack/echo/v4"

type ClassBookingHandler interface {
	Order(ctx echo.Context) error
}
