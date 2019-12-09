package store

import (
	"fmt"
	"github.com/trzhensimekh/cursesGo/task2Rest/model"
)

type MsgRep struct{
	Store *Store
}

//FindByID ...

func (r *MsgRep) FindByID(id int) (*model.Message, error) {
	m:=&model.Message{}
	if err:= r.Store.db.QueryRow(
		"SELECT id, message, user_id FROM messages WHERE id=$1",
		id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.User_id,
	); err!=nil {
		return nil,err
	}
	return m, nil
}

// GetUserMsg ...
func (r *MsgRep) GetUserMsg(id int) ([]model.Message, error) {
	var msgs []model.Message
	m := new(model.Message)
	rows,err := r.Store.db.Query(
		"SELECT * FROM messages WHERE user_id=$1",
		id)
	if err!=nil{
		return nil, err
	}
	for rows.Next() {
		err:=rows.Scan(
			&m.Id,
			&m.Message,
			&m.User_id,
		)
		if err!=nil{
			fmt.Println(err.Error())
		}
		msgs = append(msgs, *m)
	}
	return msgs,nil
}

func (r *MsgRep) CreateUser(m *model.Message) error {
	if err := r.Store.db.QueryRow(
		"INSERT INTO messages(message, user_id) VALUES ($1,$2) RETURNING id",
		m.Message,
		m.User_id,
	).Scan(&m.Id); err!=nil{
		return err
	}
	return nil
}

func (r *MsgRep) UpdatedByID(m *model.Message) error {
	if err:= r.Store.db.QueryRow(
		"UPDATE messages SET message=$1, user_id=$2  WHERE id=$3 RETURNING id, message, user_id",
		m.Message,
		m.User_id,
		m.Id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.User_id,
	); err!=nil {
		return err
	}
	return nil
}

func (r *MsgRep) DeleteByID(m *model.Message) error{
	if err:= r.Store.db.QueryRow(
		"DELETE FROM messages WHERE id=$1 RETURNING id, message, user_id",
		m.Id,
	).Scan(
		&m.Id,
		&m.Message,
		&m.User_id,
	); err!=nil {
		return err
	}
	return nil
}



