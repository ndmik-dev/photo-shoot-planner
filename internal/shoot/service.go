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

func (s *Service) List(ctx context.Context, limit, offset int32) ([]Shoot, error) {
	return s.repo.List(ctx, limit, offset)
}
