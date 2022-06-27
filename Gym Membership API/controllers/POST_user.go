package controllers

import (
	"fmt"
	"gym-membership-api/entity"
	"gym-membership-api/models"
	"gym-membership-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (C *gymControllers) CreateUser_POST(c *gin.Context) {
	var gymRequest models.Users
	err := c.ShouldBindJSON(&gymRequest)
	if err != nil {

		// Error Message 2
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage2 := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage2)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	gym, err := C.gymService.Create(gymRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToGymResponse4(gym),
	})

}

func convertToGymResponse4(g entity.User) response.UserResponse {
	return response.UserResponse{
		ID:       g.ID,
		Name:     g.Name,
		Phone:    g.Phone,
		Email:    g.Email,
		Password: g.Password,
	}
}
