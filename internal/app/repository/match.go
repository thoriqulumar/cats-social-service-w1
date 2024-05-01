package repository

import (
	"context"
	"fmt"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

var (
	createMatchQuery = `INSERT INTO "match" ("issuedId","matchCatId", "userCatId", "message", "createdAt") VALUES($1, $2, $3, $4, NOW());`
)

func (r *Repo) MatchCat(ctx context.Context, data model.MatchRequest, issuedId int64) (err error) {
	// db query goes here
	_, err = r.db.ExecContext(ctx, createMatchQuery, issuedId, data.MatchCatId, data.UserCatId, data.Message)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

var (
	getMatchByIDQuery = `SELECT * FROM "match" WHERE "id"=$1 LIMIT 1;`
)

func (r *Repo) GetMatchByID(ctx context.Context, id int) (data model.Match, err error) {
	err = r.db.QueryRowxContext(ctx, getUserByEmailQuery, id).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	getMatchByIdAndIssuedIdQuery = `SELECT * FROM "match" WHERE "id"=$1 AND "issuedId"=$2 LIMIT 1;`
)

func (r *Repo) GetMatchByIdAndIssuedId(ctx context.Context, id, issuedId int64) (data model.Match, err error){
	err = r.db.QueryRowxContext(ctx, getMatchByIdAndIssuedIdQuery, id, issuedId).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	deleteMatchQuery = `DELETE FROM "match" WHERE id = $1;`
)

func (r *Repo) DeleteMatchById(ctx context.Context, id int64) (err error){
	_, err = r.db.ExecContext(ctx, deleteMatchQuery, id)
    if err != nil {
        return err
    }
    return nil
}
