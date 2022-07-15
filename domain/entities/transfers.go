package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidAmount               = errors.New("a valid amount is required")
	ErrSameAccounts                = errors.New("the target account cannot be the same as the source account")
	ErrMandatoryAccountOrigin      = errors.New("origin account is required to create transfer")
	ErrMandatoryAccountDestination = errors.New("destination account is required to create transfer")
)

type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
}

func NewTransfer(accountOriginID, accountDestinationID string, amount int) (Transfer, error) {
	if accountOriginID == "" {
		return Transfer{}, ErrMandatoryAccountOrigin
	}

	if accountDestinationID == "" {
		return Transfer{}, ErrMandatoryAccountDestination
	}

	if amount <= 0 {
		return Transfer{}, ErrInvalidAmount
	}

	if accountOriginID == accountDestinationID {
		return Transfer{}, ErrSameAccounts
	}

	return Transfer{
		ID:                   uuid.New().String(),
		AccountOriginID:      accountOriginID,
		AccountDestinationID: accountDestinationID,
		Amount:               amount,
		CreatedAt:            time.Now(),
	}, nil
}
