package delivery

import (
	"encoding/json"
	"strconv"
	"strings"

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

func parseCatRequestFromQuery(rawQuery string) model.GetCatRequest {
	request := model.GetCatRequest{}

	if rawQuery == "" {
		return request
	}

	queryParams := strings.Split(rawQuery, "&")

	for _, param := range queryParams {
		parts := strings.SplitN(param, "=", 2)
		key := parts[0]
		value := parts[1]

		switch key {
		case "id":
			request.ID = &value
		case "limit":
			limit, _ := strconv.Atoi(value)
			request.Limit = &limit
		case "offset":
			offset, _ := strconv.Atoi(value)
			request.Offset = &offset
		case "race":
			request.Race = &value
		case "sex":
			request.Sex = &value
		case "hasMatched":
			hasMatched, _ := strconv.ParseBool(value)
			request.HasMatched = &hasMatched
		case "ageInMonth":
			request.AgeInMonth = &value
		case "owned":
			owned, _ := strconv.ParseBool(value)
			request.Owned = owned
		case "search":
			request.Search = &value
		}
	}

	return request
}
