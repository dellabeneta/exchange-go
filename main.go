package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Struct atualizada para capturar mais dados da API
type Cotacoes map[string]struct {
	Nome     string `json:"name"`
	Compra   string `json:"bid"`
	Venda    string `json:"ask"`
	Variacao string `json:"pctChange"`
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	moedas := "USD-BRL,EUR-BRL,GBP-BRL,AUD-BRL,CAD-BRL"
	url := "https://economia.awesomeapi.com.br/json/last/" + moedas

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Falha ao buscar cotações", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Falha ao ler resposta", http.StatusInternalServerError)
		return
	}

	var data Cotacoes
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Falha ao decodificar JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/cotacao", cotacaoHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}