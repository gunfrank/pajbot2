package users

import (
	"database/sql"
)

type server struct {
	sql *sql.DB
}

var _server server

func InitServer(_sql *sql.DB) {
	_server.sql = _sql
}
