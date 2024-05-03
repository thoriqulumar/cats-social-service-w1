package repository

import (
	"context"

	"github.com/lib/pq"

	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

var (
	createCatQuery = `INSERT INTO "cat" ("name", "race", "sex", "ageInMonth", "description", "imageUrls") VALUES($1, $2, $3, $4, $5, $6) RETURNING *;`
)

func (r *Repo) CreateCat(ctx context.Context, data model.Cat) (cat model.Cat, err error) {

	//konversi model.StringArray menjadi string biasa

	imagesUrlStr := pq.Array(data.ImagesUrl)

	// db query goes here

	err = r.db.QueryRowxContext(ctx, createCatQuery, data.Name, data.Race, data.Sex, data.AgeInMonth, data.Description, imagesUrlStr).StructScan(&cat)

	return cat, err
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

// var (
//     updateCatIsMatched = `UPDATE "cat" SET "isAlreadyMatched"=true WHERE "id"=$1 AND "ownerId"=$2;`
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
