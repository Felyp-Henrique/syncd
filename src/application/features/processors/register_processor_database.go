package processors

import (
	"Felyp-Henrique/syncd/src/application/domain/entities"
	"Felyp-Henrique/syncd/src/application/features/repositories"
)

type DataBaseRegisterProcessor struct {
	registerProcessorRepository repositories.RegisterProcessorRepository[any]
}

func NewDataBaseRegisterProcessor(registerProcessorRepository repositories.RegisterProcessorRepository[any]) RegisterProcessor {
	return &DataBaseRegisterProcessor{
		registerProcessorRepository: registerProcessorRepository,
	}
}

func (d *DataBaseRegisterProcessor) Execute(register *entities.Register) error {
	return nil
}
