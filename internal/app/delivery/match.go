package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	cerror "github.com/thoriqulumar/cats-social-service-w1/internal/pkg/error"
)

func (h *Handler) MatchCat(c *gin.Context) {
	ctx := c.Request.Context()
	match := model.MatchRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&match)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	issuedId := getRequestedUserIDFromRequest(c)

	err = h.service.ValidateMatchCat(ctx, match, int64(issuedId))
	if err != nil {
		c.JSON(cerror.GetCode(err), gin.H{
			"err": err.Error(),
		})
		return
	}

	// create match
	data, err := h.service.MatchCat(ctx, match, int64(issuedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusCreated, model.MatchResponse{
		Message: "Cat matched successfully. Please wait for the response of receiver",
		Data:    data,
	})
}

func (h *Handler) DeleteMatch(c *gin.Context) {
	ctx := c.Request.Context()

	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	issuedId := getRequestedUserIDFromRequest(c)

	// create match
	err := h.service.DeleteMatch(ctx, int64(id), int64(issuedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, model.MatchResponse{
		Message: "Match deleted successfully",
	})
}

func (h *Handler) ApproveMatch(c *gin.Context) {
	ctx := c.Request.Context()
	receiverId := getRequestedUserIDFromRequest(c)

	req := model.UpdateStatusRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	matchID, err := h.service.ApproveMatch(ctx, req.MatchCatId, receiverId)
	if err != nil {
		c.JSON(cerror.GetCode(err), gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Match approved successfully",
		"matchId": matchID,
	})
}

func (h *Handler) RejectMatch(c *gin.Context) {
	ctx := c.Request.Context()
	receiverId := getRequestedUserIDFromRequest(c)

	req := model.UpdateStatusRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	matchID, err := h.service.ApproveMatch(ctx, req.MatchCatId, receiverId)
	if err != nil {
		c.JSON(cerror.GetCode(err), gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Match approved successfully",
		"matchId": matchID,
	})
}
