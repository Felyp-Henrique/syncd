package repositories

import "github.com/jmoiron/sqlx"

type DataBaseRegisterProcessorRepository struct {
	dataBase *sqlx.DB
}

func NewDataBaseRegisterProcessorRepository(dataBase *sqlx.DB) RegisterProcessorRepository[string] {
	return &DataBaseRegisterProcessorRepository{
		dataBase: dataBase,
	}
}

func (d *DataBaseRegisterProcessorRepository) Execute(query string) error {
	return nil
}
