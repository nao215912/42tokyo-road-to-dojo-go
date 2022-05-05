package dao

import (
	"42tokyo-road-to-dojo-go/pkg/database/object"
	"context"
	"database/sql"
	"time"
)

type User interface {
	Create(ctx context.Context, name string) (*object.User, error)
	FindByName(ctx context.Context, name string) (*object.User, error)
	FindByToken(ctx context.Context, token string) (*object.User, error)
	UpdateByName(ctx context.Context, user *object.User) (*object.User, error)
}

type user struct {
	db *sql.DB
}

func (u *user) UpdateByName(ctx context.Context, user *object.User) (*object.User, error) {
	const query = `UPDATE user SET name = ?, updated_at = ? WHERE token = ?
`
	_, err := u.db.ExecContext(ctx, query, user.Name, time.Now(), user.Token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) FindByToken(ctx context.Context, token string) (*object.User, error) {
	const query = `select id, name, token, high_score, coin, created_at, updated_at from users where token = ?`
	var user object.User

	err := StructScan(&user, u.db.QueryRowContext(ctx, query, token))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) Create(ctx context.Context, name string) (*object.User, error) {
	const query = `insert into users (id, name, token) values(UUID(), ?, UUID())`
	_, err := u.db.ExecContext(ctx, query, name)
	if err != nil {
		return nil, err
	}
	return u.FindByName(ctx, name)
}

func (u *user) FindByName(ctx context.Context, name string) (*object.User, error) {
	const query = `select id, name, token, high_score, coin, created_at, updated_at from users where name = ?`
	var user object.User

	err := StructScan(&user, u.db.QueryRowContext(ctx, query, name))
	if err != nil {
		return nil, err
	}
	return &user, nil
}
