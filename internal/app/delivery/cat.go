package delivery

import (
	"encoding/json"
	"fmt"
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
			ID: data.ID,

			CreatedAt: data.CreatedAt,
		},
	})
}
func (h *Handler) PostCat(c *gin.Context) {
	ctx := c.Request.Context()

	catReq := model.PostCatRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&catReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	userId := getRequestedUserIDFromRequest(c)

	// err = h.service.ValidatePostCat(ctx, catReq, userId)
	// if err != nil {
	// 	c.JSON(cerror.GetCode(err), gin.H{
	// 		"err": err.Error(),
	// 	})
	// 	return
	// }

	data, err := h.service.PostCat(ctx, catReq, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	id := strconv.Itoa(int(data.ID))

	c.JSON(http.StatusCreated, model.PostCatResponse{
		Message: "success",
		Data: model.Data{
			ID:        id,
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
