package models

import "errors"

type List struct {
	Id      int64
	Name    string
	ChatId  int64
	OwnerId int64
}

func (l *List) Validate() error {
	if l.Id <= 0 {
		return errors.New("List id should be positive")
	}

	if l.ChatId <= 0 {
		return errors.New("List ChatId should be positive")
	}

	if l.OwnerId <= 0 {
		return errors.New("List OwnerId should be positive")
	}

	if l.Name == "" {
		return errors.New("List Name should not be null")
	}

	return nil
}
