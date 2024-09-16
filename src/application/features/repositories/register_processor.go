package repositories

type RegisterProcessorRepository interface {
	Execute(query string) error
}
