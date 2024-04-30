package model

type Cat struct {
	ID               int64    `json:"id"`
	OwnerId          string   `json:"ownerId"`
	Name             string   `json:"name"`
	Race             string   `json:"race"`
	Sex              string   `json:"sex"`
	AgeInMonth       int      `json:"ageInMonth"`
	Description      string   `json:"description"`
	imagesUrl        []string `json:"imagesUrl"`
	IsAlreadyMatched bool     `json:"isAlreadyMatched"`
	isDeleted        bool     `json:"isDeleted"`
	CreatedAt        string   `json:"createdAt"`
}
