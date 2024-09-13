package processors_test

import (
	"Felyp-Henrique/syncd/src/application/domain/entities"
	"Felyp-Henrique/syncd/src/application/features/processors"
	infrastructureErrors "Felyp-Henrique/syncd/src/infrastructure/errors"
	mock "Felyp-Henrique/syncd/tests/application/features/processors/mock"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=../../../../src/infrastructure/database/database.go -destination=mock/database.go
func TestDataBaseRegisterProcessorExecute(t *testing.T) {
	t.Run("Testing DataBaseRegisterProcessor execution", func(t *testing.T) {
		var (
			dataBaseConnectionMockController *gomock.Controller              = nil
			dataBaseConnectionMock           *mock.MockDataBase              = nil
			dataBaseConnection               *sqlx.DB                        = nil
			registerProcessor                processors.RegisterProcessor    = nil
			registerProcessorError           error                           = nil
			registerProcessorSyncError       *infrastructureErrors.SyncError = nil
			register                         entities.Register
		)
		register = entities.NewRegister(map[string]any{})
		dataBaseConnectionMockController = gomock.NewController(t)
		dataBaseConnectionMock = mock.NewMockDataBase(dataBaseConnectionMockController)
		dataBaseConnection, _ = dataBaseConnectionMock.GetConnection()
		registerProcessor = processors.NewDataBaseRegisterProcessor(dataBaseConnection)
		registerProcessorError = registerProcessor.Execute(&register)
		if errors.As(registerProcessorError, registerProcessorSyncError) {
			t.Errorf("DataBaseRegisterProcessor not work very well!")
		}
	})
}
