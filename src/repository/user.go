package repository

import (
	"database/sql"
	"devbook-api/src/model"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db}
}

func (r User) Create(user model.User) (uint64, error) {
	statement, error := r.db.Prepare("insert into user (name, nick, email, password) values (?, ?, ?, ?)"); if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password); if error != nil {
		return 0, error
	}

	lastInsertId, error := result.LastInsertId(); if error != nil {
		return 0, error
	}

	return uint64(lastInsertId), nil
}

func (r User) FindAll() ([]model.User, error) {
	rows, error := r.db.Query("select id, name, nick, email, password from user"); if error != nil {
		return nil, error
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		if error = rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

func (r User) FindById(id uint64) (model.User, error) {
	rows, error := r.db.Query("select id, name, nick, email, password from user where id = ?", id); if error != nil {
		return model.User{}, error
	}
	defer rows.Close()

	var user model.User

	if rows.Next() {
		if error = rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password); error != nil {
			return model.User{}, error
		}
	} else {
		return model.User{}, nil
	}

	return user, nil
}
