package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	//to connect mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func connectDB(args ...string) *sqlx.DB {
	dbConfig := Load("db")

	if dbConfig == nil {
		return nil
	}

	var user, pass string

	_, isExistUser := dbConfig["user"]
	_, isExistPass := dbConfig["pass"]

	if len(args) == 2 {
		user = args[0]
		pass = args[1]
	}

	if isExistUser && isExistPass {
		user = dbConfig["user"].(string)
		pass = dbConfig["pass"].(string)
	}

	if pass == "" || user == "" {
		log.Fatal("DB: No exist user and pass")
	}

	connectSetting := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?timeout=5s",
		user,
		pass,
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["name"],
	)

	for {
		db, err := sqlx.Connect("mysql", connectSetting)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Minute)
			continue
		}

		return db
	}
}

// Select using this DB.
// Any placeholder parameters are replaced with supplied args.
// logging error
func (app Data) Select(dest interface{}, query string, args ...interface{}) error {
	if err := app.db.Unsafe().Select(dest, query, args...); err != nil {
		app.Err("Select\nquery:\n%s\nerr: %v\n", query, err)
		return err
	}
	return nil
}

// Get using this DB.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
// logging error
func (app Data) Get(dest interface{}, query string, args ...interface{}) error {
	if err := app.db.Unsafe().Get(dest, query, args...); err != nil {
		app.Err("Get\nquery:\n%s\nerr: %v\n %v\n", query, err, args)
		return err
	}
	return nil
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// logging error
func (app Data) Exec(query string, args ...interface{}) error {
	if _, err := app.db.Unsafe().Exec(query, args...); err != nil {
		app.Err("Exec\nquery:\n%s\nerr: %v\n %v\n", query, err, args)
		return err
	}
	return nil
}

// NamedExec using this DB.
// Any named placeholder parameters are replaced with fields from arg.
// logging error
func (app Data) NamedExec(query string, arg interface{}) error {
	if _, err := app.db.Unsafe().NamedExec(query, arg); err != nil {
		app.Err("Exec\nquery:\n%s\nerr: %v\n %v\n", query, err, arg)
		return err
	}
	return nil
}

// Prepare creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's Close method
// when the statement is no longer needed.
func (app Data) Prepare(query string) (connect *sql.Stmt, err error) {
	if connect, err = app.db.Unsafe().Prepare(query); err != nil {
		app.Err("Prepare\nquery:\n%s\nerr: %s\n", query, err)
		return nil, err
	}
	return connect, nil
}
