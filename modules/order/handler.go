package order

import (
	"ambil-api/domain"
	"ambil-api/middlewares"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type orderHandler struct {
	orderService Service
}

func NewOrderHandler(v1 *gin.RouterGroup, orderService Service) {

	handler := &orderHandler{orderService}

	order := v1.Group("order")

	order.Use(middlewares.JwtAuthMiddleware())

	order.GET("", handler.GetAll)
	order.POST("", handler.Create)
	order.GET("/:id", handler.GetDetail)
	order.POST("/:id", handler.Update)
}

// @Summary Get All Order
// @Description Get All Order
// @Accept  json
// @Param OrderFilterRequest query domain.OrderFilterRequest true " OrderFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.OrderData}
// @Router /api/v1/order [get]
// @Tags Order
func (h *orderHandler) GetAll(c *gin.Context) {
	start := time.Now()
	input := domain.OrderFilterRequest{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	orders, err := h.orderService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        orders,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create Order
// @Description Create Order
// @Accept  json
// @Param OrderRequest body domain.OrderRequest true " OrderRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.OrderData}
// @Router /api/v1/order [post]
// @Tags Order
func (h *orderHandler) Create(c *gin.Context) {
	start := time.Now()
	orderInput := domain.OrderRequest{}

	err := c.ShouldBindJSON(&orderInput)
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

	order, err := h.orderService.Store(orderInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        order,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get Detail Order
// @Description Get Detail Order
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param order_id path string true " Order Id "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.OrderData}
// @Router /api/v1/order/{order_id} [get]
// @Tags Order
func (h *orderHandler) GetDetail(c *gin.Context) {
	start := time.Now()
	orderId := c.Param("order_id")

	result, err := h.orderService.GetDetail(orderId)
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

// @Summary Update Order
// @Description Update Order
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param order_id path string true " Order Id "
// @Param OrderData body domain.OrderData true " OrderData Schema "
// @Produce  json
// @Success 200 {object} domain.Response
// @Router /api/v1/order/{order_id} [post]
// @Tags Order
func (h *orderHandler) Update(c *gin.Context) {
	start := time.Now()
	orderId := c.Param("order_id")
	orderRequest := domain.OrderRequest{}

	c.ShouldBindJSON(&orderRequest)

	err := h.orderService.Update(orderId, orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message:     "order updated !",
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
