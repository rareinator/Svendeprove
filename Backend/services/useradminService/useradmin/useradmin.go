package useradmin

import (
	"context"
	"fmt"
	"time"

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
	//TODO: Clean???
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

func (u *UseradminServer) GetDepartments(ctx context.Context, e *UAEmpty) (*Departments, error) {
	departments := Departments{
		Departments: make([]*Department, 0),
	}

	dbDepartments, err := u.DB.GetDepartments()
	if err != nil {
		return nil, err
	}

	for _, dbDepartment := range dbDepartments {
		department := Department{
			Departmentid: dbDepartment.DepartmentId,
			Name:         dbDepartment.Name,
			Description:  dbDepartment.Description,
			HospitalId:   dbDepartment.HospitalId,
		}

		departments.Departments = append(departments.Departments, &department)
	}

	return &departments, nil
}

func (u *UseradminServer) GetBeds(ctx context.Context, e *UAEmpty) (*Beds, error) {
	beds := Beds{
		Beds: make([]*Bed, 0),
	}

	dbBeds, err := u.DB.GetBeds()
	if err != nil {
		return nil, err
	}

	for _, dbBed := range dbBeds {
		Bed := Bed{
			BedId:        dbBed.BedId,
			Name:         dbBed.Name,
			Departmentid: dbBed.DepartmentId,
			IsAvailable:  dbBed.IsAvailable,
		}

		beds.Beds = append(beds.Beds, &Bed)
	}

	return &beds, nil
}

func (u *UseradminServer) GetAvailableBeds(ctx context.Context, r *BedsRequest) (*Beds, error) {
	startDate, err := time.Parse("02/01/2006 15:04:05", r.BookedTime)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("02/01/2006 15:04:05", r.BookedEnd)
	if err != nil {
		return nil, err
	}

	roundedStartDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	roundedEndDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())

	beds := Beds{
		Beds: make([]*Bed, 0),
	}

	dbBeds, err := u.DB.GetAvailableBeds(roundedStartDate, roundedEndDate, r.HospitalId)
	if err != nil {
		return nil, err
	}

	for _, dbBed := range dbBeds {
		bed := Bed{
			BedId:        dbBed.BedId,
			Name:         dbBed.Name,
			Departmentid: dbBed.DepartmentId,
			IsAvailable:  dbBed.IsAvailable,
			Department: &Department{
				Departmentid: dbBed.Department.DepartmentId,
				Name:         dbBed.Department.Name,
				Description:  dbBed.Department.Description,
				HospitalId:   dbBed.Department.HospitalId,
			},
		}

		beds.Beds = append(beds.Beds, &bed)
	}

	return &beds, nil
}
