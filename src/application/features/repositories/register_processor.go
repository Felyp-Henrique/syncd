package repositories

type RegisterProcessorRepository[Q any] interface {
	Execute(query Q) error
}
