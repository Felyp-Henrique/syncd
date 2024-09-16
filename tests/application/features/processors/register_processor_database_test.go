package processors_test

import (
	"Felyp-Henrique/syncd/src/application/domain/entities"
	"Felyp-Henrique/syncd/src/application/features/processors"
	"Felyp-Henrique/syncd/src/application/features/repositories"
	mock "Felyp-Henrique/syncd/tests/application/features/processors/mock"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestDataBaseRegisterProcessorExecute(t *testing.T) {
	t.Run("Testing DataBaseRegisterProcessor execution", func(t *testing.T) {
		var (
			mockRegisterProcessorRepository *mock.MockRegisterProcessorRepository = nil
			registerProcessor               processors.RegisterProcessor          = nil
			register                        entities.Register
		)
		register = getRegisterWithEmptyData()
		mockRegisterProcessorRepository = getMockableRegisterProcessorRepository(t)
		mockRegisterProcessorRepository.EXPECT().Execute(gomock.Any()).Return(nil)
		registerProcessor = getRegisterProcessorService(mockRegisterProcessorRepository)
		if err := registerProcessor.Execute(&register); err != nil {
			t.Errorf("DataBaseRegisterProcessor not work very well!")
		}
	})
}

func getRegisterWithEmptyData() entities.Register {
	return entities.NewRegister(map[string]any{})
}

func getRegisterProcessorService(registerProcessorRepository repositories.RegisterProcessorRepository) processors.RegisterProcessor {
	return processors.NewDataBaseRegisterProcessor(registerProcessorRepository)
}

//go:generate mockgen -source=../../../../src/application/features/repositories/register_processor.go -destination=mock/register_processor.go
func getMockableRegisterProcessorRepository(t *testing.T) *mock.MockRegisterProcessorRepository {
	var (
		registerProcessorRepositoryMockController *gomock.Controller                    = nil
		registerProcessorRepositoryMock           *mock.MockRegisterProcessorRepository = nil
	)
	registerProcessorRepositoryMockController = gomock.NewController(t)
	registerProcessorRepositoryMock = mock.NewMockRegisterProcessorRepository(registerProcessorRepositoryMockController)
	return registerProcessorRepositoryMock
}
