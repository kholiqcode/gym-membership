package member

import (
	"github.com/labstack/echo/v4"
	"gym/cmd/domain/member/dto"
	"gym/cmd/domain/member/service"
	"gym/internal/protocol/http/response"
	"net/http"
)

type MemberHandlerImpl struct {
	SvcMember service.MemberService
}

func (h MemberHandlerImpl) JoinHistory(ctx echo.Context) error {
	id := ctx.Get("user_id").(float64)

	memberJoinHistory, err := h.SvcMember.GetJoinHistory(uint(id))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"members_join_history": memberJoinHistory,
	})
	return nil
}

func (h MemberHandlerImpl) Detail(ctx echo.Context) error {
	id := ctx.Get("user_id").(float64)

	member, err := h.SvcMember.GetMemberById(uint(id))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", member)
	return nil
}

func (h MemberHandlerImpl) Join(ctx echo.Context) error {
	var memberDto dto.MemberJoinRequest

	if err := ctx.Bind(&memberDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	memberJoin, err := h.SvcMember.StoreMemberJoin(ctx, &memberDto)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", memberJoin)
	return nil
}

func (h MemberHandlerImpl) Create(ctx echo.Context) error {
	var memberDto dto.MemberCreateRequest

	if err := ctx.Bind(&memberDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	member, err := h.SvcMember.Store(&memberDto)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", member)
	return nil
}

func (h MemberHandlerImpl) Login(ctx echo.Context) error {
	var memberDto dto.MemberLoginRequest

	if err := ctx.Bind(&memberDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	if err := ctx.Validate(memberDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	res, err := h.SvcMember.Login(&memberDto)

	if err != nil {
		response.Err(ctx, http.StatusUnauthorized, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", res)
	return nil
}

func (h MemberHandlerImpl) Refresh(ctx echo.Context) error {
	memberId := ctx.Get("user_id").(float64)

	res, err := h.SvcMember.Refresh(uint(memberId))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}
	response.Json(ctx, http.StatusOK, "Success", res)
	return nil
}
