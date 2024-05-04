package repository

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
	"strconv"
)

var (
	createCatQuery = `INSERT INTO "cat" ("name", "race", "sex", "ageInMonth", "description", "imageUrls", "ownerId", "createdAt", "hasMatched", "isDeleted") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(),false,false) 
	RETURNING "id", "ownerId", "name", "race", "sex", "ageInMonth", "description", "imageUrls", 
	COALESCE("hasMatched", false) AS "hasMatched", COALESCE("isDeleted", false) AS "isDeleted";`
)

func (r *Repo) CreateCat(ctx context.Context, data model.Cat) (cat model.Cat, err error) {

	//konversi model.StringArray menjadi string biasa

	imagesUrlStr := pq.Array(data.ImagesUrls)

	// db query goes here

	err = r.db.QueryRowxContext(ctx, createCatQuery, data.Name, data.Race, data.Sex, data.AgeInMonth, data.Description, imagesUrlStr, data.OwnerId).StructScan(&cat)

	return cat, err
}

var (
	prefixGetCat = `SELECT * FROM cat WHERE 1=1 AND "isDeleted"=false`
	suffixGetCat = ` ;`
)

func (r *Repo) GetCat(ctx context.Context, query string, args []interface{}) (cats []model.Cat, err error) {
	concatenatedQuery := prefixGetCat + query + suffixGetCat

	rows, err := r.db.QueryxContext(ctx, replacePlaceholders(concatenatedQuery), args...)
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
		cat.IDStr = strconv.FormatInt(cat.ID, 10)
		cats = append(cats, cat)
	}

	return cats, nil
}

var (
	getCatByID = `SELECT * FROM "cat" WHERE "id"=$1 AND "isDeleted"=false LIMIT 1;`
)

func (r *Repo) GetCatByID(ctx context.Context, id int64) (data model.Cat, err error) {
	err = r.db.QueryRowxContext(ctx, getCatByID, id).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	getCatOwnerByID = `SELECT * FROM "cat" WHERE "id"=$1 AND "ownerId"=$2 AND "isDeleted"=false LIMIT 1;`
)

func (r *Repo) GetCatOwnerByID(ctx context.Context, catId, ownerId int64) (data model.Cat, err error) {
	err = r.db.QueryRowxContext(ctx, getCatOwnerByID, catId, ownerId).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	updateCat = `UPDATE cat SET name=$1, race=$2, sex=$3, "ageInMonth"=$4, description =$5, "imageUrls"=$6 WHERE id=$7;`
)

func (r *Repo) PutCat(ctx context.Context, args []interface{}) (sql.Result, error) {
	result, err := r.db.ExecContext(ctx, updateCat, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

var (
	updateIsDeletedCat = `UPDATE cat SET "isDeleted"=true WHERE id=$1;`
)

func (r *Repo) DeleteCatById(ctx context.Context, id int64) (err error) {
	_, err = r.db.ExecContext(ctx, updateIsDeletedCat, id)
	if err != nil {
		return err
	}
	return nil
}
