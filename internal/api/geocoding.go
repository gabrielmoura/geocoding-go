package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gabrielmoura/geocoding-go/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

type Geocoding struct {
	Lat         string      `json:"lat"`
	Lon         string      `json:"lon"`
	DisplayName string      `json:"display_name"`
	Type        string      `json:"type"`
	Addresstype string      `json:"addresstype"`
	Geojson     interface{} `json:"geojson"`
}

func GetGeocoding(ctx context.Context, address string) (Geocoding, error) {
	// cria uma solicitação HTTP com o contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "https://nominatim.openstreetmap.org/search.php?q="+address+"&polygon_geojson=1&format=jsonv2", nil)
	if err != nil {
		logger.Logger.Error("erro ao criar a solicitação HTTP", zap.Error(err))
		return Geocoding{}, fmt.Errorf("erro ao criar a solicitação HTTP: %v", err)
	}

	// executa a solicitação HTTP
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Logger.Error("erro ao executar a solicitação HTTP", zap.Error(err))
		return Geocoding{}, fmt.Errorf("erro ao executar a solicitação HTTP: %v", err)
	}

	// verifica se a resposta tem um status de sucesso
	if resp.StatusCode != http.StatusOK {
		logger.Logger.Error("resposta com status de erro", zap.String("status", resp.Status), zap.Error(err))
		return Geocoding{}, fmt.Errorf("resposta com status de erro: %v", resp.Status)
	}

	bind := Geocoding{}
	err = json.NewDecoder(resp.Body).Decode(&bind)
	if err != nil {
		//logger.Logger.Error("erro ao decodificar a resposta", zap.Error(err))
		return Geocoding{}, err
	}
	return bind, nil
}
