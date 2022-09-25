package user

import (
	"github.com/cryptnode-software/hermes/pkg"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService() (result *Service, err error) {
	result = new(Service)
	if result.db, err = pkg.GetDatabase(); err != nil {
		return
	}
	return
}
