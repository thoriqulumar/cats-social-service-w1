package model

type MatchRequest struct {
	MatchCatId string `json:"matchCatId"`
	UserCatId  string `json:"userCatId"`
	Message    string `json:"message"`
}

type UpdateStatusRequest struct {
	MatchCatId int64 `json:"matchCatId"`
}

type MatchStatus string

var (
	MatchStatusWaitingForApproval = MatchStatus("waiting_for_approval")
	MatchStatusApproved           = MatchStatus("approved")
	MatchStatusRejected           = MatchStatus("rejected")
	MatchStatusDeleted            = MatchStatus("deleted")
)

type Match struct {
	IDStr      string      `json:"id"`
	ID         int64       `json:"-" db:"id"`
	IssuedID   int64       `json:"issuedId" db:"issuedId"`
	ReceiverID int64       `json:"-" db:"receiverId"` // hide on response
	MatchCatId int64       `json:"matchCatId" db:"matchCatId"`
	UserCatId  int64       `json:"userCatId" db:"userCatId"`
	Message    string      `json:"message" db:"message"`
	Status     MatchStatus `json:"status" db:"status"`
	CreatedAt  string      `json:"createdAt" db:"createdAt"`
}

type MatchResponse struct {
	Message string `json:"message"`
	Data    Match  `json:"data"`
}

type MatchListResponse struct {
	Message string      `json:"message"`
	Data    []MatchData `json:"data"`
}

type MatchData struct {
	ID             string       `json:"id"`
	IssuedBy       UserResponse `json:"issuedBy"`
	MatchCatDetail Cat          `json:"matchCatDetail"`
	UserCatDetail  Cat          `json:"userCatDetail"`
	Message        string       `json:"message"`
	CreatedAt      string       `json:"createdAt"`
}
