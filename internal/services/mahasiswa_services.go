package services

import (
	"fmt"

	integ "github.com/medivh13/mahasiswaservice/internal/integration"
	"github.com/medivh13/mahasiswaservice/internal/repository"
	"github.com/medivh13/mahasiswaservice/pkg/dto"
	"github.com/medivh13/mahasiswaservice/pkg/dto/assembler"
)

type service struct {
	repo      repository.Repository
	IntegServ integ.IntegServices
}

func NewService(repo repository.Repository, IntegServ integ.IntegServices) Services {
	return &service{repo, IntegServ}
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

func (s *service) GetMahasiswaAlamat(req *dto.GetMahasiswaAlamatReqDTO) ([]*dto.GetMahasiswaAlamatRespDTO, error) {
	resp := []*dto.GetMahasiswaAlamatRespDTO{}
	query := "%s"
	where := `'%s'`
	if req.Nama != "" {
		where = fmt.Sprintf(where, req.Nama)
		query = fmt.Sprintf(query, `a.nama = `+where)
	}

	getMahasiswaMap := make(map[int64]*dto.GetMahasiswaAlamatRespDTO)

	data, err := s.repo.GetMahasiswaAlamat(query)

	if err != nil {
		return nil, err
	}

	for _, val := range data {
		if _, ok := getMahasiswaMap[val.ID]; !ok {
			getMahasiswaMap[val.ID] = &dto.GetMahasiswaAlamatRespDTO{
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
		resp = append(resp, val)
	}

	return resp, nil

}

func (s *service) GetIntegDadJoke(req *dto.GetDadJokesInternalReqDTO) (*dto.GetDadJokesRandomRespDTO, error) {
	var resp *dto.GetDadJokesRandomRespDTO

	resp, err := s.IntegServ.GetRandomDadJokes()

	if err != nil {
		return nil, err
	}

	return resp, nil

}
