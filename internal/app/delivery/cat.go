package delivery

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCat(c *gin.Context) {
	ctx := c.Request.Context()

	rawQuery := c.Request.URL.RawQuery

	catRequest := parseCatRequestFromQuery(rawQuery)

	h.service.GetCat(ctx, catRequest)
}
