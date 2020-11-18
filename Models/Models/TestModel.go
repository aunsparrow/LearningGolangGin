package Models

import (
	"database/sql"
)

type UserModel struct {
	UserId    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Status    int    `json:"status""`
	CreateAt  string `json:"createat"`
}

type UserModelDb struct {
	UserId    sql.NullString `json:"userid"`
	FirstName sql.NullString `json:"firstname"`
	LastName  sql.NullString `json:"lastname"`
	Age       sql.NullInt32  `json:"age"`
	Status    sql.NullInt32  `json:"status""`
	CreateAt  sql.NullString `json:"createat"`
}
