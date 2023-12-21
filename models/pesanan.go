package models

type Pesanan struct{
	Id int64 `gorm:"primaryKey" json:"id"`
	NamaPemesan string `gorm:"type:varchar(255)" json:"nama_pemesan"`
	Berat int64 `gorm:"type:int" json:"berat"`
}