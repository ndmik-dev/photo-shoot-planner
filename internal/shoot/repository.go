package shoot

import (
	"context"
	"errors"

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
