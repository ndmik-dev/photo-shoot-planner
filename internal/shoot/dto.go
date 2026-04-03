package shoot

import "time"

type CreateShootRequest struct {
	Title       string    `json:"title" validate:"required,min=3,max=200"`
	Description string    `json:"description" validate:"max=200"`
	Location    string    `json:"location" validate:"required,max=255"`
	Camera      string    `json:"camera" validate:"required,max=100"`
	Lens        string    `json:"lens" validate:"required,max=100"`
	Status      string    `json:"status" validate:"required,oneof=planned shot edited published archived"`
	ShootDate   time.Time `json:"shoot_date" validate:"required"`
}

type UpdateShootRequest struct {
	Title       string    `json:"title" validate:"required,min=3,max=200"`
	Description string    `json:"description" validate:"max=2000"`
	Location    string    `json:"location" validate:"required,max=255"`
	Camera      string    `json:"camera" validate:"required,max=100"`
	Lens        string    `json:"lens" validate:"required,max=100"`
	Status      string    `json:"status" validate:"required,oneof=planned shot edited published archived"`
	ShootDate   time.Time `json:"shoot_date" validate:"required"`
}

type PathShootRequest struct {
	Status string `json:"status" validate:"required,oneof=planned shot edited published archived"`
}

type ShootResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Camera      string    `json:"camera"`
	Lens        string    `json:"lens"`
	Status      string    `json:"status"`
	ShootDate   time.Time `json:"shoot_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
