package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	"github.com/rareinator/Svendeprove/Backend/services/bookingService/booking"
	"github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
	"github.com/rareinator/Svendeprove/Backend/services/patientService/patient"
	"google.golang.org/grpc"
)

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func execute() error {
	godotenv.Load("../../.env")

	srv := newServer()

	var journalConn *grpc.ClientConn
	journalConn, err := grpc.Dial(os.Getenv("JOURNAL_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer journalConn.Close()

	var authenticationConn *grpc.ClientConn
	authenticationConn, err = grpc.Dial(os.Getenv("AUTHENTICATION_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer authenticationConn.Close()

	var patientConn *grpc.ClientConn
	patientConn, err = grpc.Dial(os.Getenv("PATIENT_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer patientConn.Close()

	var bookingConn *grpc.ClientConn
	bookingConn, err = grpc.Dial(os.Getenv("BOOKING_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer bookingConn.Close()

	srv.journalService = journal.NewJournalServiceClient(journalConn)
	srv.authenticationService = authentication.NewAuthenticationServiceClient(authenticationConn)
	srv.patientService = patient.NewPatientServiceClient(patientConn)
	srv.bookingService = booking.NewBookingServiceClient(bookingConn)

	srv.ServeHTTP()

	return nil
}
