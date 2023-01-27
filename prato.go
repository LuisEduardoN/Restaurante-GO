package main

import (
	"fmt"
	"log"
)

type prato struct {
	codigo       int     `db:"codigo"`
	nome         string  `db:"nome"`
	ingredientes string  `db:"ingredientes"`
	tipo         string  `db:"tipo"`
	preco        float64 `db:"preco"`
}

func registerPrato() {
	myPrato := prato{}

	fmt.Println("Digite o nome do prato: ")
	_, _ = fmt.Scan(&myPrato.nome)

	if myPrato.nome == "" {
		fmt.Println("Nome inválido")
		return
	}

	fmt.Println("Digite os ingredientes do prato: ")
	_, _ = fmt.Scan(&myPrato.ingredientes)

	if myPrato.ingredientes == "" {
		fmt.Println("Ingredientes inválidos")
		return
	}

	fmt.Println("Digite o tipo do prato: ")
	_, _ = fmt.Scan(&myPrato.tipo)

	if myPrato.tipo == "" {
		fmt.Println("Tipo inválido")
		return
	}

	fmt.Println("Digite o preco do preço: ")
	_, _ = fmt.Scan(&myPrato.preco)

	if myPrato.preco <= 0 {
		fmt.Println("Preço inválido")
		return
	}

	conn, err := getConnection()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	insert := `
		insert into prato (nome,ingredientes,tipo,preco)
		values ($1, $2, $3, $4)`

	_, err = conn.Exec(insert, myPrato.nome, myPrato.ingredientes, myPrato.tipo, myPrato.preco)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Prato cadastrado!")
	}

}

func showPratos() {
	pratos := buscarPratos()
	for _, prato := range pratos {
		fmt.Printf("\nCodigo: %d \nNome: %s \nIngredientes: %s \nTipo: %s \nPreço: %.2f R$\n", prato.codigo, prato.nome, prato.ingredientes, prato.tipo, prato.preco)
		fmt.Println("---------------------")
	}
}

func buscarPratos() []prato {
	conn, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rows, err := conn.Query("select codigo,nome,ingredientes,tipo,preco from prato")
	if err != nil {
		log.Fatal(err)
	}

	var pratos []prato

	for rows.Next() {
		p := prato{}
		err := rows.Scan(&p.codigo, &p.nome, &p.ingredientes, &p.tipo, &p.preco)
		if err != nil {
			log.Fatal(err)
		}

		pratos = append(pratos, p)
	}
	return pratos
}

func deletePrato() {
	var cod int
	fmt.Println("Digite o codigo do prato: ")
	fmt.Scan(&cod)

	conn, err := getConnection()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	query := "delete from prato where codigo = $1"
	r, err := conn.Exec(query, cod)
	if err != nil {
		panic(err)
	}

	linhasAfetadas, _ := r.RowsAffected()
	if linhasAfetadas > 0 {
		fmt.Println("\nOs dados foram deletados")
	} else {
		fmt.Println("\nNada foi alterado")
	}
}

func updatePrato() {
	var cod int
	fmt.Println("Digite o codigo do prato: ")
	fmt.Scan(&cod)

	conn, err := getConnection()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	var newNome string
	var newIngredientes string
	var newTipo string
	var newPreco float64
	fmt.Println("Digite um novo nome: ")
	fmt.Scan(&newNome)
	fmt.Println("Digite os novos ingredientes: ")
	fmt.Scan(&newIngredientes)
	fmt.Println("Digite um novo tipo: ")
	fmt.Scan(&newTipo)
	fmt.Println("Digite um novo preço: ")
	fmt.Scan(&newPreco)
	query := "update prato set nome = $2, ingredientes = $3, tipo = $4, preco = $5 where codigo = $1"
	_, erro := conn.Exec(query, cod, newNome, newIngredientes, newTipo, newPreco)
	if err != nil {
		log.Fatal(erro)
	}

}
