package database

import (
	"context"
	"errors"

	"alior-auth/src/types"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	_ "alior-auth/migrations"
)

func NewDB(ctx context.Context, connString string) (*DB, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	return &DB{pool: pool}, MigrateDatabase(pool)
}

type DB struct {
	pool *pgxpool.Pool
}

func (d *DB) Insert(ctx context.Context, user *types.User) (int32, error) {
	var id int32
	return id, d.pool.QueryRow(ctx, `INSERT INTO auth.public.users
    (full_name, email, password, phone_number) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.FullName, user.Email, user.Password, user.PhoneNumber).Scan(&id)
}

func (d *DB) GetUserByID(ctx context.Context, id int32) (*types.User, error) {
	query := `SELECT email, full_name, email, password FROM auth.public.users WHERE id = $1`
	user := &types.User{}

	err := d.pool.QueryRow(ctx, query, id).Scan(&user.Email, &user.FullName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *DB) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	query := `SELECT id, full_name, email, password FROM auth.public.users WHERE email = $1`
	user := &types.User{}

	err := d.pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.FullName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *DB) CheckPassword(ctx context.Context, email, password string) (int32, error) {
	user, err := d.GetUserByEmail(ctx, email)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, errors.New("incorrect password")
	}

	return user.ID, nil
}
