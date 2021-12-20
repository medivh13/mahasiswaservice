package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/medivh13/mahasiswaservice/pkg/dto"
	util "github.com/medivh13/mahasiswaservice/pkg/utils"
)

type integService struct {
}

func NewService() IntegServices {
	return &integService{}
}

func (s *integService) GetRandomDadJokes() (*dto.GetDadJokesRandomRespDTO, error) {
	var response dto.GetDadJokesRandomRespDTO

	url := util.GetIntegURL("icanhazdadjoke", "random")

	getReq, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create http request icanhazdadjoke getRandom: %v", err)
	}

	getReq.Header["Accept"] = []string{" application/json"}

	client := http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(getReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request icanhazdadjoke getRandom: %v", err)
	}
	log.Println("Success execute  : ", url)
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode icanhazdadjoke getRandom response: %v", err)
	}

	return &response, nil
}
