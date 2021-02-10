package repository

import (
	"github.com/Habibullo-1999/rest_api_gin"
	"github.com/jmoiron/sqlx"

)

type Authorization interface {
	CreateUser(user rest_api_gin.User) (int, error)
	GetUser(username, password string) (rest_api_gin.User, error)	
}

type NoteList interface {
	Create(userId int, list rest_api_gin.NoteList) (int, error)
	GetAll(userId int )([]rest_api_gin.NoteList, error)
	GetById(userId,listId int) (rest_api_gin.NoteList, error)
}

type NoteItem interface {
}

type Repository struct {
	Authorization
	NoteList
	NoteItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		NoteList: NewNoteListPostgres(db),
	}
}
