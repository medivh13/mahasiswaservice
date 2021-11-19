package services

import (
	"github.com/medivh13/mahasiswaservice/internal/repository"
	"github.com/medivh13/mahasiswaservice/pkg/dto"
	"github.com/medivh13/mahasiswaservice/pkg/dto/assembler"
)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Services {
	return &service{repo}
}

func (s *service) SaveMahasiswaAlamat(req *dto.MahasiswaReqDTO) error {

	dtAlamat := assembler.ToSaveMahasiswaAlamats(req.Alamats)
	dtMahasiswa := assembler.ToSaveMahasiswa(req)

	err := s.repo.SaveMahasiswaAlamat(dtMahasiswa, dtAlamat)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error {

	dtMhsiswa := assembler.ToUpdateMahasiswaNama(req)

	err := s.repo.UpdateMahasiswaNama(dtMhsiswa)
	if err != nil {
		return err
	}

	return nil
}
