package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

func (h *Handler) GetCat(c *gin.Context) {
	ctx := c.Request.Context()

	rawQuery := c.Request.URL.RawQuery

	catRequest := parseCatRequestFromQuery(rawQuery)

	data, err := h.service.GetCat(ctx, catRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusCreated, model.GetCatResponse{
		Message: "success",
		Data:    data,
	})
}
