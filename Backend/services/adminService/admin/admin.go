package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
	. "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

type AdminServer struct {
	UnimplementedAdminServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (u *AdminServer) GetHealth(ctx context.Context, e *Empty) (*Health, error) {
	return &Health{Message: fmt.Sprintf("Admin service is up and running on: %v 🚀", u.ListenAddress)}, nil
}

func (u *AdminServer) GetHospitals(ctx context.Context, e *Empty) (*Hospitals, error) {
	hospitals := Hospitals{
		Hospitals: make([]*Hospital, 0),
	}

	dbHospitals, err := u.DB.GetHospitals()
	if err != nil {
		return nil, err
	}

	for _, dbHospital := range dbHospitals {
		hospital := Hospital{
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

func (u *AdminServer) GetAvailableBeds(ctx context.Context, r *BedsRequest) (*Beds, error) {
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
