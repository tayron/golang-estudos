package mock

type Service interface {
	SaveAccount(person *domain.Account) error
	GetAccount(ID int64) (*domain.Account, error)
}
