// Package model contains the types for schema 'public'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Account represents a row from 'public.accounts'.
type Account struct {
	ID          int64          `json:"id"`          // id
	SecretID    sql.NullString `json:"secret_id"`   // secret_id
	SecretKey   sql.NullString `json:"secret_key"`  // secret_key
	Description sql.NullString `json:"description"` // description
	CreatedAt   pq.NullTime    `json:"created_at"`  // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Account exists in the database.
func (a *Account) Exists() bool {
	return a._exists
}

// Deleted provides information if the Account has been deleted from the database.
func (a *Account) Deleted() bool {
	return a._deleted
}

// Insert inserts the Account to the database.
func (a *Account) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.accounts (` +
		`secret_id, secret_key, description, created_at` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, a.SecretID, a.SecretKey, a.Description, a.CreatedAt)
	err = db.QueryRow(sqlstr, a.SecretID, a.SecretKey, a.Description, a.CreatedAt).Scan(&a.ID)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Account in the database.
func (a *Account) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.accounts SET (` +
		`secret_id, secret_key, description, created_at` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE id = $5`

	// run query
	XOLog(sqlstr, a.SecretID, a.SecretKey, a.Description, a.CreatedAt, a.ID)
	_, err = db.Exec(sqlstr, a.SecretID, a.SecretKey, a.Description, a.CreatedAt, a.ID)
	return err
}

// Save saves the Account to the database.
func (a *Account) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Account.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Account) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.accounts (` +
		`id, secret_id, secret_key, description, created_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, secret_id, secret_key, description, created_at` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.secret_id, EXCLUDED.secret_key, EXCLUDED.description, EXCLUDED.created_at` +
		`)`

	// run query
	XOLog(sqlstr, a.ID, a.SecretID, a.SecretKey, a.Description, a.CreatedAt)
	_, err = db.Exec(sqlstr, a.ID, a.SecretID, a.SecretKey, a.Description, a.CreatedAt)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Account from the database.
func (a *Account) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.accounts WHERE id = $1`

	// run query
	XOLog(sqlstr, a.ID)
	_, err = db.Exec(sqlstr, a.ID)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AccountsQuery returns offset-limit rows from 'public.accounts' filte by filter,
// ordered by "id" in descending order.
func AccountFilter(db XODB, filter string, offset, limit int) ([]*Account, error) {
	sqlstr := `SELECT ` +
		`id, secret_id, secret_key, description, created_at` +
		`FROM public.accounts `

	if filter != "" {
		sqlstr = sqlstr + " WHERE " + filter
	}

	sqlstr = sqlstr + " order by id desc offset $1 limit $2"

	XOLog(sqlstr, offset, limit)
	q, err := db.Query(sqlstr, offset, limit)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*Account
	for q.Next() {
		a := Account{}

		// scan
		err = q.Scan(&a.ID, &a.SecretID, &a.SecretKey, &a.Description, &a.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &a)
	}

	return res, nil
} // AccountByID retrieves a row from 'public.accounts' as a Account.
//
// Generated from index 'accounts_pkey'.
func AccountByID(db XODB, id int64) (*Account, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, secret_id, secret_key, description, created_at ` +
		`FROM public.accounts ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	a := Account{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&a.ID, &a.SecretID, &a.SecretKey, &a.Description, &a.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// AccountBySecretID retrieves a row from 'public.accounts' as a Account.
//
// Generated from index 'accounts_secret_id_key'.
func AccountBySecretID(db XODB, secretID sql.NullString) (*Account, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, secret_id, secret_key, description, created_at ` +
		`FROM public.accounts ` +
		`WHERE secret_id = $1`

	// run query
	XOLog(sqlstr, secretID)
	a := Account{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, secretID).Scan(&a.ID, &a.SecretID, &a.SecretKey, &a.Description, &a.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &a, nil
}