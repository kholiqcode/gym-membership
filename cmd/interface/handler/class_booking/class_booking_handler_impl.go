package class_booking

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_booking/dto"
	"gym/cmd/domain/class_booking/service"
	"gym/internal/protocol/http/response"
	"net/http"
)

type ClassBookingHandlerImpl struct {
	Svc service.ClassBookingService
}

func (h ClassBookingHandlerImpl) Order(ctx echo.Context) error {
	var request dto.ClassBookingCreateRequest

	if err := ctx.Bind(&request); err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	classBooking, err := h.Svc.Create(ctx, &request)

	if err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", classBooking)
	return nil
}
