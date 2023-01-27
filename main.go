package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	for {
		fmt.Println("\n1 - Cadastrar prato" + "\n2 - Listar pratos" + "\n3 - Atualizar prato" + "\n4 - Deletar prato" + "\n0 - Sair")
		var op int
		_, _ = fmt.Scan(&op)

		switch op {
		case 0:
			fmt.Println("Saindo...")
			return
		case 1:
			registerPrato()
		case 2:
			showPratos()
		case 3:
			updatePrato()
		case 4:
			deletePrato()
		}
	}
}

func getConnection() (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", "postgres://postgres:12345678@localhost:5432/restaurante_db?sslmode=disable")
	return conn, err
}
