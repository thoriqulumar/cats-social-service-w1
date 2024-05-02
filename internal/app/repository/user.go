package repository

import (
	"context"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

var (
	createUserQuery = `INSERT INTO "user" ("email", "name", "password", "createdAt") VALUES($1, $2, $3, NOW()) RETURNING *;`
)

func (r *Repo) CreateUser(ctx context.Context, data model.User) (user model.User, err error) {
	// db query goes here
	err = r.db.QueryRowxContext(ctx, createUserQuery, data.Email, data.Name, data.Password).StructScan(&user)

	return user, err
}

var (
	getUserByEmailQuery = `SELECT "id","email","name","password" FROM "user" WHERE "email"=$1 LIMIT 1;`
)

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (data model.User, err error) {
	err = r.db.QueryRowxContext(ctx, getUserByEmailQuery, email).StructScan(&data)
	if err != nil {
		return
	}
	return
}

var (
	getUserByIdQuery = `SELECT "id","email","name","createdAt" FROM "user" WHERE "id"=$1 LIMIT 1;`
)

func (r *Repo) GetUserById(ctx context.Context, id int64) (data model.UserResponse, err error) {
	err = r.db.QueryRowxContext(ctx, getUserByIdQuery, id).StructScan(&data)
	if err != nil {
		return
	}
	return
}
