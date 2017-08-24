package main

import (
	"upper.io/db.v3/mysql"
)

var c = GetDatabaseSettings()
var settings = mysql.ConnectionURL{
	Host:     c.Host,
	Database: c.Name,
	User:     c.User,
	Password: c.Password,
}


sess, err := mysql.Open(settings)
  if err != nil {
    log.Fatalf("db.Open(): %q\n", err)
  }
  defer sess.Close()