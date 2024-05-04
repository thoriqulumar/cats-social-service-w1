package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	cerror "github.com/thoriqulumar/cats-social-service-w1/internal/pkg/error"
)

func (h *Handler) GetCat(c *gin.Context) {
	ctx := c.Request.Context()

	rawQuery := c.Request.URL.RawQuery

	catRequest := parseCatRequestFromQuery(rawQuery)

	userId := getRequestedUserIDFromRequest(c)

	data, err := h.service.GetCat(ctx, catRequest, userId)
	if err != nil {
		c.JSON(cerror.GetCode(err), gin.H{
			"err": err.Error(),
		})
		return
	}

	if data == nil {
		data = []model.Cat{}
	}

	c.JSON(http.StatusOK, model.GetCatResponse{
		Message: "success",
		Data:    data,
	})
}

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
	userId := getRequestedUserIDFromRequest(c)
	fmt.Println("Extracted userId:", userId)
	data, err := h.service.RegisterCat(ctx, cat, userId)
	fmt.Println("Deliver Data", data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, model.RegisterCatResponse{
		Message: "success",
		Data: struct {
			ID        int64  `json:"id"`
			CreatedAt string `json:"createdAt"`
		}{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
		},
	})
}
