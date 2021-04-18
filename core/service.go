package beer

import (
	"database/sql"
	"fmt"
	- "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	ReadCase
	WriteCase
}

type ReadCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int) (*Beer, error)
}

type WriteCase interface {
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(b *Beer) error
}

type Service struct{
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db
	}
}

func (s *Service) GetAll() ([]*Beer, error) {
	var result []*Beer

	rows, err := s.DB.Query(`SELECT id, name,type, style FROM beer`)
	if err != nil(
		return nil, err
	)
	defer rows.Close()

	for rows.Next() {
		var b beer

		err := rows.Scan(&b.ID, &b.Name, &b.type, &b.style)
		if err != nil{
			return nil, err
		}

		result = append(result, &b)
	}
	return result, nil
}

func (s *Service) Get(ID int) (*Beer, error) {
	var b beer 

	stmt, err := s.DB.Prepare("SELECT id, name, type, style from beer where id =?")
	if err != nil{
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(ID).Scan(&b.ID, &b.name, &b.type, &b.style)
	if err != nil{
		return nil, err
	}

	return stmt, nil
}

func (s *Service) Store(b *Beer) error {
	//begin transaction
	tx, err:= s.DB.Begin()
	if err != nil{
		return err
	}

	stmt, err := tx.Prepare(`insert into beer(id, name, type, style) values(?,?,?,?)`)
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err := stmt.Exec(b.ID, b.name, b.type, b.style)
	if err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) Update(b *Beer) error {
	if b.ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil{
		return err
	}

	stmt, err := tx.Prepare("UPDATE beer set name=?,type=?,style=? where id=?")
	if err != nil{
		return err
	}

	_, err := stmt.Exec(b.name, b.type, b.style, b.ID)
	if err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) Remove(b *Beer) error {
	if b.ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil{
		return err
	}

	stmt, err := tx.Prepare("delete beer where id =?")
	if err != nil{
		return err
	}

	_, err := stmt.Exec(b.ID)
	if err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
