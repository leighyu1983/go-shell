package  daos

import (
	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

// exactly the same as the built-in
db = sqlx.Open("sqlite3", ":memory:")

// from a pre-existing sql.DB; note the required driverName
db = sqlx.NewDb(sql.Open("sqlite3", ":memory:"), "sqlite3")

// force a connection and test that it worked
err = db.Ping()

