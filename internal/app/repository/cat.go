package repository

import (
	"context"
	"database/sql"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

var (
	prefixGetCat = `SELECT * FROM cat WHERE 1=1`
	suffixGetCat = `;`
)

func (r *Repo) GetCat(ctx context.Context, query string, args []interface{}) (cats []model.Cat, err error) {
	concatenatedQuery := prefixGetCat + query + suffixGetCat

	rows, err := r.db.QueryxContext(ctx, concatenatedQuery, args...)
	if err != nil {
		return []model.Cat{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var cat model.Cat
		err = rows.StructScan(&cat)
		if err != nil {
			return []model.Cat{}, err
		}
		cats = append(cats, cat)
	}

	return cats, nil
}

var (
	getCatByID = `SELECT * FROM "cat" WHERE "id"=$1 LIMIT 1;`
)

func (r *Repo) GetCatByID(ctx context.Context, id int64) (data model.Cat, err error) {
	err = r.db.QueryRowxContext(ctx, getCatByID, id).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	getCatOwnerByID = `SELECT * FROM "cat" WHERE "id"=$1 AND "ownerId"=$2 LIMIT 1;`
)

func (r *Repo) GetCatOwnerByID(ctx context.Context, catId, ownerId int64) (data model.Cat, err error) {
	err = r.db.QueryRowxContext(ctx, getCatOwnerByID, catId, ownerId).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	updateCat = `UPDATE cat SET name = $1, race = $2, sex = $3, ageInMonth = $4, description = $5, imageUrls = $6 WHERE id = $7`
)

func (r *Repo) PutCat(ctx context.Context, args []interface{}) (sql.Result, error) {
	result, err := r.db.ExecContext(ctx, updateCat, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// var (
//     updateCatIsMatched = `UPDATE "cat" SET "hasMatched"=true WHERE "id"=$1 AND "ownerId"=$2;`
// )

// func (r *Repo) UpdateCatIsMatched(ctx context.Context, catId, ownerId int64) (err error) {
//     result, err := r.db.ExecContext(ctx, updateCatIsMatched, catId, ownerId)
//     if err != nil {
//         return err
//     }
//     rowsAffected, err := result.RowsAffected()
//     if err != nil {
//         return err
//     }
//     if rowsAffected == 0 {
//         // Handle case where no rows were affected (cat not found or owner mismatch)
//         return fmt.Errorf("cat with ID %d and owner ID %d not found", catId, ownerId)
//     }
//     return nil
// }
