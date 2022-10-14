package event

import (
	"context"

	db "github.com/cryptnode-software/hermes/gorm"
	"github.com/google/uuid"
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

func (service *Service) Save(ctx context.Context, message *db.Event) (*db.Event, error) {
	if message.ID == uuid.Nil {
		if err := service.db.Save(message).Error; err != nil {
			return nil, err
		}
		return message, nil
	}

	service.db.Model(new(db.Event)).
		Update("metadata", message.Metadata).
		Update("text", message.Text).
		Where("id = ?", message.ID)
	return message, nil
}

func (service *Service) Delete(ctx context.Context, message *db.Event) error {
	return service.db.Delete(message).Error
}
