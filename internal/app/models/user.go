package models

import (
	"database/sql"

	pbv1 "github.com/ThorifArtanel/grpc-sandbox/gen/proto/v1"
	"github.com/google/uuid"
)

type Users struct {
	DB *sql.DB
}

func (m *Users) All(ret *pbv1.UserGetResponse) error {
	rows, err := m.DB.Query("SELECT id, firstname, lastname FROM users")
	if err != nil {
		return err
	}
	for rows.Next() {
		var usr pbv1.User
		err = rows.Scan(&usr.Id, &usr.Firstname, &usr.Lastname)
		if err != nil {
			return err
		}

		ret.Users = append(ret.Users, &usr)
	}

	return nil
}

func (m *Users) One(id string, ret *pbv1.UserOneResponse) error {
	var usr pbv1.User
	err := m.DB.QueryRow("SELECT firstname, lastname FROM users WHERE id=$1", id).Scan(&usr.Firstname, &usr.Lastname)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	ret.User = &usr

	return nil
}

func (m *Users) Insert(req *pbv1.UserInsertRequest) error {
	_, err := m.DB.Exec("INSERT INTO users(id, firstname, lastname) VALUES ($1, $2, $3)", uuid.NewString(), req.User.GetFirstname(), req.User.GetLastname())
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (m *Users) Update(id string, req *pbv1.UserUpdateRequest) error {
	_, err := m.DB.Exec("UPDATE users SET firstname=$2, lastname=$3 WHERE id=$1", id, req.User.GetFirstname(), req.User.GetLastname())
	if err != nil {
		return err
	}

	return nil
}

func (m *Users) Delete(id string) error {
	_, err := m.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
