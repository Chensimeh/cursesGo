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

