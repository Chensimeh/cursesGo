package store

import (
	"fmt"
	"github.com/trzhensimekh/cursesGo/task2Rest/model"
)

// UserRep ...
type UserRep struct {
Store *Store
}

// Create ...
func (r *UserRep) Create(u *model.User)(*model.User,error){
if err := r.Store.db.QueryRow(
	"INSERT INTO users(id, firstname, lastname, email) VALUES ($1,$2,$3,$4) RETURNING id",
	u.Id,
	u.FistName,
	u.LastName,
	u.Email,
	).Scan(&u.Id); err!=nil{
	return nil,err
}
return u, nil
}

//FindByID ...
func (r *UserRep) FindByID(id int)(*model.User,error){
	u:=&model.User{}
	if err:= r.Store.db.QueryRow(
		"SELECT id, firstname, lastname, email FROM USERS WHERE id=$1",
		id,
		).Scan(
			&u.Id,
			&u.FistName,
			&u.LastName,
			&u.Email,
			); err!=nil {
	return nil,err
	}

	return u, nil
}

func (r *UserRep) GetUsers() ([]model.User, error) {
	var users []model.User
	u := new(model.User)
	rows,err := r.Store.db.Query(
		"SELECT * FROM USERS")
	if err!=nil{
		return nil, err
	}

	for rows.Next() {
		err:=rows.Scan(
			&u.Id,
			&u.FistName,
			&u.LastName,
			&u.Email,
		)
		if err!=nil{
			fmt.Println(err.Error())
		}
		users = append(users, *u)
	}
	return users,nil
}