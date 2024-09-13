package processors

import "Felyp-Henrique/syncd/src/application/domain/entities"

type RegisterProcessor interface {
	Execute(*entities.Register) error
}
