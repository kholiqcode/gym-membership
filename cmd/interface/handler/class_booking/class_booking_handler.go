package class_booking

import "github.com/labstack/echo/v4"

type ClassBookingHandler interface {
	Get(ctx echo.Context) error
	Order(ctx echo.Context) error
}
