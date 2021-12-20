package repository

import (
	"fmt"
	"log"

	"github.com/medivh13/mahasiswaservice/internal/models"
	"github.com/medivh13/mahasiswaservice/internal/repository"

	"github.com/jmoiron/sqlx"
	mhsErrors "github.com/medivh13/mahasiswaservice/pkg/errors"
)

const (
	SaveMahasiswa          = `INSERT INTO kampus.mahasiswas (nama, nim, created_at) VALUES ($1, $2, now()) RETURNING id`
	SaveMahasiswaAlamat    = `INSERT INTO kampus.mahasiswa_alamats (jalan, no_rumah, created_at, id_mahasiswas) VALUES ($1,$2, now(), $3)`
	UpdateMahasiswaNama    = `UPDATE kampus.mahasiswas SET nama = $1, updated_at = now() where id = $2`
	TampilMahasiswa        = `SELECT id, nama, nim from kampus.mahasiswas WHERE %s`
	GetMahasiswaAlamatByID = `SELECT a.id, a.nama, a.nim, b.jalan, b.no_rumah from kampus.mahasiswas a JOIN kampus.mahasiswa_alamats b ON a.id = b.id_mahasiswas
	WHERE a.id = $1`

	GetMahasiswaAlamatAllOrByParam = `SELECT a.id, a.nama, a.nim, b.jalan, b.no_rumah from kampus.mahasiswas a JOIN kampus.mahasiswa_alamats b ON a.id = b.id_mahasiswas
	WHERE %s`
)

var statement PreparedStatement

type PreparedStatement struct {
	updateMahasiswaNama    *sqlx.Stmt
	getMahasiswaAlamatByID *sqlx.Stmt
}

type PostgreSQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &PostgreSQLRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

func (p *PostgreSQLRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *PostgreSQLRepo) {
	statement = PreparedStatement{
		updateMahasiswaNama:    m.Preparex(UpdateMahasiswaNama),
		getMahasiswaAlamatByID: m.Preparex(GetMahasiswaAlamatByID),
	}
}

func (p *PostgreSQLRepo) SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatModels) error {

	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx SaveMahasiswa Alamat : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}
	var idMahasiswa int64
	err = tx.QueryRow(SaveMahasiswa, dataMahasiswa.Name, dataMahasiswa.Nim).Scan(&idMahasiswa)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query SaveMahasiswa: ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	for _, val := range dataAlamat {
		_, err = tx.Exec(SaveMahasiswaAlamat, val.Jalan, val.NoRumah, idMahasiswa)
		if err != nil {
			tx.Rollback()
			log.Println("Failed Query SaveMahasiswaAlamat : ", err.Error())
			return fmt.Errorf(mhsErrors.ErrorDB)
		}
	}

	return tx.Commit()
}

func (p *PostgreSQLRepo) UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error {

	result, err := statement.updateMahasiswaNama.Exec(dataMahasiswa.Name, dataMahasiswa.ID)

	if err != nil {
		log.Println("Failed Query UpdateMahasiswaNama : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Failed RowAffectd UpdateMahasiswaNama : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("UpdateMahasiswaNama: No Data Changed")
		return fmt.Errorf(mhsErrors.ErrorNoDataChange)
	}

	return nil
}

func (p *PostgreSQLRepo) GetMahasiswaAlamatByID(id int64) ([]*models.GetMahasiswaAlamatsModels, error) {
	var data []*models.GetMahasiswaAlamatsModels

	err := statement.getMahasiswaAlamatByID.Select(&data, id)
	if err != nil {
		log.Println("Failed Query GetMahasiswaAlamatByID: ", err.Error())
		return data, fmt.Errorf(mhsErrors.ErrorDB)
	}
	if len(data) == 0 {
		return data, fmt.Errorf(mhsErrors.ErrorDataNotFound)
	}
	return data, nil
}

func (p *PostgreSQLRepo) GetMahasiswaAlamat(where string) ([]*models.GetMahasiswaAlamatsModels, error) {
	var data []*models.GetMahasiswaAlamatsModels
	var query string

	if where != "" && where != "%s" {
		query = fmt.Sprintf(GetMahasiswaAlamatAllOrByParam, where)
	} else {
		query = fmt.Sprintf(GetMahasiswaAlamatAllOrByParam, "1=1")
	}

	err := p.Conn.Select(&data, query)
	if err != nil {
		log.Println("Failed Query GetMahasiswaAlamat: ", err.Error())
		return data, fmt.Errorf(mhsErrors.ErrorDB)
	}

	if len(data) == 0 {
		return data, fmt.Errorf(mhsErrors.ErrorDataNotFound)
	}
	return data, nil
}
