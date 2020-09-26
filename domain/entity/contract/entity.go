package contract

import "fypm.com/domain/entity"

type Contract struct {
	ID entity.ID
	PremiumInCents int
	CourtageInPercent float64
	Insurer string
}
