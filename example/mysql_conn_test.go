package example

import (
	"goPro4/db"
	"testing"
)

func TestGetMysqlConn(t *testing.T) {
	db.GetMysqlConnection()
}
