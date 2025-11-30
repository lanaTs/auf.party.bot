package repository

import "auf.party.bot/domain/models"

type ListItem interface {
	Create(li models.ListItem) (models.ListItem, error)
	Update(li models.ListItem) (models.ListItem, error)
	Delete(id int64) error

	GetByID(id int64) (*models.ListItem, error)
	GetByListID(id int64) ([]*models.ListItem, error)
	ToggleCheck(id int64) error
	// GetList(chatId int64) ([]*models.ListItem, error)
}
