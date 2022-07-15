package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrMandatoryName   = errors.New("a name is required to create the account")
	ErrMandatoryCPF    = errors.New("CPF is required to create the account")
	ErrInvalidCPF      = errors.New("CPF is invalid")
	ErrNegativeBalance = errors.New("the account balance cannot be negative")
	ErrNotFound        = errors.New("not found")
	ErrMandatoryPIN    = errors.New("PIN is required to create the account")
)

type Account struct {
	ID        string
	Name      string
	CPF       string
	PIN       string
	Balance   int
	CreatedAt time.Time
}

func NewAccount(name, cpf, pin string, balance int) (Account, error) {
	if name == "" {
		return Account{}, ErrMandatoryName
	}

	if cpf == "" {
		return Account{}, ErrMandatoryCPF
	}

	ValidateCPF := ValidateCPF(cpf)
	if !ValidateCPF {
		return Account{}, ErrInvalidCPF
	}

	if pin == "" {
		return Account{}, ErrMandatoryPIN
	}

	hash, err := HashPIN(pin)
	if err != nil {
		return Account{}, fmt.Errorf("could not generate hash: %w", err)
	}

	if balance < 0 {
		return Account{}, ErrNegativeBalance
	}

	return Account{
		ID:        uuid.New().String(),
		Name:      name,
		CPF:       cpf,
		Balance:   balance,
		PIN:       hash,
		CreatedAt: time.Now(),
	}, nil
}

func HashPIN(pin string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pin), 7)
	return string(bytes), err
}

func CheckPINHash(pin, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pin)) != nil
}
