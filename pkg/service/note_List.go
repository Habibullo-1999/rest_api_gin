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

func(s *NoteListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId,listId)
}

func(s *NoteListService) Update(userId, listId int, input rest_api_gin.UpdateListInput) error{
	if err := input.Validate(); err !=nil {
		return err
	} 
	return s.repo.Update(userId,listId,input)
}
