package controllers

import (
	"html/template"
	"log"
	"net/http"
	"projetos-web/models"
	"strconv"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosLivros := models.BuscaTodosOsLIvros()
	templ.ExecuteTemplate(w, "Index", todosLivros)
}

func New(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		titulo := r.FormValue("titulo")
		autor := r.FormValue("autor")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		models.CriaNovoLivro(titulo, autor, precoConvertidoParaFloat, quantidadeConvertidaParaInt)

	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaLivro(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idLivro := r.URL.Query().Get("id")
	livro := models.EditaLivro(idLivro)
	templ.ExecuteTemplate(w, "Edit", livro)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		titulo := r.FormValue("titulo")
		autor := r.FormValue("autor")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na convesão da quantidade para int:", err)
		}

		models.AtualizaLivro(idConvertidaParaInt, titulo, autor, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
