package asset

import (
	"context"

	"Regma/internal/domain"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, asset *domain.Asset) error {
	// TODO: Validate asset
	// TODO: Set default status
	return s.repo.Create(ctx, asset)
}

func (s *Service) GetByID(ctx context.Context, id string) (*domain.Asset, error) {
	return s.repo.FindByID(ctx, id)
}
