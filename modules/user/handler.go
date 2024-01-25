package user

import (
	"ambil-api/domain"
	"ambil-api/middlewares"
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

	user := v1.Group("user")
	user.POST("bookmark/merchant", middlewares.JwtAuthMiddleware(), handler.AddBookmark)

	userLevels := user.Group("levels")
	userLevels.GET("", middlewares.JwtAuthMiddleware(), handler.GetUserLevels)
	userLevels.POST("", middlewares.JwtAuthMiddleware(), handler.CreateUserLevels)
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
			"errors": errorMessages,
		})

		return
	}

	user, err := h.userService.Create(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        user,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get All User Level
// @Description Get All User Level
// @Accept  json
// @Param UserLevelRequest query domain.UserLevelRequest true " UserLevelRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.UserLevelData}
// @Router /api/v1/user/levels [get]
// @Tags User
func (h *userHandler) GetUserLevels(c *gin.Context) {
	start := time.Now()
	input := domain.UserLevelRequest{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	userLevels, err := h.userService.GetAllUserLevel(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        userLevels,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Add Favourite Merchant
// @Description Add Favourite Merchant
// @Accept  json
// @Param UserMerchantFavouriteRequest body domain.UserMerchantFavouriteRequest true " UserMerchantFavouriteRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.UserMerchantFavouriteData}
// @Router /api/v1/user/bookmark [post]
// @Tags User
func (h *userHandler) AddBookmark(c *gin.Context) {
	start := time.Now()
	input := domain.UserMerchantFavouriteRequest{}

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	result, err := h.userService.AddMerchantFavourite(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        result,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create User Level
// @Description Create User Level
// @Accept  json
// @Param UserLevelRequest body domain.UserLevelRequest true " UserLevelRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.UserLevelData}
// @Router /api/v1/user/levels [post]
// @Tags User
func (h *userHandler) CreateUserLevels(c *gin.Context) {
	start := time.Now()
	input := domain.UserLevelRequest{}

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	userLevels, err := h.userService.CreateUserLevel(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        userLevels,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
