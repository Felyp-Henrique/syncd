package databases

import "database/sql"

var (
	_databases []*sql.DB = make([]*sql.DB, 0)
)

func ServicesAll() []*sql.DB {
	return _databases
}

func ServicesCreate() {

}
