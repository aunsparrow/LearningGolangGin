package Models

import (
	"database/sql"
)

type UserModel struct {
	UserId    string `json:"UserId"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
	Active    int    `json:"Active""`
	CreatedAt string `json:"CreatedAt"`
}

type UserModelDb struct {
	UserId    sql.NullString `json:"UserId"`
	FirstName sql.NullString `json:"FirstName"`
	LastName  sql.NullString `json:"LastName"`
	Age       sql.NullInt32  `json:"Age"`
	Active    sql.NullInt32  `json:"Active""`
	CreatedAt sql.NullString `json:"CreatedAt"`
}
