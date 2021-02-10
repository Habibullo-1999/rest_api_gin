package service

import (
	"github.com/Habibullo-1999/rest_api_gin"
	"github.com/Habibullo-1999/rest_api_gin/pkg/repository"

)

type NoteListService struct {
	repo repository.NoteList
}

func NewNoteListService(repo repository.NoteList) *NoteListService {
	return &NoteListService{repo: repo}
}

func (s *NoteListService) Create(userId int, list rest_api_gin.NoteList) (int, error) {
	return s.repo.Create(userId,list)
}

func (s *NoteListService) GetAll(userId int, )([]rest_api_gin.NoteList, error) {
	return s.repo.GetAll(userId)
}

func (s *NoteListService) GetById(userId,listId int) (rest_api_gin.NoteList, error){
	return s.repo.GetById(userId,listId)
}