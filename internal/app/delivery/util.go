package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"net/http"
	"strconv"
)

func getRequestedUserIDFromRequest(c *gin.Context) int64 {
	userData, ok := c.Get("userData")
	if !ok {
		return 0
	}
	userID, _ := strconv.ParseInt(userData.(map[string]any)["id"].(string), 10, 64)
	return userID
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

func parseCatRequestFromQuery(r *http.Request) (model.GetCatRequest, error) {
	var req model.GetCatRequest

	// Parse query parameters
	err := r.ParseForm()
	if err != nil {
		return req, err
	}

	// Convert form values to struct fields
	req.ID = r.FormValue("id")
	limit := parseFormInt(r.FormValue("limit"))
	if limit == 0 {
		limit = 5 // default limit
	}
	req.Limit = limit
	req.Offset = parseFormInt(r.FormValue("offset"))
	req.Race = r.FormValue("race")
	req.Sex = r.FormValue("sex")
	req.HasMatched = parseFormBoolPtr(r.FormValue("hasMatched"))
	req.AgeInMonth = r.FormValue("ageInMonth")
	req.Owned = parseFormBoolPtr(r.FormValue("owned"))
	req.Search = r.FormValue("search")

	return req, nil
}

func parseFormInt(value string) int {
	// Handle conversion error gracefully
	var result int
	_, err := fmt.Sscanf(value, "%d", &result)
	if err != nil {
		return 0
	}
	return result
}

func parseFormBool(value string) bool {
	// Handle conversion error gracefully
	if value == "true" || value == "1" {
		return true
	}
	return false
}

func parseFormBoolPtr(value string) *bool {
	// Handle conversion error gracefully
	if value == "" {
		return nil
	}
	result := parseFormBool(value)
	return &result
}
