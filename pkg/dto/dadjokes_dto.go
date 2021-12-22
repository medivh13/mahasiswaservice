package dto

import (
	"errors"

	"github.com/medivh13/mahasiswaservice/pkg/common/crypto"
	"github.com/medivh13/mahasiswaservice/pkg/common/validator"
	"github.com/medivh13/mahasiswaservice/pkg/env"
	util "github.com/medivh13/mahasiswaservice/pkg/utils"
)

type GetDadJokesInternalReqDTO struct {
	Authorization string `json:"Authorization" valid:"required" validname:"datetime"`
	Signature     string `json:"signature" valid:"required" validname:"signature"`
	DateTime      string `json:"datetime" valid:"required" validname:"datetime"`
	ID            string `json:"id,omitempty,string"`
}

func (dto *GetDadJokesInternalReqDTO) Validate() error {
	v := validator.NewValidate(dto)
	v.SetCustomValidation(true, func() error {
		return dto.customValidation()
	})
	return v.Validate()
}

func (dto *GetDadJokesInternalReqDTO) customValidation() error {

	signature := crypto.EncodeSHA256HMAC(util.GetBTBPrivKeySignature(), dto.Authorization, dto.DateTime)
	if signature != dto.Signature {
		if env.IsProduction() {
			return errors.New("invalid signature")
		}
		return errors.New("invalid signature" + " --> " + signature)
	}

	return nil
}

type GetDadJokesRandomRespDTO struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int64  `json:"status"`
}
