package rest_api_gin

import "errors"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type NoteList struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title" binding:"required" `
	Content string `json:"content" db:"content"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}


type UpdateListInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Content == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
