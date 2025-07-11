package employee

import "database/sql"

type Repository interface {
	GetAll() ([]Employee, error)
}

type employeeRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &employeeRepository{db}
}

func (r *employeeRepository) GetAll() ([]Employee, error) {
	rows, err := r.db.Query("SELECT id, name FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Name); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}
