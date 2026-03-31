package shoot

import "time"

type Status string

const (
	StatusPlanned   Status = "planned"
	StatusShot      Status = "shot"
	StatusEdited    Status = "edited"
	StatusPublished Status = "published"
	StatusArchived  Status = "archived"
)

type Shoot struct {
	ID          int64
	Title       string
	Description string
	Location    string
	Camera      string
	Lens        string
	Status      Status
	ShootDate   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func isValidStatus(s string) bool {
	switch Status(s) {
	case StatusPlanned, StatusShot, StatusEdited, StatusPublished, StatusArchived:
		return true
	default:
		return false
	}
}
