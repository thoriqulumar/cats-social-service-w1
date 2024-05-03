package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

func (h *Handler) RegisterCat(c *gin.Context) {
	ctx := c.Request.Context()
	cat := model.Cat{}
	err := json.NewDecoder(c.Request.Body).Decode(&cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// validation
	err = h.service.ValidateCat(ctx, cat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	data, err := h.service.RegisterCat(ctx, cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, model.RegisterCatResponse{
		Message: "success",
		Data:    data,
	})
}
