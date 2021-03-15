package booking

import (
	"context"
	"fmt"
	"time"

	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
)

type BookingServer struct {
	UnimplementedBookingServiceServer
	DB            *mssql.MSSQL
	ListenAddress string
}

func (b *BookingServer) GetHealth(ctx context.Context, e *BEmpty) (*BHealth, error) {
	return &BHealth{Message: fmt.Sprintf("Booking service is up and running on: %v 🚀", b.ListenAddress)}, nil
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
		BookingId:          booking.BookingId,
		Bookedtime:         bookedTime,
		BookedEnd:          bookedEnd,
		Patient:            booking.Patient,
		ApprovedByEmployee: booking.ApprovedByEmployee,
	}

	if err := b.DB.CreateBooking(&dbBooking); err != nil {
		return nil, err
	}

	booking.BookingId = dbBooking.BookingId

	return booking, nil
}

func (b *BookingServer) GetBooking(ctx context.Context, br *BRequest) (*Booking, error) {
	dbBooking, err := b.DB.GetBooking(br.Id)
	if err != nil {
		return nil, err
	}

	var bookingType int32
	var description string

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
			return nil, fmt.Errorf("Could not find either hospitilization or examination data")
		} else {
			bookingType = 0
			description = examination.Description
		}
	} else {
		bookingType = 1
		description = hospitilization.Description

	}

	result := Booking{
		BookingId:          dbBooking.BookingId,
		BookedTime:         dbBooking.Bookedtime.Format("02/01/2006 15:04:05"),
		BookedEnd:          dbBooking.BookedEnd.Format("02/01/2006 15:04:05"),
		Patient:            dbBooking.Patient,
		ApprovedByEmployee: dbBooking.ApprovedByEmployee,
		Type:               bookingType,
		Description:        description,
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
		BookingId:          booking.BookingId,
		Bookedtime:         bookedTime,
		BookedEnd:          bookedEnd,
		Patient:            booking.Patient,
		ApprovedByEmployee: booking.ApprovedByEmployee,
	}

	if err := b.DB.UpdateBooking(&dbBooking); err != nil {
		return nil, err
	}

	return booking, nil
}

func (b *BookingServer) DeleteBooking(ctx context.Context, br *BRequest) (*BStatus, error) {
	dbBooking := mssql.DBBooking{
		BookingId: br.Id,
	}

	if err := b.DB.DeleteBooking(&dbBooking); err != nil {
		return &BStatus{Success: false}, err
	}

	return &BStatus{Success: true}, nil
}

func (b *BookingServer) GetBookingsByPatient(ctx context.Context, br *BRequest) (*Bookings, error) {
	bookings := Bookings{
		Bookings: make([]*Booking, 0),
	}

	dbBookings, err := b.DB.GetBookingsByPatient(br.Username)
	if err != nil {
		return nil, err
	}

	for _, dbBooking := range dbBookings {
		booking := Booking{
			BookingId:          dbBooking.BookingId,
			BookedTime:         dbBooking.Bookedtime.Format("02/01/2006 15:04:05"),
			BookedEnd:          dbBooking.BookedEnd.Format("02/01/2006 15:04:05"),
			Patient:            dbBooking.Patient,
			ApprovedByEmployee: dbBooking.ApprovedByEmployee,
			Employee:           dbBooking.Employee,
		}

		bookings.Bookings = append(bookings.Bookings, &booking)
	}

	return &bookings, nil
}
