package app

import (
	"database/sql"
	"fmt"
	"time"

	//to connect DB
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func connectDB(args ...string) *sqlx.DB {
	dbConfig := Load("db")

	user := dbConfig["user"].(string)
	pass := dbConfig["pass"].(string)

	if len(args) == 2 {
		user = args[0]
		pass = args[1]
	}

	connectSetting := fmt.Sprintf(
		"%s:%s@tcp(localhost:3306)/%s",
		user,
		pass,
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
func (app Data) Select(dest interface{}, query string, args ...interface{}) {
	if err := app.db.Unsafe().Select(dest, query, args...); err != nil {
		app.Err("Select\nquery:\n%s\nerr: %v", query, err)
	}
}

// Get using this DB.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
// logging error
func (app Data) Get(dest interface{}, query string, args ...interface{}) {
	if err := app.db.Unsafe().Get(dest, query, args...); err != nil {
		app.Err("Select\nquery:\n%s\nerr: %v", query, err)
	}
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// logging error
func (app Data) Exec(query string, args ...interface{}) {
	if _, err := app.db.Unsafe().Exec(query, args...); err != nil {
		app.Err("Select\nquery:\n%s\nerr: %v", query, err)
	}
}

// Prepare creates a prepared statement for later queries or executions.
// Multiple queries or executions may be run concurrently from the
// returned statement.
// The caller must call the statement's Close method
// when the statement is no longer needed.
// logging error
func (app Data) Prepare(query string) *sql.Stmt {
	conn, err := app.db.Prepare(query)
	if err != nil {
		app.Err("Prepare\nquery:\n%s\nerr: %v", query, err)
	}
	return conn
}
