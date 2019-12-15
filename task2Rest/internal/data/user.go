package data

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	FistName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
}

type userData struct {
	db *sql.DB
}

type IUserData interface {
	FindByID() (*User, error)
	GetUsers() ([]User, error)
	CreateUser() error
	UpdatedByID() error
	DeleteByID() error
}

func NewUserData() *userData {
	db, err := (&Store{}).Open()
	if err != nil {
		log.Fatal(err)
	}
	return &userData{db: db}
}

func (*User) GetUsers() ([]User, error) {
	var users []User
	u := new(User)
	db := NewUserData().db
	defer db.Close()
	rows, err := db.Query(
		"SELECT * FROM USERS")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&u.Id,
			&u.FistName,
			&u.LastName,
			&u.Email,
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		users = append(users, *u)
	}
	return users, nil
}

func (u *User) CreateUser() error {
	db := NewUserData().db
	defer db.Close()
	if err := db.QueryRow(
		"INSERT INTO users(firstname, lastname, email) VALUES ($1,$2,$3) RETURNING id",
		u.FistName,
		u.LastName,
		u.Email,
	).Scan(&u.Id); err != nil {
		return err
	}
	return nil
}

func (u *User) UpdatedByID() error {
	db := NewUserData().db
	defer db.Close()
	if err := db.QueryRow(
		"UPDATE USERS SET firstname=$1 ,lastname=$2, email=$3  WHERE id=$4 RETURNING id, firstname ,lastname, email",
		u.FistName,
		u.LastName,
		u.Email,
		u.Id,
	).Scan(
		&u.Id,
		&u.FistName,
		&u.LastName,
		&u.Email,
	); err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteByID() error {
	db := NewUserData().db
	defer db.Close()
	if err := db.QueryRow(
		"DELETE FROM users WHERE id=$1 RETURNING id, firstname ,lastname, email",
		u.Id,
	).Scan(
		&u.Id,
		&u.FistName,
		&u.LastName,
		&u.Email,
	); err != nil {
		return err
	}
	return nil
}

func (u *User) FindByID() (*User, error) {
	db := NewUserData().db
	defer db.Close()
	if err := db.QueryRow(
		"SELECT id, firstname, lastname, email FROM USERS WHERE id=$1",
		u.Id,
	).Scan(
		&u.Id,
		&u.FistName,
		&u.LastName,
		&u.Email,
	); err != nil {
		return nil, err
	}

	return u, nil
}
