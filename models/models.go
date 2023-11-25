package models

import (
	"context"
	"database/sql"
	"log"
	"omni/utils"
)

type User struct {
	ID       int64  `json:"ID"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserModel struct {
	DB  *sql.DB
	Log *log.Logger
}

func (u UserModel) ListUsers() ([]User, error) {
	rows, err := u.DB.Query("SELECT id, Name FROM Users")
	if err != nil {
		log.Fatal()
		return nil, err
	}

	defer rows.Close()
	var us []User
	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return us, nil
}

func (u UserModel) ListUserById(id string) (User, error) {
	var err error
	row := u.DB.QueryRow("SELECT id, name FROM Users where id=?", id)

	user := User{}

	if err := row.Scan(&user.ID, &user.Name); err != nil {
		log.Fatalf("error while scanning row %v", err)

	}

	if err := row.Err(); err != nil {
		log.Fatalf("error while scanning row %v", err)
	}
	return user, err
}

func (u UserModel) CreateUser(usr User) (User, error) {
	q := "INSERT INTO Users (Name, Password) VALUES (?, ?)"
	hash, err := utils.HashPassword(usr.Password)

	if err != nil {
		log.Fatalf("error when hashing password: %v", err)
	}

	r, err := u.DB.ExecContext(context.Background(), q, usr.Name, hash)

	if err != nil {
		log.Fatalf("error while inserting user: %v", err)
	}

	id, err := r.LastInsertId()

	if err != nil {
		log.Fatalf("cannot retrieve last inserted id: %s", err)
	}

	user := User{
		ID:   id,
		Name: usr.Name,
	}

	return user, nil
}
