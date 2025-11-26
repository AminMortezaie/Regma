package investor

import (
	"context"

	"Regma/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, inv *domain.Investor) error
	FindByID(ctx context.Context, id string) (*domain.Investor, error)
	FindByWallet(ctx context.Context, wallet string) (*domain.Investor, error)
	FindByJurisdiction(ctx context.Context, jurisdiction string) ([]*domain.Investor, error)
	Update(ctx context.Context, inv *domain.Investor) error
	Delete(ctx context.Context, id string) error
}
