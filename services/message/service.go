package message

import (
	"github.com/cryptnode-software/hermes/model"
	"github.com/cryptnode-software/hermes/pkg"
	"github.com/google/uuid"
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

func (service *Service) Save(message *model.Event) (*model.Event, error) {
	if message.ID == uuid.Nil {
		if err := service.db.Save(message).Error; err != nil {
			return nil, err
		}
		return message, nil
	}

	service.db.Model(new(model.Event)).
		Update("metadata", message.Metadata).
		Update("text", message.Text).
		Where("id = ?", message.ID)
	return message, nil
}

func (service *Service) Delete(message *model.Event) error {
	return service.db.Delete(message).Error
}
