package booking

import (
	"context"
	"fmt"

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
