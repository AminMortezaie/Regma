package domain

import "time"

type AssetType string

const (
	AssetTypeRealEstate AssetType = "real_estate"
	AssetTypeEquity     AssetType = "equity"
	AssetTypeDebt       AssetType = "debt"
	AssetTypeCommodity  AssetType = "commodity"
)

type AssetStatus string

const (
	AssetStatusDraft   AssetStatus = "draft"
	AssetStatusActive  AssetStatus = "active"
	AssetStatusPaused  AssetStatus = "paused"
	AssetStatusRetired AssetStatus = "retired"
)

type Asset struct {
	ID                   string
	Name                 string
	Symbol               string
	AssetType            AssetType
	Status               AssetStatus
	TotalSupply          uint64
	Decimals             uint8
	IssuerID             string
	AllowedJurisdictions []string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
