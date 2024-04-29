package repository

import (
	"context"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/model"
)

var (
	createUserQuery = `INSERT INTO "user" (email,"name", "password") VALUES($1, $2, $3);`
)

func (r *Repo) CreateUser(ctx context.Context, data model.User) (err error) {
	// db query goes here
	_, err = r.db.ExecContext(ctx, createUserQuery, data.Email, data.Name, data.Password)
	if err != nil {
		return err
	}
	return nil
}

var (
	getUserByEmailQuery = `SELECT email,name,password FROM "user" WHERE "email"=$1 LIMIT 1;`
)

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (data model.User, err error) {
	err = r.db.QueryRowxContext(ctx, getUserByEmailQuery, email).StructScan(&data)
	if err != nil {
		return
	}
	return
}
