package repository

import (
	"fmt"
	"strings"

	"github.com/Habibullo-1999/rest_api_gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	// "github.com/pelletier/go-toml/query"

)

type NoteListPostgres struct {
	db *sqlx.DB
}

func NewNoteListPostgres(db *sqlx.DB) *NoteListPostgres {
	return &NoteListPostgres{db: db}
}

func (r *NoteListPostgres) Create(userId int, list rest_api_gin.NoteList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, content) VALUES ($1,$2) RETURNING id", noteListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Content)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1,$2)", userLIstTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *NoteListPostgres) GetAll(userId int) ([]rest_api_gin.NoteList, error) {
	var lists []rest_api_gin.NoteList

	query := fmt.Sprintf("SELECT nl.id, nl.title, nl.content FROM %s nl INNER JOIN %s ul on nl.id=ul.list_id WHERE ul.user_id = $1", noteListsTable, userLIstTable)

	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *NoteListPostgres) GetById(userId, listId int) (rest_api_gin.NoteList, error) {
	var list rest_api_gin.NoteList

	query := fmt.Sprintf("SELECT nl.id, nl.title, nl.content FROM %s nl INNER JOIN %s ul on nl.id=ul.list_id WHERE ul.user_id = $1 AND ul.list_id= $2", noteListsTable, userLIstTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *NoteListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s nl USING %s ul WHERE nl.id =ul.list_id AND ul.user_id=$1 AND ul.list_id = $2", noteListsTable, userLIstTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *NoteListPostgres) Update(userId, listId int, input rest_api_gin.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	var argId int64 = 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s nl SET %s FROM %s ul WHERE nl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		noteListsTable, setQuery, userLIstTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}
