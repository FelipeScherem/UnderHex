package rest

import (
	"UnderProject/internal/core/models"
	"UnderProject/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) UserCreate(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UserCreate(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UserGetById(c *gin.Context) {

	c.JSON(http.StatusOK, 1)
}

func (h *UserHandler) UserGetAll(c *gin.Context) {

	c.JSON(http.StatusOK, 1)
}

func (h *UserHandler) UserUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, 1)
}

func (h *UserHandler) UserDelete(c *gin.Context) {
	var userID int
	if err := c.BindQuery(&userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userID)
}

// TODO: IMPLEMENTAR A FUNÇÃO
func bindAndValidateJSON(c *gin.Context) (models.User, error) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return models.User{}, err
	}
	return user, nil
}
