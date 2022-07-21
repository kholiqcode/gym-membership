package class_booking

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/class_booking/dto"
	"gym/cmd/domain/class_booking/service"
	"gym/internal/protocol/http/response"
	"gym/pkg/database"
	"net/http"
)

type ClassBookingHandlerImpl struct {
	Svc service.ClassBookingService
}

func (h ClassBookingHandlerImpl) Get(ctx echo.Context) error {
	pagination := database.NewPagination(ctx)

	classBookings, err := h.Svc.GetAll(ctx, pagination)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"class_bookings": map[string]interface{}{
			"data":       classBookings,
			"sort":       pagination.GetSort(),
			"page":       pagination.GetPage(),
			"page_size":  pagination.GetLimit(),
			"total_page": pagination.GetTotalPage(),
			"total_rows": pagination.GetTotalRows(),
		},
	})
	return nil
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
