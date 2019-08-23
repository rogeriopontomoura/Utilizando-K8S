package main

// Realiza a importação
import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("O servidor está escutando na porta 8000")
	http.HandleFunc("/", handler) // Declaração do path e Handler
	http.ListenAndServe(":8000", nil) // Porta que vai escutar

}

// Trata a requisição

func  handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, greeting("Code.education Rocks!!") )
}


// Função greeting

func greeting(text string) string {

	msg := "<b>" + text + "</b>"

	return msg
}