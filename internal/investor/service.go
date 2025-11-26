package investor

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

func (s *Service) Create(ctx context.Context, inv *domain.Investor) error {
	// TODO: Validate investor
	// TODO: Check for duplicates
	return s.repo.Create(ctx, inv)
}

func (s *Service) GetByID(ctx context.Context, id string) (*domain.Investor, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) GetByWallet(ctx context.Context, wallet string) (*domain.Investor, error) {
	return s.repo.FindByWallet(ctx, wallet)
}
