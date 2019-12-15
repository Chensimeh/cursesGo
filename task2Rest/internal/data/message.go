package data

import (
	"database/sql"
	"fmt"
	"log"
)

type Message struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	UserId  int    `json:"user_id"`
}

type msgData struct {
	db *sql.DB
}

type IMsgData interface {
	FindMsgByID() (*Message, error)
	GetUserMsg() ([]Message, error)
	CreateMsg() error
	UpdatedByID() error
	DeleteByID() error
}

func NewMsgData() *msgData {

	db, err := (&Store{}).Open()
	if err != nil {
		log.Fatal(err)
	}
	return &msgData{db: db}
}

func (m *Message) FindMsgByID() (*Message, error) {
	db := NewMsgData().db
	defer db.Close()
	if err := db.QueryRow(
		"SELECT id, message, user_id FROM messages WHERE id=$1",
		m.Id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.UserId,
	); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Message) GetUserMsg() ([]Message, error) {
	var msgs []Message
	db := NewMsgData().db
	defer db.Close()
	rows, err := db.Query(
		"SELECT * FROM messages WHERE user_id=$1",
		m.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&m.Id,
			&m.Message,
			&m.UserId,
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		msgs = append(msgs, *m)
	}
	return msgs, nil
}

func (m *Message) CreateMsg() error {
	db := NewMsgData().db
	defer db.Close()
	if err := db.QueryRow(
		"INSERT INTO messages(message, user_id) VALUES ($1,$2) RETURNING id",
		m.Message,
		m.UserId,
	).Scan(&m.Id); err != nil {
		return err
	}
	return nil
}

func (m *Message) UpdatedByID() error {
	db := NewMsgData().db
	defer db.Close()
	if err := db.QueryRow(
		"UPDATE messages SET message=$1, user_id=$2  WHERE id=$3 RETURNING id, message, user_id",
		m.Message,
		m.UserId,
		m.Id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.UserId,
	); err != nil {
		return err
	}
	return nil
}

func (m *Message) DeleteByID() error {
	db := NewMsgData().db
	defer db.Close()
	if err := db.QueryRow(
		"DELETE FROM messages WHERE id=$1 RETURNING id, message, user_id",
		m.Id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.UserId,
	); err != nil {
		return err
	}
	return nil
}
