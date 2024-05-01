package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
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

	// TODO get issuedId from access token
	mockIssuedId := 1

	// validation create match cat
	err = h.service.ValidationMatchCat(ctx, match, int64(mockIssuedId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// create match
	err = h.service.MatchCat(ctx, match, int64(mockIssuedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, model.MatchResponse{
		Message: "Cat matched successfully. Please wait for the response of receiver",
	})
}

func (h *Handler) DeleteMatch(c *gin.Context) {
	ctx := c.Request.Context()
	
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	// TODO get issuedId from access token
	mockIssuedId := 1

	
	// create match
	err := h.service.DeleteMatch(ctx, int64(id), int64(mockIssuedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, model.MatchResponse{
		Message: "Match deleted successfully",
	})
}


