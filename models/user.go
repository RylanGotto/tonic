package models

import (
	"context"
	"database/sql"
	"log"
	"omni/utils"
	"strconv"
)

type User struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) ListUsers() ([]User, error) {
	rows, err := u.DB.Query("SELECT id, Name, Email FROM Users")
	if err != nil {
		log.Fatal()
		return nil, err
	}

	defer rows.Close()
	var us []User
	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Name, &u.Email)
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

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return user, err
	}

	if err := row.Err(); err != nil {
		return user, err
	}
	return user, err
}

func (u UserModel) ListUserByEmail(email string) (User, error) {
	var err error

	row := u.DB.QueryRow("SELECT * FROM Users where email=?", email)

	user := User{}

	if err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email); err != nil {
		return user, err
	}

	if err := row.Err(); err != nil {
		return user, err
	}
	return user, err
}

func (u UserModel) CreateUser(usr User) (User, error) {
	q := "INSERT INTO Users (Name, Password, Email) VALUES (?, ?, ?)"
	hash, err := utils.HashPassword(usr.Password)

	if err != nil {
		return usr, err
	}

	r, err := u.DB.ExecContext(context.Background(), q, usr.Name, hash, usr.Email)

	if err != nil {
		return usr, err
	}

	id, err := r.LastInsertId()

	if err != nil {
		return usr, err
	}

	user := User{
		ID:    strconv.FormatInt(id, 10),
		Name:  usr.Name,
		Email: usr.Email,
	}

	return user, nil
}
