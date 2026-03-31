package shoot

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ndmik-dev/photo-shoot-planner/internal/platform/dbgen"
)

func FromDB(m dbgen.Shoot) Shoot {
	return Shoot{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Location:    m.Location,
		Camera:      m.Camera,
		Lens:        m.Lens,
		Status:      Status(m.Status),
		ShootDate:   toTime(m.ShootDate),
		CreatedAt:   toTime(m.CreatedAt),
		UpdatedAt:   toTime(m.UpdatedAt),
	}
}

func ToResponse(s Shoot) ShootResponse {
	return ShootResponse{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		Location:    s.Location,
		Camera:      s.Camera,
		Lens:        s.Lens,
		Status:      string(s.Status),
		ShootDate:   s.ShootDate,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func toTime(ts pgtype.Timestamptz) time.Time {
	if !ts.Valid {
		return time.Time{}
	}

	return ts.Time
}

func toTimestamptz(t time.Time) pgtype.Timestamptz {
	if t.IsZero() {
		return pgtype.Timestamptz{}
	}

	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}
