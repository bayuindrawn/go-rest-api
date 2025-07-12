package employee

import "sync"

type Service interface {
	GetEmployees() ([]Employee, error)
	IncrementCounter()
	GetCounter() int
}

type employeeService struct {
	repo    Repository
	counter int
	mu      sync.Mutex
}

func NewService(repo Repository) Service {
	return &employeeService{
		repo:    repo,
		counter: 0,
	}
}

func (s *employeeService) GetEmployees() ([]Employee, error) {
	return s.repo.GetAll()
}

func (s *employeeService) IncrementCounter() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
}

func (s *employeeService) GetCounter() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.counter
}
