package models

type User struct {
	ID       int64     `gorm:"primaryKey" json:"id"`
	Nama     string    `gorm:"type:varchar(255)" json:"nama"`
	Pesanans []Pesanan `gorm:"foreignKey:UserID" json:"pesanans"`
}

type Pesanan struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	UserID    int64  `gorm:"index" json:"user_id"`
	Berat     int64  `gorm:"type:int" json:"berat"`
	Deskripsi string `gorm:"type:varchar(255)" json:"deskripsi"`
}
