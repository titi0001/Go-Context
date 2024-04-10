package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou request no servidor")
	defer log.Println("Finalizou a Request no servidor")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Requisição processada com sucesso no servidor")
		w.Write([]byte("Requisição processa com sucesso no client"))
	case <-ctx.Done():
		http.Error(w, "Request cancelada no cliente", http.StatusRequestTimeout)
	}

}
