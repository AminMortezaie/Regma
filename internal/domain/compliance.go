package domain

import "time"

type RuleType string

const (
	RuleTypeJurisdiction  RuleType = "jurisdiction"
	RuleTypeInvestorType  RuleType = "investor_type"
	RuleTypeTransferLimit RuleType = "transfer_limit"
	RuleTypeHoldingPeriod RuleType = "holding_period"
	RuleTypeInvestorCount RuleType = "investor_count"
)

type ComplianceRule struct {
	ID        string
	AssetID   string
	RuleType  RuleType
	Config    map[string]any // Flexible rule configuration
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransferCheck struct {
	FromInvestorID string
	ToInvestorID   string
	AssetID        string
	Amount         uint64
}

type TransferResult struct {
	Allowed bool
	Reason  string
}
