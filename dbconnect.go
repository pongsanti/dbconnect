package dbconnect

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConnect struct {
	host             string
	dbname           string
	user             string
	password         string
	options          string
	connectionString string
	Db               *sql.DB
}

func NewDBConnect(host string, dbname string, user string, password string, options string) (*DBConnect, error) {

	connectionString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s %s",
		host, dbname, user, password, options)
	log.Print(connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to database!")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database!")
		return nil, err
	}

	con := &DBConnect{
		host:             host,
		dbname:           dbname,
		user:             user,
		password:         password,
		options:          options,
		connectionString: connectionString,
		Db:               db,
	}

	return con, nil
}

func (con *DBConnect) Close() {
	con.Db.Close()
}
