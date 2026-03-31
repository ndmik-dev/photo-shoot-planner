-- name: CreateShoot :one
INSERT INTO shoots (
    title,
    description,
    location,
    camera,
    lens,
    status,
    shoot_date
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         )
RETURNING id, title, description, location, camera, lens, status, shoot_date, created_at, updated_at;

-- name: GetShootByID :one
SELECT id, title, description, location, camera, lens, status, shoot_date, created_at, updated_at
FROM shoots
WHERE id = $1;

-- name: DeleteShoot :exec
DELETE FROM shoots
WHERE id = $1;

-- name: ListShoots :many
SELECT id, title, description, location, camera, lens, status, shoot_date, created_at, updated_at
FROM shoots
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateShoot :one
UPDATE shoots
SET
    title = $2,
    description = $3,
    location = $4,
    camera = $5,
    lens = $6,
    status = $7,
    shoot_date = $8,
    updated_at = NOW()
WHERE id = $1
RETURNING id, title, description, location, camera, lens, status, shoot_date, created_at, updated_at;

-- name: UpdateShootStatus :one
UPDATE shoots
SET
    status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, title, description, location, camera, lens, status, shoot_date, created_at, updated_at;