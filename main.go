package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, forneça um CEP como argumento.")
		fmt.Println("Exemplo: go run main.go 01153000")
		os.Exit(1)
	}

	cep := os.Args[1]
	if len(cep) != 8 {
		fmt.Println("O CEP deve conter 8 dígitos.")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	chViaCEP := make(chan ViaCEP)
	chBrasilAPI := make(chan BrasilAPI)
	chErro := make(chan error)

	// Goroutine para ViaCEP
	go func() {
		url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			chErro <- err
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			chErro <- err
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			chErro <- err
			return
		}

		var cep ViaCEP
		if err := json.Unmarshal(body, &cep); err != nil {
			chErro <- err
			return
		}

		chViaCEP <- cep
	}()

	// Goroutine para BrasilAPI
	go func() {
		url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			chErro <- err
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			chErro <- err
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			chErro <- err
			return
		}

		var cep BrasilAPI
		if err := json.Unmarshal(body, &cep); err != nil {
			chErro <- err
			return
		}

		chBrasilAPI <- cep
	}()

	select {
	case viaCEP := <-chViaCEP:
		fmt.Println("Resposta mais rápida da API ViaCEP:")
		fmt.Printf("CEP: %s\n", viaCEP.Cep)
		fmt.Printf("Logradouro: %s\n", viaCEP.Logradouro)
		fmt.Printf("Complemento: %s\n", viaCEP.Complemento)
		fmt.Printf("Bairro: %s\n", viaCEP.Bairro)
		fmt.Printf("Cidade: %s\n", viaCEP.Localidade)
		fmt.Printf("UF: %s\n", viaCEP.Uf)

	case brasilAPI := <-chBrasilAPI:
		fmt.Println("Resposta mais rápida da API BrasilAPI:")
		fmt.Printf("CEP: %s\n", brasilAPI.Cep)
		fmt.Printf("Estado: %s\n", brasilAPI.State)
		fmt.Printf("Cidade: %s\n", brasilAPI.City)
		fmt.Printf("Bairro: %s\n", brasilAPI.Neighborhood)
		fmt.Printf("Rua: %s\n", brasilAPI.Street)

	case err := <-chErro:
		fmt.Printf("Erro: %v\n", err)

	case <-ctx.Done():
		fmt.Println("Timeout: nenhuma API respondeu em tempo hábil")
	}
}
