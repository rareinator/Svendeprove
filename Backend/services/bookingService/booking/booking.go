package booking

import (
	"context"
	"fmt"
	"time"

	"github.com/rareinator/Svendeprove/Backend/packages/models"
	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
	. "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

type BookingServer struct {
	UnimplementedBookingServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (b *BookingServer) GetHealth(ctx context.Context, e *Empty) (*Health, error) {
	return &Health{Message: fmt.Sprintf("Booking service is up and running on: %v 🚀", b.ListenAddress)}, nil
}

func (b *BookingServer) CreateBooking(ctx context.Context, booking *Booking) (*Booking, error) {
	bookedTime, err := time.Parse("02/01/2006 15:04:05", booking.BookedTime)
	if err != nil {
		return nil, err
	}

	bookedEnd, err := time.Parse("02/01/2006 15:04:05", booking.BookedEnd)
	if err != nil {
		return nil, err
	}

	dbBooking := mssql.DBBooking{
		BookingId:  booking.BookingId,
		Bookedtime: bookedTime,
		BookedEnd:  bookedEnd,
		Patient:    booking.Patient,
		Employee:   booking.Employee,
		Approved:   booking.Approved,
		HospitalId: booking.Hospital.HospitalId,
	}

	if err := b.DB.CreateBooking(&dbBooking); err != nil {
		return nil, err
	}

	booking.BookingId = dbBooking.BookingId

	switch booking.Type {
	case string(models.Examination):
		examination := mssql.DBExamination{
			Description: booking.Description,
			StartedTime: bookedTime,
			EndedTime:   bookedEnd,
			BookingId:   booking.BookingId,
		}

		if err := b.DB.CreateExamination(&examination); err != nil {
			return nil, err
		}

	case string(models.Hospitilization): //Hospitilization
		hospitilization := mssql.DBHospitilization{
			Description: booking.Description,
			StartedTime: bookedTime,
			EndedTime:   bookedEnd,
			BookingId:   booking.BookingId,
			BedId:       booking.Bed.BedId,
		}

		if err := b.DB.CreateHospitilization(&hospitilization); err != nil {
			return nil, err
		}
	}

	return booking, nil
}

func (b *BookingServer) GetBooking(ctx context.Context, br *Request) (*Booking, error) {
	dbBooking, err := b.DB.GetBooking(br.Id)
	if err != nil {
		return nil, err
	}

	var bookingType string
	var description string

	hospitilization, err := b.DB.GetHospitilizationByBookingId(dbBooking.BookingId)
	if err != nil {
		return nil, err
	}

	if true {
		examination, err := b.DB.GetExaminationByBookingId(dbBooking.BookingId)
		if err != nil {
			return nil, err
		}

		if examination == nil {
			return nil, fmt.Errorf("could not find either hospitilization or examination data")
		} else {
			bookingType = "0"
			description = examination.Description
		}
	} else {
		bookingType = "1"
		description = hospitilization.Description
	}

	result := Booking{
		BookingId:   dbBooking.BookingId,
		BookedTime:  dbBooking.Bookedtime.Format("02/01/2006 15:04:05"),
		BookedEnd:   dbBooking.BookedEnd.Format("02/01/2006 15:04:05"),
		Patient:     dbBooking.Patient,
		Employee:    dbBooking.Employee,
		Approved:    dbBooking.Approved,
		Type:        bookingType,
		Description: description,
		Hospital: &Hospital{
			HospitalId: dbBooking.Hospital.HospitalId,
			Name:       dbBooking.Hospital.Name,
			Address:    dbBooking.Hospital.Address,
			City:       dbBooking.Hospital.City,
			PostalCode: dbBooking.Hospital.PostalCode,
			Country:    dbBooking.Hospital.Country,
		},
	}

	return &result, nil
}

func (b *BookingServer) UpdateBooking(ctx context.Context, booking *Booking) (*Booking, error) {
	bookedTime, err := time.Parse("02/01/2006 15:04:05", booking.BookedTime)
	if err != nil {
		return nil, err
	}

	bookedEnd, err := time.Parse("02/01/2006 15:04:05", booking.BookedEnd)
	if err != nil {
		return nil, err
	}

	dbBooking := mssql.DBBooking{
		BookingId:  booking.BookingId,
		Bookedtime: bookedTime,
		BookedEnd:  bookedEnd,
		Employee:   booking.Employee,
		Patient:    booking.Patient,
		Approved:   booking.Approved,
		HospitalId: booking.Hospital.HospitalId,
	}

	if err := b.DB.UpdateBooking(&dbBooking); err != nil {
		return nil, err
	}

	switch booking.Type {
	case string(models.Examination):
		examination := mssql.DBExamination{
			Description: booking.Description,
			StartedTime: bookedTime,
			EndedTime:   bookedEnd,
			BookingId:   booking.BookingId,
		}

		if err := b.DB.UpdateExamination(&examination); err != nil {
			return nil, err
		}

	case string(models.Hospitilization): //Hospitilization
		hospitilization := mssql.DBHospitilization{
			Description: booking.Description,
			StartedTime: bookedTime,
			EndedTime:   bookedEnd,
			BookingId:   booking.BookingId,
			BedId:       booking.Bed.BedId,
		}

		if err := b.DB.UpdateHospitilization(&hospitilization); err != nil {
			return nil, err
		}
	}

	return booking, nil
}

func (b *BookingServer) DeleteBooking(ctx context.Context, br *Request) (*Status, error) {
	dbBooking := mssql.DBBooking{
		BookingId: br.Id,
	}

	if err := b.DB.DeleteHospitilization(&mssql.DBHospitilization{BookingId: br.Id}); err != nil {
		return nil, err
	}

	if err := b.DB.DeleteExamination(&mssql.DBExamination{BookingId: br.Id}); err != nil {
		return nil, err
	}

	if err := b.DB.DeleteBooking(&dbBooking); err != nil {
		return &Status{Success: false}, err
	}

	return &Status{Success: true}, nil
}

func (b *BookingServer) GetBookingsByPatient(ctx context.Context, br *Request) (*Bookings, error) {
	bookings := Bookings{
		Bookings: make([]*Booking, 0),
	}

	dbBookings, err := b.DB.GetBookingsByPatient(br.UserId)
	if err != nil {
		return nil, err
	}

	for _, dbBooking := range dbBookings {
		booking := Booking{
			BookingId:  dbBooking.BookingId,
			BookedTime: dbBooking.Bookedtime.Format("02/01/2006 15:04:05"),
			BookedEnd:  dbBooking.BookedEnd.Format("02/01/2006 15:04:05"),
			Patient:    dbBooking.Patient,
			Employee:   dbBooking.Employee,
			Approved:   dbBooking.Approved,
			Hospital: &Hospital{
				HospitalId: dbBooking.Hospital.HospitalId,
				Name:       dbBooking.Hospital.Name,
				Address:    dbBooking.Hospital.Address,
				City:       dbBooking.Hospital.City,
				PostalCode: dbBooking.Hospital.PostalCode,
				Country:    dbBooking.Hospital.Country,
			},
		}

		hospitilization, err := b.DB.GetHospitilizationByBookingId(dbBooking.BookingId)
		if err != nil {
			return nil, err
		}

		if hospitilization == nil {
			examination, err := b.DB.GetExaminationByBookingId(dbBooking.BookingId)
			if err != nil {
				return nil, err
			}

			if examination == nil {
				return nil, fmt.Errorf("could not find either hospitilization or examination data")
			} else {
				booking.Type = string(models.Examination)
				booking.Description = examination.Description
			}
		} else {
			booking.Type = string(models.Hospitilization)
			booking.Description = hospitilization.Description
			booking.Bed = &Bed{
				BedId:        hospitilization.Bed.BedId,
				Name:         hospitilization.Bed.Name,
				Departmentid: hospitilization.Bed.DepartmentId,
				IsAvailable:  hospitilization.Bed.IsAvailable,
			}
		}

		bookings.Bookings = append(bookings.Bookings, &booking)
	}

	return &bookings, nil
}

func (b *BookingServer) GetBookingsByEmployee(ctx context.Context, br *Request) (*Bookings, error) {
	bookings := Bookings{
		Bookings: make([]*Booking, 0),
	}

	dbBookings, err := b.DB.GetBookingsByEmployee(br.UserId)
	if err != nil {
		return nil, err
	}

	for _, dbBooking := range dbBookings {

		booking := Booking{
			BookingId:  dbBooking.BookingId,
			BookedTime: dbBooking.Bookedtime.Format("02/01/2006 15:04:05"),
			BookedEnd:  dbBooking.BookedEnd.Format("02/01/2006 15:04:05"),
			Patient:    dbBooking.Patient,
			Employee:   dbBooking.Employee,
			Approved:   dbBooking.Approved,
			Hospital: &Hospital{
				HospitalId: dbBooking.Hospital.HospitalId,
				Name:       dbBooking.Hospital.Name,
				Address:    dbBooking.Hospital.Address,
				City:       dbBooking.Hospital.City,
				PostalCode: dbBooking.Hospital.PostalCode,
				Country:    dbBooking.Hospital.Country,
			},
		}

		hospitilization, err := b.DB.GetHospitilizationByBookingId(dbBooking.BookingId)
		if err != nil {
			return nil, err
		}

		if hospitilization == nil {
			examination, err := b.DB.GetExaminationByBookingId(dbBooking.BookingId)
			if err != nil {
				return nil, err
			}

			if examination == nil {
				return nil, fmt.Errorf("could not find either hospitilization or examination data, for booking: %v", dbBooking.BookingId)
			} else {
				booking.Type = string(models.Examination)
				booking.Description = examination.Description
			}
		} else {
			booking.Type = string(models.Hospitilization)
			booking.Description = hospitilization.Description
			booking.Bed = &Bed{
				BedId:        hospitilization.Bed.BedId,
				Name:         hospitilization.Bed.Name,
				Departmentid: hospitilization.Bed.DepartmentId,
				IsAvailable:  hospitilization.Bed.IsAvailable,
			}
		}

		bookings.Bookings = append(bookings.Bookings, &booking)
	}

	return &bookings, nil
}

func (b *BookingServer) GetAvailableTimesForDoctor(ctx context.Context, request *TimeFrameRequest) (*Strings, error) {
	parsedTime, err := time.Parse("02/01/2006 15:04:05", request.Day)
	if err != nil {
		return nil, err
	}

	roundedTime := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 0, 0, 0, 0, parsedTime.Location())

	bookedTimes, err := b.DB.GetBookedTimesForDoctor(roundedTime, request.Doctor)
	if err != nil {
		return nil, err
	}

	times := Strings{
		Strings: []string{},
	}

	availableTimes := make([]time.Time, 8)
	availableTimes[0] = roundedTime.Add(time.Hour * 8)
	availableTimes[1] = roundedTime.Add(time.Hour * 9)
	availableTimes[2] = roundedTime.Add(time.Hour * 10)
	availableTimes[3] = roundedTime.Add(time.Hour * 11)
	availableTimes[4] = roundedTime.Add(time.Hour * 12)
	availableTimes[5] = roundedTime.Add(time.Hour * 13)
	availableTimes[6] = roundedTime.Add(time.Hour * 14)
	availableTimes[7] = roundedTime.Add(time.Hour * 15)

	for idx, availableTime := range availableTimes {
		add := true

		for _, bookedTime := range bookedTimes {
			roundedBookedTime := time.Date(bookedTime.Year(), bookedTime.Month(), bookedTime.Day(), bookedTime.Hour(), 0, 0, 0, bookedTime.Location())
			if availableTime == roundedBookedTime {
				add = false
			}
		}

		if add {
			times.Strings = append(times.Strings, availableTimes[idx].Format("02/01/2006 15:04:05"))
		}
	}

	return &times, nil
}
