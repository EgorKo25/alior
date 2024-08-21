package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "alior-sms/migrations" // линтер жалуется, но это нужно, чтобы MigrateDatabase работал
	"alior-sms/src/types"
)

type DB struct {
	pool *pgxpool.Pool
}

type ConfigDatabase interface {
	GetConnString() string
}

func NewDB(ctx context.Context, cfgDB ConfigDatabase, migrationDir string) (*DB, error) {
	pool, err := pgxpool.New(ctx, cfgDB.GetConnString())
	if err != nil {
		return nil, err
	}

	if err := MigrateDatabase(pool, migrationDir); err != nil {
		pool.Close()
		return nil, err
	}

	return &DB{pool: pool}, nil
}

func (db *DB) InsertService(ctx context.Context, service *types.Service) (int32, error) {
	query := `INSERT INTO services
    (name, description, price) VALUES ($1, $2, $3) RETURNING id`
	return service.ID, db.pool.QueryRow(ctx, query, service.Name, service.Description, service.Price).Scan(&service.ID)
}

func (db *DB) GetServiceByID(ctx context.Context, id int32) (*types.Service, error) {
	query := `SELECT name, description, price FROM services WHERE id = $1`
	service := &types.Service{ID: id}

	return service, db.pool.QueryRow(ctx, query, id).Scan(&service.Name, &service.Description, &service.Price)
}

func (db *DB) DelServiceByID(ctx context.Context, id int32) error {
	query := `DELETE FROM services WHERE id = $1`
	commTag, err := db.pool.Exec(ctx, query, id)

	if commTag.RowsAffected() != 1 {
		return errors.New("No row found to delete")
	}

	return err
}

func (db *DB) GetPaginatedServices(ctx context.Context, limit, offset int32) ([]*types.Service, error) {
	query := `SELECT id, name, description, price
	FROM services ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := db.pool.Query(ctx, query, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	services := make([]*types.Service, 0, limit)

	for rows.Next() {
		var service types.Service
		if err := rows.Scan(&service.ID, &service.Name, &service.Description, &service.Price); err != nil {
			return nil, err
		}

		services = append(services, &service)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}

func (db *DB) InsertProject(ctx context.Context, project *types.Project) (int32, error) {
	query := `INSERT INTO projects
    (url, description) VALUES ($1, $2) RETURNING id`
	return project.ID, db.pool.QueryRow(ctx, query, project.URL, project.Description).Scan(&project.ID)
}

func (db *DB) GetProjectByID(ctx context.Context, id int32) (*types.Project, error) {
	query := `SELECT url, description FROM projects WHERE id = $1`
	project := &types.Project{ID: id}

	return project, db.pool.QueryRow(ctx, query, id).Scan(&project.URL, &project.Description)
}

func (db *DB) DelProjectByID(ctx context.Context, id int32) error {
	query := `DELETE FROM projects WHERE id = $1`
	commTag, err := db.pool.Exec(ctx, query, id)

	if commTag.RowsAffected() != 1 {
		return errors.New("No row found to delete")
	}

	return err
}
