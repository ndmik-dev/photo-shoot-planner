package shoot

import (
	"context"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateShootRequest) (Shoot, error) {
	req.Title = strings.TrimSpace(req.Title)
	req.Location = strings.TrimSpace(req.Location)
	req.Camera = strings.TrimSpace(req.Camera)
	req.Lens = strings.TrimSpace(req.Lens)

	return s.repo.Create(ctx, req)
}

func (s *Service) GetByID(ctx context.Context, id int64) (Shoot, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context, limit, offset int32) ([]Shoot, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *Service) Update(ctx context.Context, id int64, req UpdateShootRequest) (Shoot, error) {
	req.Title = strings.TrimSpace(req.Title)
	req.Location = strings.TrimSpace(req.Location)
	req.Camera = strings.TrimSpace(req.Camera)
	req.Lens = strings.TrimSpace(req.Lens)

	return s.repo.Update(ctx, id, req)
}

func (s *Service) UpdateStatus(ctx context.Context, id int64, status string) (Shoot, error) {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
