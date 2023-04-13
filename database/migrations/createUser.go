package migrations

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id      uint      `json:"id" gorm:"primarykey" key:"autoIncrement"`
	Nama    *string   `json:"nama"`
	Email   *string   `json:"email"`
	Tanggal time.Time `json:"tanggal"`
	Kota    *string   `json:"kota"`
	Negara  *string   `json:"negara"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
