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
	//TODO: implement okta sdk

	result := Employee{
		Name:         "Morten Nissen",
		WorktitleId:  1,
		DepartmentId: 1,
		Username:     er.Employee,
	}

	return &result, nil
}
