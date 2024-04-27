package databases

type Driver int

const (
	DriverPostgres Driver = iota
	DriverSqlite
)

type DataBase struct {
	Host   string
	Port   int
	Driver Driver
}

type Option func(*DataBase)

func NewDataBase(options ...Option) *DataBase {
	var (
		database *DataBase = &DataBase{}
	)
	for _, option := range options {
		option(database)
	}
	return database
}
