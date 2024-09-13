package services_test

import (
	"Felyp-Henrique/syncd/src/application/features/services"
	"testing"
)

func TestNewUserDataBaseService(t *testing.T) {
	t.Run("Testing a new UsersDataBaseService instance", func(t *testing.T) {
		if services.NewUserDataBaseService(nil) == nil {
			t.Errorf("UsersDataBaseService can't be nil!")
		}
	})
}
