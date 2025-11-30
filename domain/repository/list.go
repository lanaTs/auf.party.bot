package repository

import "auf.party.bot/domain/models"

type List interface {
	Create(l models.List) (models.List, error)
	Update(l models.List) (models.List, error)
	Delete(id int64) error

	GetByID(id int64) (*models.List, error)
	GetByChatID(id int64) ([]*models.List, error)
	// GetList(chatId int64) ([]*models.List, error)
}
