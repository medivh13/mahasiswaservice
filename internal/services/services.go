package services

import "github.com/medivh13/mahasiswaservice/pkg/dto"

type Services interface {
	SaveMahasiswaAlamat(req *dto.MahasiswaReqDTO) error
	UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error
	GetMahasiswaAlamatByID(req *dto.GetMahasiswaAlamatByIDReqDTO) (*dto.GetMahasiswaAlamatByIDRespDTO, error)
}
