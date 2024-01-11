package banner

import (
	"ambil-api/domain"
	"ambil-api/middlewares"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bannerHandler struct {
	bannerService Service
}

func NewCategoryHandler(v1 *gin.RouterGroup, bannerService Service) {

	handler := &bannerHandler{bannerService}

	banner := v1.Group("banner")
	banner.GET("", handler.GetAll)
	banner.GET("/:id", handler.GetDetail)

	banner.Use(middlewares.JwtAuthMiddleware())

	banner.POST("", handler.Create)
	banner.POST("/:id", handler.Update)
}

// @Summary Get All Banner
// @Description Get All Banner
// @Accept  json
// @Param BannerFilterRequest query domain.BannerFilterRequest true " BannerFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.BannerData}
// @Router /api/v1/banner [get]
// @Tags Banner
func (h *bannerHandler) GetAll(c *gin.Context) {
	start := time.Now()
	input := domain.BannerFilterRequest{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	banner, err := h.bannerService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        banner,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create Banner
// @Description Create Banner
// @Accept  json
// @Param BannerRequest body domain.BannerRequest true " BannerRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.BannerData}
// @Router /api/v1/banner [post]
// @Tags Banner
func (h *bannerHandler) Create(c *gin.Context) {
	start := time.Now()
	bannerInput := domain.BannerRequest{}

	err := c.ShouldBindJSON(&bannerInput)
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

	banner, err := h.bannerService.Store(bannerInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        banner,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get Detail Banner
// @Description Get Detail Banner
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param banner_id path string true " Banner Id "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.BannerData}
// @Router /api/v1/banner/{banner_id} [get]
// @Tags Banner
func (h *bannerHandler) GetDetail(c *gin.Context) {
	start := time.Now()
	bannerId := c.Param("id")

	result, err := h.bannerService.GetDetail(bannerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        result,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Update Banner
// @Description Update Banner
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param banner_id path string true " Banner Id "
// @Param BannerData body domain.BannerData true " BannerData Schema "
// @Produce  json
// @Success 200 {object} domain.Response
// @Router /api/v1/banner/{banner_id} [post]
// @Tags Banner
func (h *bannerHandler) Update(c *gin.Context) {
	start := time.Now()
	bannerId := c.Param("id")
	bannerRequest := domain.BannerRequest{}

	c.ShouldBindJSON(&bannerRequest)

	err := h.bannerService.Update(bannerId, bannerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message:     "banner updated !",
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
