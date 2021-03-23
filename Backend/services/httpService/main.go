package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/packages/protocol"
	"github.com/rareinator/Svendeprove/Backend/services/httpService/server"
	"google.golang.org/grpc"
)

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func execute() error {
	if err := godotenv.Load("../../.env"); err != nil {
		return err
	}

	srv := server.NewServer()

	journalConn, err := grpc.Dial(os.Getenv("JOURNAL_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer journalConn.Close()

	patientConn, err := grpc.Dial(os.Getenv("PATIENT_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer patientConn.Close()

	bookingConn, err := grpc.Dial(os.Getenv("BOOKING_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer bookingConn.Close()

	adminConn, err := grpc.Dial(os.Getenv("ADMIN_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer adminConn.Close()

	iotConn, err := grpc.Dial(os.Getenv("IOT_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer iotConn.Close()

	srv.JournalService = protocol.NewJournalServiceClient(journalConn)
	srv.PatientService = protocol.NewPatientServiceClient(patientConn)
	srv.BookingService = protocol.NewBookingServiceClient(bookingConn)
	srv.AdminService = protocol.NewAdminServiceClient(adminConn)
	srv.IotService = protocol.NewIotServiceClient(iotConn)

	srv.StaticFileDir = "./static"

	if err := srv.ServeHTTP(); err != nil {
		return err
	}

	return nil
}
