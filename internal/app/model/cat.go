package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Cat struct {
	ID          int64       `json:"id" db:"id"`
	OwnerId     int64       `json:"ownerId" db:"ownerId"`
	Name        string      `json:"name" db:"name"`
	Race        string      `json:"race" db:"race"`
	Sex         string      `json:"sex" db:"sex"`
	AgeInMonth  int         `json:"ageInMonth" db:"ageInMonth"`
	Description string      `json:"description" db:"description"`
	ImagesUrl   StringArray `json:"imagesUrl" db:"imageUrls"`
	HasMatched  bool        `json:"hasMatched" db:"hasMatched"`
	IsDelseted  bool        `json:"isDeleted" db:"isDeleted"`
	CreatedAt   string      `json:"createdAt" db:"createdAt"`
}

type GetCatRequest struct {
	ID         *string `json:"id,omitempty"`
	Limit      *int    `json:"limit,omitempty"`
	Offset     *int    `json:"offset,omitempty"`
	Race       *string `json:"race,omitempty"`
	Sex        *string `json:"sex,omitempty"`
	HasMatched *bool   `json:"hasMatched,omitempty"`
	AgeInMonth *string `json:"ageInMonth,omitempty"`
	Owned      *bool   `json:"owned,omitempty"`
	Search     *string `json:"search,omitempty"`
}

type GetCatResponse struct {
	Message string `json:"message"`
	Data    []Cat  `json:"data"`
}

// StringArray represents a string array that can be scanned from the database.
type StringArray []string

// Scan scans a database value into StringArray.
func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}
	if bv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := bv.([]byte); ok {
			*s = strings.Split(string(v), ",") // Assuming the array is stored as a comma-separated string
			return nil
		}
	}
	return fmt.Errorf("failed to scan StringArray")
}
