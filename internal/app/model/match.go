package model

type MatchRequest struct {
	MatchCatId int64  `json:"matchCatId"`
	UserCatId  int64  `json:"userCatId"`
	Message    string `json:"message"`
}

type Match struct {
	ID                   int64  `json:"id" db:"id"`
	IssuedID             int64  `json:"issuedId" db:"issuedId"`
	MatchCatId           int64  `json:"matchCatId" db:"matchCatId"`
	UserCatId            int64  `json:"userCatId" db:"userCatId"`
	Message              string `json:"message" db:"message"`
	IsApprovedOrRejected bool   `json:"isApprovedOrRejected" db:"isApprovedOrRejected"`
	CreatedAt            string `json:"createdAt" db:"createdAt"`
}

type MatchResponse struct {
	Message string `json:"message"`
	Data    Match  `json:"data"`
}
