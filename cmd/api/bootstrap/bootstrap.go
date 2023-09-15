package bootstrap

import (
	"database/sql"
	"fmt"

	"gohex.com/m/internal/platform/server"
	"gohex.com/m/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "user"
	dbPass = "pass"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "courses"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}