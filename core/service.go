package Beer

type UseCase interface {
	ReadCase
	WriteCase
}

type ReadCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int) (*Beer, error)
}

type WriteCase interface {
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(b *Beer) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAll() ([]*Beer, error) {
	return nil, nil
}

func (s *Service) Get(ID int) (*Beer, error) {
	return nil, nil
}

func (s *Service) Store(b *Beer) error {
	return nil
}

func (s *Service) Update(b *Beer) error {
	return nil
}

func (s *Service) Remove(b *Beer) error {
	return nil
}
