package entities

func ValidateCPF(cpf string) bool {
	return len(cpf) == 11
}
