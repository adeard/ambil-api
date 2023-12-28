package category

import (
	"ambil-api/domain"
	"ambil-api/middlewares"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type categoryHandler struct {
	categoryService Service
}

func NewCategoryHandler(v1 *gin.RouterGroup, categoryService Service) {

	handler := &categoryHandler{categoryService}

	category := v1.Group("category")
	category.GET("", handler.GetAll)
	category.GET("/:id", handler.GetDetail)

	category.Use(middlewares.JwtAuthMiddleware())

	category.POST("", handler.Create)
	category.POST("/:id", handler.Update)
}

// @Summary Get All Category
// @Description Get All Category
// @Accept  json
// @Param CategoryFilterRequest query domain.CategoryFilterRequest true " CategoryFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.CategoryData}
// @Router /api/v1/category [get]
// @Tags Category
func (h *categoryHandler) GetAll(c *gin.Context) {
	start := time.Now()
	input := domain.CategoryFilterRequest{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	categories, err := h.categoryService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        categories,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create Category
// @Description Create Category
// @Accept  json
// @Param CategoryRequest body domain.CategoryRequest true " CategoryRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.CategoryData}
// @Router /api/v1/category [post]
// @Tags Category
func (h *categoryHandler) Create(c *gin.Context) {
	start := time.Now()
	categoryInput := domain.CategoryRequest{}

	err := c.ShouldBindJSON(&categoryInput)
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

	category, err := h.categoryService.Store(categoryInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        category,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get Detail Category
// @Description Get Detail Category
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param category_id path string true " Category Id "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.CategoryData}
// @Router /api/v1/category/{category_id} [get]
// @Tags Category
func (h *categoryHandler) GetDetail(c *gin.Context) {
	start := time.Now()
	categoryId := c.Param("id")

	result, err := h.categoryService.GetDetail(categoryId)
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

// @Summary Update Category
// @Description Update Category
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param category_id path string true " Category Id "
// @Param CategoryData body domain.CategoryData true " CategoryData Schema "
// @Produce  json
// @Success 200 {object} domain.Response
// @Router /api/v1/category/{category_id} [post]
// @Tags Category
func (h *categoryHandler) Update(c *gin.Context) {
	start := time.Now()
	categoryId := c.Param("id")
	CategoryRequest := domain.CategoryRequest{}

	c.ShouldBindJSON(&CategoryRequest)

	err := h.categoryService.Update(categoryId, CategoryRequest)
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
