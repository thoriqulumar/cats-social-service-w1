package delivery

import (
	"encoding/json"
	"net/http"

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
	// check if userCatId is owned by userId

	// check is matchCatId and userCatId is found  
	err = h.service.ValidateIdCat(ctx, match)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
		})
		return
	}

	// check is matchCatId and userCatId is found 
	err = h.service.ValidateGenderCat(ctx, match)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// create match

	c.JSON(http.StatusOK, model.MatchResponse{
		Message: "User logged successfully",
	})
}
