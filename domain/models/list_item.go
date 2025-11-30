package models

import "errors"

type ListItem struct {
	Id      int64
	ListId  int64
	Done    bool
	Name    string
	OwnerId int64
}

func (li *ListItem) Validate() error {
	if li.Id <= 0 {
		return errors.New("ListItem id should be positive")
	}

	if li.OwnerId <= 0 {
		return errors.New("ListItem OwnerId should be positive")
	}

	if li.Name == "" {
		return errors.New("ListItem Name should not be null")
	}

	return nil
}
