package compliance

import (
	"context"

	"Regma/internal/domain"
)

type Repository interface {
	CreateRule(ctx context.Context, rule *domain.ComplianceRule) error
	FindRulesByAsset(ctx context.Context, assetID string) ([]*domain.ComplianceRule, error)
	UpdateRule(ctx context.Context, rule *domain.ComplianceRule) error
	DeleteRule(ctx context.Context, id string) error
}
