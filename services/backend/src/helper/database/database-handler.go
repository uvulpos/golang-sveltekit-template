package database

import "github.com/go-sqlx/sqlx"

// struct so you can modify the database handler more easily
// did this because people believe in these drivers / libraries as in god(s)
type Sql struct {
	DB *sqlx.DB
}
