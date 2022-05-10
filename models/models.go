package models

import (
	"projetos-web/db"
)

func BuscaTodosOsLIvros() []Livros {

	db := db.ConectaComBancoDeDados()

	selectDeTodosOsLivros, err := db.Query("select * from livros")
	if err != nil {
		panic(err.Error())
	}

	p := Livros{}
	livros := []Livros{}

	for selectDeTodosOsLivros.Next() {

		err = selectDeTodosOsLivros.Scan(&p.Id, &p.Titulo, &p.Autor, &p.Preco, &p.Quantidade)
		if err != nil {
			panic(err.Error())
		}

		livros = append(livros, p)
	}
	defer db.Close()
	return livros
}

func CriaNovoLivro(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insertDadosNoBanco, err := db.Prepare(`insert into livros(titulo, autor, preco, quantidade) 
												  values ($1, $2, $3,$4)`)

	if err != nil {
		panic(err.Error())
	}

	insertDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	db.Close()
}

func DeletaLivro(id string) {
	db := db.ConectaComBancoDeDados()

	deletarLivro, err := db.Prepare("delete from livros where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarLivro.Exec(id)
	defer db.Close()
}

func EditaLivro(id string) Livros {
	db := db.ConectaComBancoDeDados()

	livroDoBanco, err := db.Query("select * from livros where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	livroParaAtualizar := Livros{}

	for livroDoBanco.Next() {

		err = livroDoBanco.Scan(&livroParaAtualizar.Id,
			&livroParaAtualizar.Titulo,
			&livroParaAtualizar.Autor,
			&livroParaAtualizar.Preco,
			&livroParaAtualizar.Quantidade)
		if err != nil {
			panic(err.Error())
		}
	}
	defer db.Close()
	return livroParaAtualizar
}

func AtualizaLivro(id int, titulo, autor string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	AtualizarLivro, err := db.Prepare("update livros set titulo=$1, autor=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizarLivro.Exec(titulo, autor, preco, quantidade, id)
	defer db.Close()
}
