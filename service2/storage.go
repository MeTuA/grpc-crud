package main

import (
	"github.com/jmoiron/sqlx"
	datapb "service2/proto"
)

type Storage struct {
	dbConn *sqlx.DB
}

func NewStorage(dbConn *sqlx.DB) *Storage {
	return &Storage{dbConn: dbConn}
}

func (s *Storage) Delete(id int) error {
	query := `DELETE FROM data WHERE id=:id`
	queryArgs := map[string]interface{}{"id":id}

	_, err := s.dbConn.NamedExec(query, queryArgs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Get(id int) (*datapb.Data, error) {
	data := &datapb.Data{}

	query := `SELECT id,user_id as userId, title, body FROM data WHERE id=$1`

	err := s.dbConn.Get(data, query, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Storage) GetAll() (*datapb.ListData, error) {
	data := &datapb.ListData{}

	query := `SELECT id, user_id as userId, title, body FROM data;`
	err := s.dbConn.Select(&data.Data, query)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Storage) Update(data *datapb.Data) error {
	query := `UPDATE data SET user_id=:userId, title=:title, body=:body WHERE id=:id`
	queryArgs := map[string]interface{}{"userId": data.GetUserId(), "title": data.GetTitle(), "body": data.GetBody(), "id": data.GetId()}

	_, err := s.dbConn.NamedExec(query, queryArgs)
	if err != nil {
		return err
	}

	return nil
}
