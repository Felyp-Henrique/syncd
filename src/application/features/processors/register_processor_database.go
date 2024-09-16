package processors

import (
	"Felyp-Henrique/syncd/src/application/domain/entities"
	"Felyp-Henrique/syncd/src/application/features/repositories"
)

type DataBaseRegisterProcessor struct {
	registerProcessorRepository repositories.RegisterProcessorRepository
}

func NewDataBaseRegisterProcessor(registerProcessorRepository repositories.RegisterProcessorRepository) RegisterProcessor {
	return &DataBaseRegisterProcessor{
		registerProcessorRepository: registerProcessorRepository,
	}
}

func (d *DataBaseRegisterProcessor) Execute(register *entities.Register) error {
	return d.registerProcessorRepository.Execute("")
}
