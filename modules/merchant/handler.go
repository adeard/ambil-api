package merchant

import (
	"ambil-api/domain"
	"ambil-api/middlewares"
	"fmt"
	"net/http"
	"strconv"
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
	merchant.GET(":id", handler.GetDetail)
	merchant.GET(":id/gallery", handler.GetGalleryByMerchantId)
	merchant.POST("", middlewares.JwtAuthMiddleware(), handler.Create)
	merchant.POST(":id", middlewares.JwtAuthMiddleware(), handler.Update)
	merchant.POST("gallery", middlewares.JwtAuthMiddleware(), handler.AddPhotoGallery)

	item := merchant.Group("item")
	item.GET("", handler.GetAllItem)
	item.Use(middlewares.JwtAuthMiddleware())
	item.POST("", handler.CreateItem)
	item.POST(":id", handler.UpdateItem)

	rating := merchant.Group("rating")
	rating.Use(middlewares.JwtAuthMiddleware())
	rating.POST("", handler.CreateRating)
	rating.POST("image", handler.CreateRatingImage)

	category := merchant.Group("category")
	category.GET("", handler.GetAllCategory)
	category.POST("", middlewares.JwtAuthMiddleware(), handler.CreateCategory)
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
	input := domain.MerchantGalleryFilterRequest{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	merchants, err := h.merchantService.GetAllGallery(input)
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

// @Summary Get All Merchant Gallery
// @Description Get All Merchant Gallery
// @Accept  json
// @Param MerchantGalleryFilterRequest query domain.MerchantGalleryFilterRequest true " MerchantGalleryFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantGalleryData}
// @Router /api/v1/merchant/gallery [get]
// @Tags Merchant
func (h *merchantHandler) GetGalleryByMerchantId(c *gin.Context) {
	start := time.Now()
	merchantId, _ := strconv.Atoi(c.Param("id"))
	input := domain.MerchantGalleryFilterRequest{
		MerchantGalleryRequest: domain.MerchantGalleryRequest{
			MerchantId: merchantId,
		},
	}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	merchantGalleries, err := h.merchantService.GetAllGallery(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchantGalleries,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get All Merchant Category
// @Description Get All Merchant Category
// @Accept  json
// @Param MerchantCategoryFilterRequest query domain.MerchantCategoryFilterRequest true " MerchantCategoryFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantCategoryData}
// @Router /api/v1/merchant/category [get]
// @Tags Merchant
func (h *merchantHandler) GetAllCategory(c *gin.Context) {
	start := time.Now()
	input := domain.MerchantCategoryFilterRequest{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	merchantCategories, err := h.merchantService.GetAllCategory(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchantCategories,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get All Merchant Item
// @Description Get All Merchant Item
// @Accept  json
// @Param MerchantItemFilterRequest query domain.MerchantItemFilterRequest true " MerchantItemFilterRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantItemData}
// @Router /api/v1/merchant/item [get]
// @Tags Merchant
func (h *merchantHandler) GetAllItem(c *gin.Context) {
	start := time.Now()
	input := domain.MerchantItemFilterRequest{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})
		return
	}

	merchantItems, err := h.merchantService.GetAllItem(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchantItems,
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

// @Summary Create Gallery Merchant
// @Description Create Gallery Merchant
// @Accept  json
// @Param MerchantGalleryRequest body domain.MerchantGalleryRequest true " MerchantGalleryRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantGalleryData}
// @Router /api/v1/merchant/gallery [post]
// @Tags Merchant
func (h *merchantHandler) AddPhotoGallery(c *gin.Context) {
	start := time.Now()
	merchantGalleryInput := domain.MerchantGalleryRequest{}

	err := c.ShouldBindJSON(&merchantGalleryInput)
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

	merchantGallery, err := h.merchantService.StoreGallery(merchantGalleryInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchantGallery,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create Category Merchant
// @Description Create Category Merchant
// @Accept  json
// @Param MerchantCategoryRequest body domain.MerchantCategoryRequest true " MerchantCategoryRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantCategoryData}
// @Router /api/v1/merchant/category [post]
// @Tags Merchant
func (h *merchantHandler) CreateCategory(c *gin.Context) {
	start := time.Now()
	merchantCategoryInput := domain.MerchantCategoryRequest{}

	err := c.ShouldBindJSON(&merchantCategoryInput)
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

	merchantCategory, err := h.merchantService.StoreCategory(merchantCategoryInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchantCategory,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create Merchant Item
// @Description Create Merchant Item
// @Accept  json
// @Param MerchantItemRequest body domain.MerchantItemRequest true " MerchantItemRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantItemData}
// @Router /api/v1/merchant/item [post]
// @Tags Merchant
func (h *merchantHandler) CreateItem(c *gin.Context) {
	start := time.Now()
	merchantItemInput := domain.MerchantItemRequest{}

	err := c.ShouldBindJSON(&merchantItemInput)
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

	merchantItem, err := h.merchantService.StoreItem(merchantItemInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchantItem,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create Merchant Rating
// @Description Create Merchant Rating
// @Accept  json
// @Param MerchantRatingRequest body domain.MerchantRatingRequest true " MerchantRatingRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantRatingData}
// @Router /api/v1/merchant/rating [post]
// @Tags Merchant
func (h *merchantHandler) CreateRating(c *gin.Context) {
	start := time.Now()
	merchantRatingInput := domain.MerchantRatingRequest{}

	err := c.ShouldBindJSON(&merchantRatingInput)
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

	merchantRating, err := h.merchantService.StoreRating(merchantRatingInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchantRating,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Create Merchant Rating Image
// @Description Create Merchant Rating Image
// @Accept  json
// @Param MerchantRatingImageRequest body domain.MerchantRatingImageRequest true " MerchantRatingImageRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response{data=domain.MerchantRatingImageData}
// @Router /api/v1/merchant/rating/image [post]
// @Tags Merchant
func (h *merchantHandler) CreateRatingImage(c *gin.Context) {
	start := time.Now()
	merchantRatingImageInput := domain.MerchantRatingImageRequest{}

	err := c.ShouldBindJSON(&merchantRatingImageInput)
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

	merchantRatingImage, err := h.merchantService.StoreRatingImage(merchantRatingImageInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:        merchantRatingImage,
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
	merchantId := c.Param("id")

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
	merchantId := c.Param("id")
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

// @Summary Update Merchant Item
// @Description Update Merchant Item
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param merchant_item_id path string true " Merchant Item Id "
// @Param MerchantItemRequest body domain.MerchantItemRequest true " MerchantItemRequest Schema "
// @Produce  json
// @Success 200 {object} domain.Response
// @Router /api/v1/merchant/item/{merchant_item_id} [post]
// @Tags Merchant
func (h *merchantHandler) UpdateItem(c *gin.Context) {
	start := time.Now()
	merchantItemId := c.Param("id")
	merchantItemRequest := domain.MerchantItemRequest{}

	c.ShouldBindJSON(&merchantItemRequest)

	err := h.merchantService.UpdateItem(merchantItemId, merchantItemRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message:     "item updated !",
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
