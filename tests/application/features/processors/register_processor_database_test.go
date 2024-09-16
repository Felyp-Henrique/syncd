package processors_test

import (
	"Felyp-Henrique/syncd/src/application/domain/entities"
	"Felyp-Henrique/syncd/src/application/features/processors"
	"Felyp-Henrique/syncd/src/application/features/repositories"
	infrastructureErrors "Felyp-Henrique/syncd/src/infrastructure/errors"
	mock "Felyp-Henrique/syncd/tests/application/features/processors/mock"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestDataBaseRegisterProcessorExecute(t *testing.T) {
	t.Run("Testing DataBaseRegisterProcessor execution", func(t *testing.T) {
		var (
			mockRegisterProcessorRepository *mock.MockRegisterProcessorRepository[any] = nil
			registerProcessor               processors.RegisterProcessor               = nil
			registerProcessorError          error                                      = nil
			registerProcessorSyncError      *infrastructureErrors.SyncError            = nil
			register                        entities.Register
		)
		register = getRegisterWithEmptyData()
		mockRegisterProcessorRepository = getMockableRegisterProcessorRepository(t)
		registerProcessor = getRegisterProcessorService(mockRegisterProcessorRepository)
		mockRegisterProcessorRepository.EXPECT().Execute(nil).Return(nil)
		registerProcessorError = registerProcessor.Execute(&register)
		if errors.As(registerProcessorError, registerProcessorSyncError) {
			t.Errorf("DataBaseRegisterProcessor not work very well!")
		}
	})
}

func getRegisterWithEmptyData() entities.Register {
	return entities.NewRegister(map[string]any{})
}

func getRegisterProcessorService(registerProcessorRepository repositories.RegisterProcessorRepository[any]) processors.RegisterProcessor {
	return processors.NewDataBaseRegisterProcessor(registerProcessorRepository)
}

//go:generate mockgen -source=../../../../src/application/features/repositories/register_processor.go -destination=mock/register_processor.go
func getMockableRegisterProcessorRepository(t *testing.T) *mock.MockRegisterProcessorRepository[any] {
	var (
		registerProcessorRepositoryMockController *gomock.Controller                         = nil
		registerProcessorRepositoryMock           *mock.MockRegisterProcessorRepository[any] = nil
	)
	registerProcessorRepositoryMockController = gomock.NewController(t)
	registerProcessorRepositoryMock = mock.NewMockRegisterProcessorRepository[any](registerProcessorRepositoryMockController)
	return registerProcessorRepositoryMock
}
