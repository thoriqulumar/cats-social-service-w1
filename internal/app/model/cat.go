package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type RegisterCatResponse struct {
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Cat struct {
	IDStr       string      `json:"id"`
	ID          int64       `json:"-" db:"id"`
	OwnerId     int64       `json:"ownerId" db:"ownerId"`
	Name        string      `json:"name" db:"name"`
	Race        string      `json:"race" db:"race"`
	Sex         string      `json:"sex" db:"sex"`
	AgeInMonth  int         `json:"ageInMonth" db:"ageInMonth"`
	Description string      `json:"description" db:"description"`
	ImagesUrls  StringArray `json:"imageUrls" db:"imageUrls"`
	HasMatched  bool        `json:"hasMatched" db:"hasMatched"`
	IsDeleted   bool        `json:"isDeleted" db:"isDeleted"`
	CreatedAt   string      `json:"createdAt" db:"createdAt"`
}

type GetCatRequest struct {
	ID         string `json:"id,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	Race       string `json:"race,omitempty"`
	Sex        string `json:"sex,omitempty"`
	HasMatched *bool  `json:"hasMatched,omitempty"`
	AgeInMonth string `json:"ageInMonth,omitempty"`
	Owned      *bool  `json:"owned,omitempty"`
	Search     string `json:"search,omitempty"`
}

type GetCatResponse struct {
	Message string `json:"message"`
	Data    []Cat  `json:"data"`
}

type PostCatRequest struct {
	Name        string      `json:"name"`
	Race        string      `json:"race"`
	Sex         string      `json:"sex"`
	AgeInMonth  int         `json:"ageInMonth"`
	Description string      `json:"description"`
	ImageUrls   StringArray `json:"imageUrls"`
}

type Data struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type PostCatResponse struct {
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type PutCatResponse struct {
	Message string `json:"message"`
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
