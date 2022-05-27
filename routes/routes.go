package routes

import (
	"net/http"

	"github.com/devopscodeck/controlles"
)

func CarregaRotas() {

	http.HandleFunc("/", controlles.Index)  //rota definida
	http.HandleFunc("/new", controlles.New) //rota definida
	http.HandleFunc("/insert", controlles.Insert)
	http.HandleFunc("/delete", controlles.Delete)
	http.HandleFunc("/edit", controlles.Edit)
	http.HandleFunc("/update", controlles.Update)
}
