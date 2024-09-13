package entities_test

import (
	"Felyp-Henrique/syncd/src/application/domain/entities"
	"testing"
)

func TestNewUser(t *testing.T) {
	t.Run("Testing a new user instance", func(t *testing.T) {
		if entities.NewUser() == nil {
			t.Errorf("User can't be nil!")
		}
	})
}
