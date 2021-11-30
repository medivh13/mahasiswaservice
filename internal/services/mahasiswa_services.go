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

func (s *service) GetMahasiswaAlamatByID(req *dto.GetMahasiswaAlamatByIDReqDTO) (*dto.GetMahasiswaAlamatByIDRespDTO, error) {
	var resp *dto.GetMahasiswaAlamatByIDRespDTO

	getMahasiswaMap := make(map[int64]*dto.GetMahasiswaAlamatByIDRespDTO)
	data, err := s.repo.GetMahasiswaAlamatByID(req.ID)

	if err != nil {
		return nil, err
	}

	for _, val := range data {
		if _, ok := getMahasiswaMap[val.ID]; !ok {
			getMahasiswaMap[val.ID] = &dto.GetMahasiswaAlamatByIDRespDTO{
				ID:   val.ID,
				Nama: val.Name,
				Nim:  val.Nim,
			}
			getMahasiswaMap[val.ID].Alamats = append(getMahasiswaMap[val.ID].Alamats, &dto.AlamatRespDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		} else {
			getMahasiswaMap[val.ID].Alamats = append(getMahasiswaMap[val.ID].Alamats, &dto.AlamatRespDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		}
	}

	for _, val := range getMahasiswaMap {
		resp = val
	}

	return resp, nil
}
