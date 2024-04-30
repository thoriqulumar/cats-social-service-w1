package repository

import (
	"context"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

var (
	createMatchQuery = `INSERT INTO "match" ("issuedId","matchCatId", "userCatId", "message", "createdAt") VALUES($1, $2, $3, $4, NOW());`
)

func (r *Repo) MatchCat(ctx context.Context, data model.MatchRequest, issuedId int64) (err error) {
	// db query goes here
	_, err = r.db.ExecContext(ctx, createMatchQuery, issuedId, data.MatchCatId, data.UserCatId, data.Message)
	if err != nil {
		return err
	}
	return nil
}

var(
	getMatchByIdQuery = `SELECT * FROM "match" WHERE "id"=$1 LIMIT 1;`
)

func (r *Repo) GetMatchById(ctx context.Context, id int) (data model.Match, err error){
	err = r.db.QueryRowxContext(ctx, getUserByEmailQuery, id).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var(
	getCatOwnerById = `SELECT * FROM "cat" WHERE "id"=$1 AND "ownerId"=$3 LIMIT 1;`
)

func (r *Repo) GetCatOwnerById(ctx context.Context, catId, ownerId int64) (data model.Cat, err error){
	err = r.db.QueryRowxContext(ctx, getCatOwnerById, catId, ownerId).StructScan(&data)
	if err != nil {
		return
	}
	return
}