package user

import (
	"ambil-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Service
}

func NewUserHandler(v1 *gin.RouterGroup, userService Service) {

	handler := &userHandler{userService}

	v1.POST("/login", handler.Login)
	v1.POST("/register", handler.PostUser)
}

func (h *userHandler) Login(c *gin.Context) {
	var input domain.AuthRequest

	c.ShouldBindJSON(&input)

	token, err := h.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
	}

	c.JSON(http.StatusOK, domain.Response{
		Data: token,
	})
}
