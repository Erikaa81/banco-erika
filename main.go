package main

import (
	"fmt"

	"github.com/Erikaa81/banco-erika/domain/entities"
)

func main() {

	account, err := entities.NewAccount("João", "33344455566", "Efr2", 45.0)
	if err != nil {
		fmt.Printf("Erro ao criar nova conta: %v", err)
	} else {
		fmt.Println("Conta: ", account.ID)
		fmt.Println("Criada em : ", account.CreatedAt)
		fmt.Println("Saldo: ", account.Balance)

	}
}
