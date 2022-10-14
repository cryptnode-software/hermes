package user

import (
	db "github.com/cryptnode-software/hermes/gorm"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService() (result *Service, err error) {
	result = new(Service)
	if result.db, err = db.Get(); err != nil {
		return
	}
	return
}
