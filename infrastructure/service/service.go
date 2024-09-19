package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MikelSot/trival-proyect/model"
)

type Service struct {
	config model.ConfigRouter
}

func New(config model.ConfigRouter) Service {
	return Service{config: config}
}

func (s Service) GetRandomUser() (model.Results, error) {
	client := http.Client{}

	resp, err := client.Get(s.config.RandomUserUrl)
	if err != nil {
		log.Println("Error: al hacer la petición al servicio de RandomUser:", err)

		return nil, err
	}
	defer resp.Body.Close()

	var response model.ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println("Error: decodificando la respuesta del servicio de RandomUser:", err)

		return nil, err
	}

	var results model.Results
	for _, result := range response.Results {
		results = append(results, result)
	}

	return results, nil
}
