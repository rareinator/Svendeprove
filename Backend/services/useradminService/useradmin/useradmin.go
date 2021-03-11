package useradmin

import (
	"context"
	"fmt"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
)

type UseradminServer struct {
	UnimplementedUseradminServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (u *UseradminServer) GetHealth(ctx context.Context, e *UAEmpty) (*UAHealth, error) {
	return &UAHealth{Message: fmt.Sprintf("Useradmin service is up and running on: %v 🚀", u.ListenAddress)}, nil
}

func (u *UseradminServer) GetEmployee(ctx context.Context, er *EmployeeRequest) (*Employee, error) {
	dbEmployee, err := u.DB.GetEmployee(er.EmployeeId)
	if err != nil {
		return nil, err
	}

	result := Employee{
		EmployeeId:   dbEmployee.EmployeeId,
		Name:         dbEmployee.Name,
		WorktitleId:  dbEmployee.WorktitleId,
		DepartmentId: dbEmployee.DepartmentId,
		Username:     dbEmployee.Username,
	}

	return &result, nil
}
