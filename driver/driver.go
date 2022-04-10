package driver

import "database/sql"


type DB struct {
	Con sql.Conn
}