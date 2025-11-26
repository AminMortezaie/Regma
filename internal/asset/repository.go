package asset

import (
	"context"

	"Regma/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, asset *domain.Asset) error
	FindByID(ctx context.Context, id string) (*domain.Asset, error)
	FindBySymbol(ctx context.Context, symbol string) (*domain.Asset, error)
	FindByIssuer(ctx context.Context, issuerID string) ([]*domain.Asset, error)
	Update(ctx context.Context, asset *domain.Asset) error
}
