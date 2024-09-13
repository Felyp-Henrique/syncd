package processors

import (
	"Felyp-Henrique/syncd/src/application/domain/entities"

	"github.com/jmoiron/sqlx"
)

type DataBaseRegisterProcessor struct {
	dataBaseConnection *sqlx.DB
}

func NewDataBaseRegisterProcessor(dataBaseConnection *sqlx.DB) RegisterProcessor {
	return &DataBaseRegisterProcessor{
		dataBaseConnection: dataBaseConnection,
	}
}

func (d *DataBaseRegisterProcessor) Execute(register *entities.Register) error {
	return nil
}
