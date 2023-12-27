package merchant

import (
	"ambil-api/domain"
	"ambil-api/middlewares"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type merchantHandler struct {
	merchantService Service
}

func NewMerchantHandler(v1 *gin.RouterGroup, merchantService Service) {

	handler := &merchantHandler{merchantService}

	merchant := v1.Group("merchant")
	merchant.GET("", handler.GetAll)

	merchant.Use(middlewares.JwtAuthMiddleware())

	merchant.POST("", handler.Create)
	merchant.GET("/:id", handler.GetDetail)
	merchant.POST("/:id", handler.Update)
}

// @Summary Get All Merchant
// @Description Get All Merchant
// @Accept  json
// @Param MerchantFilterRequest query domain.MerchantFilterRequest true " MerchantFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantData}
// @Router /api/v1/merchant [get]
// @Tags Merchant
func (h *merchantHandler) GetAll(c *gin.Context) {
	start := time.Now()
	input := domain.MerchantFilterRequest{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	merchants, err := h.merchantService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchants,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create Merchant
// @Description Create Merchant
// @Accept  json
// @Param MerchantRequest body domain.MerchantRequest true " MerchantRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantData}
// @Router /api/v1/merchant [post]
// @Tags Merchant
func (h *merchantHandler) Create(c *gin.Context) {
	start := time.Now()
	merchantInput := domain.MerchantRequest{}

	err := c.ShouldBindJSON(&merchantInput)
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

	merchant, err := h.merchantService.Store(merchantInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchant,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get Detail Merchant
// @Description Get Detail Merchant
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param merchant_id path string true " Merchant Id "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantData}
// @Router /api/v1/merchant/{merchant_id} [get]
// @Tags Merchant
func (h *merchantHandler) GetDetail(c *gin.Context) {
	start := time.Now()
	merchantId := c.Param("merchant_id")

	result, err := h.merchantService.GetDetail(merchantId)
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

// @Summary Update Merchant
// @Description Update Merchant
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param merchant_id path string true " Merchant Id "
// @Param MerchantData body domain.MerchantData true " MerchantData Schema "
// @Produce  json
// @Success 200 {object} domain.Response
// @Router /api/v1/merchant/{merchant_id} [post]
// @Tags Merchant
func (h *merchantHandler) Update(c *gin.Context) {
	start := time.Now()
	merchantId := c.Param("merchant_id")
	merchantRequest := domain.MerchantRequest{}

	c.ShouldBindJSON(&merchantRequest)

	err := h.merchantService.Update(merchantId, merchantRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message:     "driver updated !",
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
