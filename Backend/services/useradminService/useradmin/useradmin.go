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

func (u *UseradminServer) GetUser(ctx context.Context, er *UserRequest) (*UAUser, error) {
	//TODO: implement okta sdk

	result := UAUser{
		Name:     "Morten Nissen",
		Username: er.Username,
	}

	return &result, nil
}

func (u *UseradminServer) GetHospitals(ctx context.Context, e *UAEmpty) (*Hospitals, error) {
	hospitals := Hospitals{
		Hospitals: make([]*UAHospital, 0),
	}

	dbHospitals, err := u.DB.GetHospitals()
	if err != nil {
		return nil, err
	}

	for _, dbHospital := range dbHospitals {
		hospital := UAHospital{
			HospitalId: dbHospital.HospitalId,
			Name:       dbHospital.Name,
			Address:    dbHospital.Address,
			City:       dbHospital.City,
			PostalCode: dbHospital.PostalCode,
			Country:    dbHospital.Country,
		}

		hospitals.Hospitals = append(hospitals.Hospitals, &hospital)
	}

	return &hospitals, nil
}
