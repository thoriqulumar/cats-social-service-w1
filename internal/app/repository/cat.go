package repository

import (
	"context"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

var(
	getCatById = `SELECT * FROM "cat" WHERE "id"=$1 LIMIT 1;`
)

func (r *Repo) GetCatByID(ctx context.Context, id int) (data model.Cat, err error){
	err = r.db.QueryRowxContext(ctx, getCatById, id).StructScan(&data)
	if err != nil {
		return
	}
	return
}