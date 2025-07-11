package employee

type Service interface {
	GetEmployees() ([]Employee, error)
}

type employeeService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &employeeService{repo}
}

func (s *employeeService) GetEmployees() ([]Employee, error) {
	return s.repo.GetAll()
}
