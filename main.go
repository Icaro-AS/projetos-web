package main

import (
	"net/http"
	"projetos-web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
