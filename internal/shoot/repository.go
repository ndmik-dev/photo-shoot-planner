package shoot

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/ndmik-dev/photo-shoot-planner/internal/platform/dbgen"
)

var ErrNotFound = errors.New("shoot not found")

type Repository struct {
	q *dbgen.Queries
}

func NewRepository(q *dbgen.Queries) *Repository {
	return &Repository{q: q}
}

func (r *Repository) Create(ctx context.Context, req CreateShootRequest) (Shoot, error) {
	row, err := r.q.CreateShoot(ctx, dbgen.CreateShootParams{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Camera:      req.Camera,
		Lens:        req.Lens,
		Status:      req.Status,
		ShootDate:   toTimestamptz(req.ShootDate),
	})

	if err != nil {
		return Shoot{}, err
	}

	return FromDB(row), nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (Shoot, error) {
	row, err := r.q.GetShootByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Shoot{}, ErrNotFound
		}
		return Shoot{}, err
	}
	return FromDB(row), nil
}

func (r *Repository) List(ctx context.Context, limit, offset int32) ([]Shoot, error) {
	rows, err := r.q.ListShoots(ctx, dbgen.ListShootsParams{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return nil, err
	}

	result := make([]Shoot, 0, len(rows))
	for _, row := range rows {
		result = append(result, FromDB(row))
	}

	return result, nil
}

func (r *Repository) Update(ctx context.Context, id int64, req UpdateShootRequest) (Shoot, error) {
	row, err := r.q.UpdateShoot(ctx, dbgen.UpdateShootParams{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Camera:      req.Camera,
		Lens:        req.Lens,
		Status:      req.Status,
		ShootDate:   toTimestamptz(req.ShootDate),
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Shoot{}, ErrNotFound
		}
		return Shoot{}, err
	}

	return FromDB(row), nil
}

func (r *Repository) UpdateStatus(ctx context.Context, id int64, status string) (Shoot, error) {
	row, err := r.q.UpdateShootStatus(ctx, dbgen.UpdateShootStatusParams{
		ID:     id,
		Status: status,
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Shoot{}, ErrNotFound
		}
		return Shoot{}, err
	}

	return FromDB(row), nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteShoot(ctx, id)
}
