package util

import (
	"database/sql"

	"github.com/lib/pq"
)

// user queries
const (
	SIGNUP           = "INSERT INTO USERS (FIRST_NAME, LAST_NAME, EMAIL, PASSWORD, MONTHLY_SALARY, CREATED_AT) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID"
	FETCH_USER_BY_ID = "SELECT ID, PASSWORD FROM USERS WHERE EMAIL=$1"
	IS_USER_EXISTS   = "SELECT EXISTS(SELECT 1 FROM USERS WHERE ID=$1"
	UPDATE_BUDGET    = "update users set budget_per_month=monthly_salary - (select sum(amount) from savings where user_id=$1) where id=$1"
)

type ExecFunc func(stmt *sql.Stmt) error

func BulkInsert(db *sql.DB, tableName string, f ExecFunc, columns ...string) error {
	txn, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := txn.Prepare(pq.CopyIn(tableName, columns...))
	if err != nil {
		return err
	}
	if err := f(stmt); err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	if err := stmt.Close(); err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}
