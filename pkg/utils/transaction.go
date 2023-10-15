package utils

import (
	"time"

	"github.com/Arjun-P17/tax-go/models"
)

type Transaction interface {
	GetDate() time.Time
	GetBasis() float64
}

func IsUniqueTransaction[T Transaction](transactions []T, transaction models.Transaction) bool {
	for _, t := range transactions {
		if t.GetDate() == transaction.Date && t.GetBasis() == transaction.Basis {
			return false
		}
	}
	return true
}
