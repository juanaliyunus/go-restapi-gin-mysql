package models

type Product struct {
	Id          int64  `gorm:"primary_key" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
}
