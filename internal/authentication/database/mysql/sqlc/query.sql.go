// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :execresult
INSERT INTO Users (UserID, FirstName, LastName, Email, PhoneNo, Password, Role)
VALUES (?,?,?,?,?,?,?)
`

type CreateAccountParams struct {
	Userid    string
	Firstname string
	Lastname  string
	Email     string
	Phoneno   string
	Password  string
	Role      sql.NullString
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAccount,
		arg.Userid,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Phoneno,
		arg.Password,
		arg.Role,
	)
}

const getPassword = `-- name: GetPassword :one
SELECT Password from Users
WHERE Email = ?
`

func (q *Queries) GetPassword(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRowContext(ctx, getPassword, email)
	var password string
	err := row.Scan(&password)
	return password, err
}

const getRole = `-- name: GetRole :one
SELECT RoleID from Roles
WHERE RoleName = ?
`

func (q *Queries) GetRole(ctx context.Context, rolename string) (string, error) {
	row := q.db.QueryRowContext(ctx, getRole, rolename)
	var roleid string
	err := row.Scan(&roleid)
	return roleid, err
}
