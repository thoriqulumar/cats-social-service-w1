package delivery

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

func getRequestedUserIDFromRequest(c *gin.Context) int64 {
	userData, ok := c.Get("userData")
	if !ok {
		return 0
	}
	return int64(userData.(map[string]any)["id"].(float64))
}

func getRequestedUserDataFromRequest(c *gin.Context) (user model.User) {
	userData, ok := c.Get("userData")
	if !ok {
		return user
	}
	tmp, err := json.Marshal(userData)
	if err != nil {
		return user
	}
	err = json.Unmarshal(tmp, &user)
	if err != nil {
		return user
	}
	return
}
