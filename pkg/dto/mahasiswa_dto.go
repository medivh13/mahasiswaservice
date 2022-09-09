package dto

import (
	"errors"

	"github.com/medivh13/mahasiswaservice/pkg/common/crypto"
	"github.com/medivh13/mahasiswaservice/pkg/common/validator"
	"github.com/medivh13/mahasiswaservice/pkg/env"
	util "github.com/medivh13/mahasiswaservice/pkg/utils"
)

type MahasiswaReqDTO struct {
	Nama    string         `json:"nama" valid:"required" validname:"nama"`
	Nim     string         `json:"nim" valid:"required" validname:"nim"`
	Alamats []AlamatReqDTO `json:"alamat" valid:"required" `
}

type AlamatReqDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}

func (dto *MahasiswaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type UpadeMahasiswaNamaReqDTO struct {
	Nama string `json:"nama" valid:"required" validname:"nama"`
	ID   int64  `json:"id" valid:"required,integer,non_zero" validname:"id"`
}

func (dto *UpadeMahasiswaNamaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type GetMahasiswaAlamatByIDReqDTO struct {
	ID int64 `json:"id" valid:"required,integer,non_zero" validname:"id"`
}

func (dto *GetMahasiswaAlamatByIDReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type GetMahasiswaAlamatByIDRespDTO struct {
	ID      int64            `json:"id"`
	Nama    string           `json:"nama"`
	Nim     string           `json:"nim"`
	Alamats []*AlamatRespDTO `json:"alamat"`
}

type AlamatRespDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}

type GetMahasiswaAlamatReqDTO struct {
	Authorization string `json:"Authorization" valid:"required" validname:"datetime"`
	Signature     string `json:"signature" valid:"required" validname:"signature"`
	DateTime      string `json:"datetime" valid:"required" validname:"datetime"`
	Nama          string `json:"nama,omitempty,string"`
}

func (dto *GetMahasiswaAlamatReqDTO) Validate() error {
	v := validator.NewValidate(dto)
	v.SetCustomValidation(true, func() error {
		return dto.customValidation()
	})
	return v.Validate()
}

func (dto *GetMahasiswaAlamatReqDTO) customValidation() error {

	signature := crypto.EncodeSHA256HMAC(util.GetBTBPrivKeySignature(), dto.Authorization, dto.DateTime)
	if signature != dto.Signature {
		if env.IsProduction() {
			return errors.New("invalid signature")
		}
		return errors.New("invalid signature" + " --> " + signature)
	}

	return nil
}

type GetMahasiswaAlamatRespDTO struct {
	ID      int64            `json:"id"`
	Nama    string           `json:"nama"`
	Nim     string           `json:"nim"`
	Alamats []*AlamatRespDTO `json:"alamat"`
}

// vendor
// district
// PO
//  PO item
//  Qty Request
//  Qty receive
//  Price
//  UOM
// warehouse

// {
// 	vendor :Jody,
// 	district : Cilacap,
// 	PO :
//   		[
// 			Item :Baju,
//   			QtyReq : 10,
// 			QtyReceive : 10,
//   			Price : 10000,
//   			Uom : EA
// 	  	],
//   		[
// 			Item : "Baju",
//   			QtyReq : 10,
// 			QtyReceive : 10,
//   			Price : 10000,
//   			Uom : "EA"
//   		],
// 	warehouse : "Pucang"
// }
// a b c d e e f g h
// type GetVendorRespDTO struct {
// 	Vendor string `json: "vendor"`
// 	District string `json: "district"`
// 	PO []*PoDTO `json: "po"`
// 	Warehouse string `json:"warehouse"`
// }

// type PoDTO struct {
// 	Item string `json:"po_item"`
// 	QtyReq string `json:"qty_request"`
// 	QtyReceive string `json:"qty_receive"`
// 	Price string `json:"price"`
// 	UOM string `json : "uom"`
// }
