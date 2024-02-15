package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gabrielmoura/geocoding-go/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

type Cep struct {
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

func GetCep(ctx context.Context, cep string) (Cep, error) {
	// cria uma solicitação HTTP com o contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		logger.Logger.Error("erro ao criar a solicitação HTTP", zap.Error(err))
		return Cep{}, fmt.Errorf("erro ao criar a solicitação HTTP: %v", err)
	}

	// executa a solicitação HTTP
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Logger.Error("erro ao executar a solicitação HTTP", zap.Error(err))
		return Cep{}, fmt.Errorf("erro ao executar a solicitação HTTP: %v", err)
	}

	// verifica se a resposta tem um status de sucesso
	if resp.StatusCode != http.StatusOK {
		logger.Logger.Error("resposta com status de erro", zap.String("status", resp.Status), zap.Error(err))
		return Cep{}, fmt.Errorf("resposta com status de erro: %v", resp.Status)
	}

	bind := Cep{}
	err = json.NewDecoder(resp.Body).Decode(&bind)
	if err != nil {
		fmt.Println("Erro ao decodificar a resposta:", err)
		return Cep{}, err
	}
	return bind, nil
}
