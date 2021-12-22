package integration

import "github.com/medivh13/mahasiswaservice/pkg/dto"

type IntegServices interface {
	GetRandomDadJokes(req *dto.GetDadJokesInternalReqDTO) (*dto.GetDadJokesRandomRespDTO, error)
}
