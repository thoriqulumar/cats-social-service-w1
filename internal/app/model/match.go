package model

type MatchRequest struct {
	MatchCatId int64  `json:"matchCatId"`
	UserCatId  int64  `json:"userCatId"`
	Message    string `json:"message"`
}

type Match struct {
	ID         int64  `json:"id"`
	IssuedID   int64  `json:"issuedId"`
	MatchCatId int64  `json:"matchCatId"`
	UserCatId  int64  `json:"userCatId"`
	Message    string `json:"message"`
	CreatedAt  string `json:"createdAt"`
}

type MatchResponse struct {
	Message string `json:"message"`
	Data    Match  `json:"data"`
}
