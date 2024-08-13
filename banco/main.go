package main

import (
	"fmt"

	"github.com\caahpaiva\Golang\banco\contas"
    "github.com\caahpaiva\Golang\banco\clientes"

)

func main() {
		contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
			Nome: "Bruno", 
			CPF: "123.111.123.12",
			Profissao: "Desenvolvedor"},
			NumeroAgencia: 123, NumeroConta:123456, Saldo:100
		}

		fmt.Println(contaDoBruno)

	}




	// contaDaSilvia := ContaCorrente{}
	// contaDaSilvia.titular = "Silvia"
	// contaDaSilvia.saldo = 500

	// fmt.Println(contaDaSilvia.saldo)
	// // contaDaSilvia.Depositar(500)
	// status, valor := contaDaSilvia.Depositar(2000)
	// fmt.Println(status, valor)

	// fmt.Println(contaDaSilvia.Sacar(-100))
	// fmt.Println(contaDaSilvia.saldo)
	// contaDaSilvia := contas.ContaCorrenteContaCorrente{Titular: "Silvia", Saldo: 300}
	// contaDoGustavo := contas.ContaCorrenteContaCorrente{Titular: "Gustavo", Saldo: 100}

	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)

	// fmt.Println(status)
	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)
}
