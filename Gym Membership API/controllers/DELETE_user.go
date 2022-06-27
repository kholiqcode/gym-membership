package controllers

import (
	"gym-membership-api/entity"
	"gym-membership-api/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (C *gymControllers) DeleteUserByID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	g, err := C.gymService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": converttoUserResponse(g),
	})
}

func converttoUserResponse(g entity.User) response.UserResponse {
	return response.UserResponse{
		ID:       g.ID,
		Name:     g.Name,
		Phone:    g.Phone,
		Email:    g.Email,
		Password: g.Password,
	}
}
