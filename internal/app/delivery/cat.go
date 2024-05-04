package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (h *Handler) PutCat(c *gin.Context) {
	ctx := c.Request.Context()

	catReq := model.PostCatRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&catReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	int64Id := int64(intId)

	userId := getRequestedUserIDFromRequest(c)

	err = h.service.ValidatePutCat(ctx, catReq, int64Id, userId)
	if err != nil {
		c.JSON(cerror.GetCode(err), gin.H{
			"err": err.Error(),
		})
		return
	}

	_, err = h.service.PutCat(ctx, catReq, int64Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.PutCatResponse{
		Message: "successfully update cat",
	})
}
