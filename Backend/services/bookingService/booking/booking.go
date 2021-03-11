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
	return &BHealth{Message: fmt.Sprintf("Booing service is up and running on: %v 🚀", b.ListenAddress)}, nil
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
		PatientId:          booking.PatientId,
		ApprovedByEmployee: booking.ApprovedByEmployee,
	}

	if err := b.DB.CreateBooking(&dbBooking); err != nil {
		return nil, err
	}

	return booking, nil
}

func (b *BookingServer) GetBooking(ctx context.Context, br *BRequest) (*Booking, error) {
	dbBooking, err := b.DB.GetBooking(br.Id)
	if err != nil {
		return nil, err
	}

	result := Booking{
		BookingId:          dbBooking.BookingId,
		BookedTime:         dbBooking.Bookedtime.Format("02/01/2006 15:04:05"),
		BookedEnd:          dbBooking.BookedEnd.Format("02/01/2006 15:04:05"),
		PatientId:          dbBooking.PatientId,
		ApprovedByEmployee: dbBooking.ApprovedByEmployee,
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
		PatientId:          booking.PatientId,
		ApprovedByEmployee: booking.ApprovedByEmployee,
	}

	if err := b.DB.UpdateBooking(&dbBooking); err != nil {
		return nil, err
	}

	return booking, nil
}
