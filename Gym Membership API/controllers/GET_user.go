package controllers

import (
	"gym-membership-api/entity"
	"gym-membership-api/response"
	"gym-membership-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type gymControllers struct {
	gymService service.Service
}

func NewGymControllers(gymService service.Service) *gymControllers {
	return &gymControllers{gymService}
}

func (C *gymControllers) GetRootControllers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response_code": "200",
		"message":       "Success!",
		"version":       "1.0.0-alpha-1"})
}

func (C *gymControllers) GetAllUser(c *gin.Context) {
	gyms, err := C.gymService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var gymsResponse []response.UserResponse
	for _, g := range gyms {
		gymResponse := ConvertToGymResponse(g)

		gymsResponse = append(gymsResponse, gymResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gymsResponse,
	})
}

func (C *gymControllers) GetUserById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	g, err := C.gymService.FindByID(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	gymResponse := ConvertToGymResponse(g)
	c.JSON(http.StatusOK, gin.H{
		"data": gymResponse,
	})
}

func ConvertToGymResponse(g entity.User) response.UserResponse {
	return response.UserResponse{
		ID:       g.ID,
		Name:     g.Name,
		Phone:    g.Phone,
		Email:    g.Email,
		Password: g.Password,
	}
}

// func (C *gymControllers)GetAllUser(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"name":    "Al Tsaqif Nugraha Ahmad",
// 		"class":   "06TPLP008",
// 		"address": "Perumahan Panorama Bali Residence Ciseeng Bogor",
// 	})
// }

// func (C *gymControllers)GetUserById(c *gin.Context) {
// 	id := c.Param("id")
// 	name := c.Param("name")
// 	c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
// }

// func (C *gymControllers)GetUserByQuery(c *gin.Context) {
// 	id := c.Query("id")
// 	name := c.Query("name")
// 	c.JSON(http.StatusOK, gin.H{"name": name, "id": id})
// }
