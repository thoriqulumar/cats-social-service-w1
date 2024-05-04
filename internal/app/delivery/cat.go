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

	catRequest, err := parseCatRequestFromQuery(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

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
	data, err := h.service.RegisterCat(ctx, cat, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusCreated, model.RegisterCatResponse{
		Message: "success",
		Data: model.Data{
			ID:        data.IDStr,
			CreatedAt: data.CreatedAt,
		},
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

func (h *Handler) DeleteCat(c *gin.Context) {
	ctx := c.Request.Context()

	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	ownerId := getRequestedUserIDFromRequest(c)

	err := h.service.ValidateDeleteCat(ctx, int64(id), int64(ownerId))
	if err != nil {
		c.JSON(cerror.GetCode(err), gin.H{
			"err": err.Error(),
		})
		return
	}

	err = h.service.DeleteCat(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetCatResponse{
		Message: "Match deleted successfully",
	})
}
