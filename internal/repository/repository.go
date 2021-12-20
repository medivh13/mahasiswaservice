package repository

import "github.com/medivh13/mahasiswaservice/internal/models"

type Repository interface {
	SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatModels) error
	UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error
	GetMahasiswaAlamatByID(id int64) ([]*models.GetMahasiswaAlamatsModels, error)
	GetMahasiswaAlamat(where string) ([]*models.GetMahasiswaAlamatsModels, error)
}
