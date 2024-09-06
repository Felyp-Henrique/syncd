package database

import (
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed configuration.yml
	_embedConfiguration []byte

	_databaseConnection *sql.DB = nil
)

func init() {
	var (
		configuration *DataSource = nil
		err           error       = nil
	)
	if configuration, err = GetConfigurationFromBytes(_embedConfiguration); err != nil {
		panic(err)
	}
	switch configuration.Driver {
	case DRIVER_SQLITE:
		_databaseConnection, err = GetConnectionSqlite(configuration.DataBase)
	case DRIVER_POSTGRESQL:
		_databaseConnection, err = GetConnectionPostgres(
			configuration.Host,
			configuration.Port,
			configuration.Username,
			configuration.Password,
			configuration.DataBase)
	}
	if err != nil {
		panic(err)
	}
}

const (
	DRIVER_SQLITE     string = "sqlite3"
	DRIVER_POSTGRESQL string = "postgres"
)

type DataSource struct {
	Host                  string `yml:"host"`
	Port                  int    `yml:"port"`
	DataBase              string `yml:"database"`
	Driver                string `yml:"driver"`
	Username              string `yml:"username"`
	Password              string `yml:"password"`
	ConnectionLifetimeMax int    `yml:"connection_lifetime_max"`
	ConnectionIdleMax     int    `yml:"connection_idle_max"`
	ConnectionOpenMax     int    `yml:"connection_open_max"`
}

func GetConfigurationFromBytes(bytes []byte) (*DataSource, error) {
	var (
		configuration *DataSource = new(DataSource)
	)
	if err := yaml.Unmarshal(bytes, configuration); err != nil {
		return nil, err
	} else {
		return configuration, nil
	}
}

func GetConnection() *sql.DB {
	return _databaseConnection
}

func GetConnectionSqlite(database string) (*sql.DB, error) {
	return sql.Open(DRIVER_SQLITE, database)
}

func GetConnectionPostgres(
	host string,
	port int,
	username string,
	password string,
	database string) (*sql.DB, error) {
	var datasource = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		username,
		password,
		database)
	return sql.Open(DRIVER_POSTGRESQL, datasource)
}
