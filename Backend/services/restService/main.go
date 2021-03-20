package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/packages/protocol"
	"github.com/rareinator/Svendeprove/Backend/services/restService/server"
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

	var journalConn *grpc.ClientConn

	journalConn, err := grpc.Dial(os.Getenv("JOURNAL_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer journalConn.Close()

	var patientConn *grpc.ClientConn

	patientConn, err = grpc.Dial(os.Getenv("PATIENT_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer patientConn.Close()

	var bookingConn *grpc.ClientConn

	bookingConn, err = grpc.Dial(os.Getenv("BOOKING_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer bookingConn.Close()

	var useradminConn *grpc.ClientConn

	useradminConn, err = grpc.Dial(os.Getenv("USERADMIN_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer useradminConn.Close()

	var iotConn *grpc.ClientConn

	iotConn, err = grpc.Dial(os.Getenv("IOT_REMOTE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer iotConn.Close()

	srv.JournalService = protocol.NewJournalServiceClient(journalConn)
	srv.PatientService = protocol.NewPatientServiceClient(patientConn)
	srv.BookingService = protocol.NewBookingServiceClient(bookingConn)
	srv.UseradminService = protocol.NewUseradminServiceClient(useradminConn)
	srv.IotService = protocol.NewIotServiceClient(iotConn)

	srv.StaticFileDir = "./static"

	if err := srv.ServeHTTP(); err != nil {
		return err
	}

	return nil
}
