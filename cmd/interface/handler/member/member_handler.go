package member

import (
	"github.com/labstack/echo/v4"
)

type MemberHandler interface {
	JoinHistory(ctx echo.Context) error
	Detail(ctx echo.Context) error
	Join(ctx echo.Context) error
	Create(ctx echo.Context) error
	Login(ctx echo.Context) error
	Refresh(ctx echo.Context) error
}
