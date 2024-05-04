package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	cerror "github.com/thoriqulumar/cats-social-service-w1/internal/pkg/error"
)

func (h *Handler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	user := model.User{}
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// validation
	err = h.service.ValidateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	data, err := h.service.Register(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusCreated, model.RegisterResponse{
		Message: "User logged successfully",
		Data:    data,
	})
}

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	lr := model.LoginRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&lr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	data, err := h.service.Login(ctx, lr)
	if err != nil {
		c.JSON(cerror.GetCode(err), gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.RegisterResponse{
		Message: "User logged successfully",
		Data:    data,
	})
}
