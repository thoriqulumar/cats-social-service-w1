package model

type MatchRequest struct {
	MatchCatId int64  `json:"matchCatId"`
	UserCatId  int64  `json:"userCatId"`
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
	ID         int64       `json:"id" db:"id"`
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
