package repository

import (
	"github.com/medivh13/mahasiswaservice/internal/repository"

	"github.com/jmoiron/sqlx"
)

const (
	SaveInitSubmitRDL       = `INSERT INTO rdl_account (partnerRefNum, fullName, short_name, date_of_birth, place_of_birth, gender, religion, marital_status, mother_maiden_name, nationality, identSerialNum, npwp, phone_number, handphone, email, address, kelurahan, kecamatan, rt_dan_rw, city, province, zipcode, occupation, income_range) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	GetPartnerRefNumQuery   = `SELECT partnerRefNum from rdl_account where partnerRefNum = ? OR identSerialNum = ?`
	UpdateRDLStatusInfo     = `UPDATE rdl_account SET datetime = ?, emailAddr = ?, phoneNumber = ?, statusCode = ?, statusDesc = ?, custNo = ?, accountNo = ?, branchName = ?, expDateKYC = ?, expDateInitialDeposit = ?, applyDate = ?, acctOpenDate = ?, acctClosedDate = ?, remarksVBO = ?, VBOScheduleDate = ?, VBOScheduleTime = ?, intervalDaysEmail = ?, cardLinkFlag = ?, cardDeliveryFlag = ?, remarksDLCU = ?, remarksOther = ? WHERE partnerRefNum= ?`
	GetAverageTransaction   = `SELECT id, jhi_key, jhi_value, language, created_date FROM averageTransaction`
	GetBusinessSector       = `SELECT id, jhi_key, jhi_value, language, created_date FROM businessSector`
	GetBloodType            = `SELECT id, jhi_key, jhi_value, language, created_date FROM bloodType`
	GetEducation            = `SELECT id, jhi_key, jhi_value, language, created_date FROM education`
	GetFamilyRelation       = `SELECT id, jhi_key, jhi_value, language, created_date FROM familyRelation`
	GetIndustrialType       = `SELECT id, jhi_key, jhi_value, language, created_date FROM industrialType`
	GetJobPosition          = `SELECT id, jhi_key, jhi_value, language, created_date FROM jobPosition where language = "id"`
	GetMonthlySalary        = `SELECT id, jhi_key, jhi_value, language, created_date FROM monthLySalary`
	GetOccupation           = `SELECT id, jhi_key, jhi_value, language, created_date FROM occupation`
	GetPurposeOpenAccount   = `SELECT id, jhi_key, jhi_value, language, created_date FROM purposeOpenAccount`
	GetSourceOfFund         = `SELECT id, jhi_key, jhi_value, language, created_date FROM sourceOfFund`
	GetMaritalStatus        = `SELECT id, source_look_up, sode_source_look_up_desc, created_date FROM marital_status`
	GetReligion             = `SELECT id, source_look_up, sode_source_look_up_desc, created_date FROM religion`
	GetGender               = `SELECT id, source_look_up, sode_source_look_up_desc, created_date FROM gender`
	GetProvinces            = `SELECT id, name, province_code, created_date FROM provinces`
	GetRegenciesByProvinces = `SELECT id, name, provinces_id, created_date FROM regencies where provinces_id = ?`
	GetDistrictsByRegency   = `SELECT id, name, regencies_id, created_date FROM district where regencies_id = ?`
	GetVillagesByDistrict   = `SELECT id, name, districs_id, postal_code, created_date FROM villages where districs_id = ?`
	GetNPWPReasons          = `SELECT id, description, created_date FROM npwp_reason`
	GetPurposeFunds         = `SELECT id, jhi_group, jhi_key, jhi_value, language, created_date FROM purposefund`
	GetRDLACcountByEmail    = `SELECT partnerRefNum, dateTime, fullName, emailAddr, phoneNumber, statusCode, statusDesc, custNo, accountNo, branchName, expDateKYC, expDateInitialDeposit, applyDate, acctOpenDate, acctClosedDate, remarksVBO, VBOScheduleDate, VBOScheduleTime, intervalDaysEmail, cardLinkFlag, cardDeliveryFlag, remarksDLCU, remarksOther from rdl_account WHERE emailAddr = ?`
)

var statement PreparedStatement

type PreparedStatement struct {
	saveInitSubmit          *sqlx.Stmt
	existingPartnerRefNum   *sqlx.Stmt
	updateStatusInfo        *sqlx.Stmt
	getAverageTransaction   *sqlx.Stmt
	getBusinessSector       *sqlx.Stmt
	getBloodType            *sqlx.Stmt
	getEducation            *sqlx.Stmt
	getFamilyRelation       *sqlx.Stmt
	getIndustrialType       *sqlx.Stmt
	getJobPosition          *sqlx.Stmt
	getMonthlySalary        *sqlx.Stmt
	getOccupation           *sqlx.Stmt
	getPurposeOpenAccount   *sqlx.Stmt
	getSourceOfFund         *sqlx.Stmt
	getMaritalStatus        *sqlx.Stmt
	getReligion             *sqlx.Stmt
	getGender               *sqlx.Stmt
	getProvinces            *sqlx.Stmt
	getRegenciesByProvinces *sqlx.Stmt
	getDistrictsByRegency   *sqlx.Stmt
	getVillagesByDistric    *sqlx.Stmt
	getNPWPReasons          *sqlx.Stmt
	getPurposeFunds         *sqlx.Stmt
	getRDLAccountByEmail    *sqlx.Stmt
}

type MySQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &MySQLRepo{Conn}
	// InitPreparedStatement(repo)
	return repo
}
