package entities

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrMandatorySecret = errors.New("password is required for login")
	ErrCPF             = errors.New("CPF is required for login")
)

type Login struct {
	CPF    string
	Secret string
}

func NewLogin(cpf, secret string) (Login, error) {
	if cpf == "" {
		return Login{}, ErrCPF
	}

	ValidateCPF := ValidateCPF(cpf)
	if !ValidateCPF {
		return Login{}, ErrInvalidCPF
	}

	if secret == "" {
		return Login{}, ErrMandatorySecret
	}

	hash, err := HashSecret(secret)
	if err != nil {
		return Login{}, fmt.Errorf("could not generate hash: %w", err)
	}

	return Login{
		CPF:    cpf,
		Secret: hash,
	}, nil
}

func HashSecret(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), 7)
	return string(bytes), err
}

func CheckSecretHash(secret, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret)) != nil
}
