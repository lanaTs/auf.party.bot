package models

import "errors"

type ListItem struct {
	Id        int64
	ListId    int64
	Done      bool
	Name      string
	CreatedBy int64
	Order     int
}

func (li *ListItem) Validate() error {
	if li.Id <= 0 {
		return errors.New("ListItem id should be positive")
	}

	if li.CreatedBy <= 0 {
		return errors.New("ListItem CreatedBy should be positive")
	}

	if li.Name == "" {
		return errors.New("ListItem Name should not be null")
	}

	return nil
}
