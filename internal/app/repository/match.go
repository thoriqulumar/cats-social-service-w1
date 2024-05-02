package repository

import (
	"context"
	"fmt"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

var (
	defaultColumnForMatchTable = ` "id", "issuedId", "receiverId", "matchCatId", "userCatId", "message", "status", "createdAt" `
)

var (
	createMatchQuery = `INSERT INTO "match" ("issuedId","receiverId","matchCatId", "userCatId", "message", "status", "createdAt") 
                        VALUES($1, $2, $3, $4, $5, false, NOW()) RETURNING *;` // Select all inserted columns with *
)

func (r *Repo) MatchCat(ctx context.Context, data model.MatchRequest, issuedId, receiverID int64) (model.Match, error) {
	var match model.Match
	err := r.db.QueryRowContext(ctx, createMatchQuery, issuedId, receiverID, data.MatchCatId, data.UserCatId, data.Message).Scan(&match.ID,
		&match.IssuedID, &match.MatchCatId, &match.UserCatId, &match.Message, &match.Status, &match.CreatedAt, // Scan into struct fields
	)

	fmt.Println(err)
	if err != nil {
		return model.Match{}, err
	}
	return match, nil
}

var (
	getMatchByIDQuery = `SELECT * FROM "match" WHERE "id"=$1 LIMIT 1;`
)

func (r *Repo) GetMatchByID(ctx context.Context, id int64) (data model.Match, err error) {
	err = r.db.QueryRowxContext(ctx, getUserByEmailQuery, id).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	getMatchByIdAndIssuedIdQuery = `SELECT * FROM "match" WHERE "id"=$1 AND "issuedId"=$2 LIMIT 1;`
)

func (r *Repo) GetMatchByIdAndIssuedId(ctx context.Context, id, issuedId int64) (data model.Match, err error) {
	err = r.db.QueryRowxContext(ctx, getMatchByIdAndIssuedIdQuery, id, issuedId).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	deleteMatchQuery = `DELETE FROM "match" WHERE id = $1;`
)

func (r *Repo) DeleteMatchById(ctx context.Context, id int64) (err error) {
	_, err = r.db.ExecContext(ctx, deleteMatchQuery, id)
	if err != nil {
		return err
	}
	return nil
}

var (
	updateMatchStatusQuery = `UPDATE "match"
SET status = $2 WHERE id = $1;`
)

func (r *Repo) UpdateMatchStatus(ctx context.Context, id int64, status model.MatchStatus) (err error) {
	_, err = r.db.ExecContext(ctx, updateMatchStatusQuery, id, status)
	return
}

var (
	getMatchByBothOwnerQuery = `SELECT` + defaultColumnForMatchTable + `FROM "match" WHERE issuedId=$1 AND receiverId=$2;`
)

func (r *Repo) GetMatchByBothOwner(ctx context.Context, issuedId int64, receiverID int64) (listData []model.Match, err error) {
	rows, err := r.db.QueryxContext(ctx, getMatchByBothOwnerQuery, issuedId, receiverID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var match model.Match
		err = rows.StructScan(&match)
		if err != nil {
			return
		}
		listData = append(listData, match)
	}

	return
}
