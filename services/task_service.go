package services

import (
	"database/sql"
	"errors"
	"fmt"
	"gin_test/db"
	"gin_test/models"
)

type TaskService struct {
	DB *sql.DB
}

func (service *TaskService) GetById(id int) (models.Todo, error) {
	var todo models.Todo

	stmt, err := db.Get().Prepare("SELECT * FROM todo where id = ?")
	if err != nil {
		return todo, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return todo, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.Done,
			&todo.Created,
			&todo.Updated,
		)
		if err != nil {
			return todo, err
		}
	}

	return todo, nil
}

func (service *TaskService) GetAll() (models.Todos, error) {
	var todos models.Todos
	stmt, err := db.Get().Prepare("SELECT * FROM todo")
	if err != nil {
		return todos, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return todos, err
	}

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.Done,
			&todo.Created,
			&todo.Updated,
		)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (service *TaskService) Insert(name string) (int64, error) {
	var id int64
	stmt, err := db.Get().Prepare("INSERT INTO todo(name) values (?)")
	if err != nil {
		return id, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		return id, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, err
}

func (service *TaskService) Update(id int, todo models.Todo) error {
	stmt, err := db.Get().Prepare("UPDATE todo SET name = ?, done = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Name, todo.Done, id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if !(rowAffected > 0) {
		return errors.New(fmt.Sprintf("{row_affected=%d}", rowAffected))
	}
	return nil
}

func (service *TaskService) Delete(id int) error {
	stmt, err := db.Get().Prepare("DELETE FROM Todo WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if !(rowAffected > 0) {
		return errors.New(fmt.Sprintf("no such Todo id = %d", id))
	}

	return nil
}
