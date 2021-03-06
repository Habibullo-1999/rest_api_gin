package service

import (
	"github.com/Habibullo-1999/rest_api_gin"
	"github.com/Habibullo-1999/rest_api_gin/pkg/repository"

)

type Authorization interface {
	CreateUser(user rest_api_gin.User) (int, error)
	CreateToToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type NoteList interface {
	Create(userId int, list rest_api_gin.NoteList) (int, error)
	GetAll(userId int, )([]rest_api_gin.NoteList, error)
	GetById(userId,listId int) (rest_api_gin.NoteList, error)	
	Delete(userId, listId int) error
	Update(userId, listId int, input rest_api_gin.UpdateListInput) error
}

type NoteItem interface {
}

type Service struct {
	Authorization
	NoteList
	NoteItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		NoteList: NewNoteListService(repos.NoteList),
		}
}
