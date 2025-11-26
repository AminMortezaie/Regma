package domain

import "time"

type InvestorType string

const (
	InvestorTypeRetail        InvestorType = "retail"
	InvestorTypeAccredited    InvestorType = "accredited"
	InvestorTypeInstitutional InvestorType = "institutional"
)

type KYCStatus string

const (
	KYCStatusPending  KYCStatus = "pending"
	KYCStatusVerified KYCStatus = "verified"
	KYCStatusRejected KYCStatus = "rejected"
	KYCStatusExpired  KYCStatus = "expired"
)

type Investor struct {
	ID            string
	WalletAddress string
	Jurisdiction  string // ISO 3166-1 alpha-2
	InvestorType  InvestorType
	KYCStatus     KYCStatus
	KYCProvider   string
	KYCExpiry     *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
