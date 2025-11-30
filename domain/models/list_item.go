package models

type ListItem struct {
	Id      int64
	ListId  string
	Done    bool
	Name    string
	OwnerId int64
}
