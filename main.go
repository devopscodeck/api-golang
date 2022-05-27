package main

import (
	"net/http"

	"github.com/devopscodeck/routes"
)

//Funcao principal que inicia o servidor a partir do diretorio raiz localhost/
func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
