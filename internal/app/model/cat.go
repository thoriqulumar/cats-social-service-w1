package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type RegisterCatResponse struct {
	Message string `json:"message"`
	Data    Cat    `json:"data"`
}

type Cat struct {
	ID               int64       `json:"id" db:"id"`
	OwnerId          int64       `json:"ownerId" db:"ownerId"`
	Name             string      `json:"name" db:"name"`
	Race             string      `json:"race" db:"race"`
	Sex              string      `json:"sex" db:"sex"`
	AgeInMonth       int         `json:"ageInMonth" db:"ageInMonth"`
	Description      string      `json:"description" db:"description"`
	ImagesUrl        StringArray `json:"imagesUrl" db:"imageUrls"`
	IsAlreadyMatched bool        `json:"isAlreadyMatched" db:"isAlreadyMatched"`
	IsDeleted        bool        `json:"isDeleted" db:"isDeleted"`
	CreatedAt        string      `json:"createdAt" db:"createdAt"`
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
