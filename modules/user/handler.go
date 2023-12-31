package user

import (
	"ambil-api/domain"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService Service
}

func NewUserHandler(v1 *gin.RouterGroup, userService Service) {

	handler := &userHandler{userService}

	v1.POST("/login", handler.Login)
	v1.POST("/register", handler.PostUser)
}

// @Summary Login
// @Description Login
// @Accept  json
// @Param AuthRequest body domain.AuthRequest true " AuthRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{}
// @Router /api/v1/login [post]
// @Tags User
func (h *userHandler) Login(c *gin.Context) {
	start := time.Now()
	input := domain.AuthRequest{}

	c.ShouldBindJSON(&input)

	token, err := h.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        token,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Register
// @Description Register
// @Accept  json
// @Param RegisterRequest body domain.RegisterRequest true " RegisterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.UserData}
// @Router /api/v1/register [post]
// @Tags User
func (h *userHandler) PostUser(c *gin.Context) {
	start := time.Now()
	userInput := domain.RegisterRequest{}

	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": errorMessages,
		})

		return
	}

	user, err := h.userService.Create(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        user,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
