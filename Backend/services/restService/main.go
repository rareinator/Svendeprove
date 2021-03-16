package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	"github.com/rareinator/Svendeprove/Backend/services/bookingService/booking"
	"github.com/rareinator/Svendeprove/Backend/services/iotService/iot"
	"github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
	"github.com/rareinator/Svendeprove/Backend/services/patientService/patient"
	"github.com/rareinator/Svendeprove/Backend/services/restService/server"
	"github.com/rareinator/Svendeprove/Backend/services/useradminService/useradmin"
	"github.com/tidwall/buntdb"
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

	srv := server.NewServer()

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

	var useradminConn *grpc.ClientConn
	useradminConn, err = grpc.Dial(os.Getenv("USERADMIN_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer useradminConn.Close()

	var iotConn *grpc.ClientConn
	iotConn, err = grpc.Dial(os.Getenv("IOT_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer iotConn.Close()

	localDB, err := buntdb.Open("data.db")
	if err != nil {
		return err
	}
	defer localDB.Close()

	srv.LocalDB = localDB

	srv.JournalService = journal.NewJournalServiceClient(journalConn)
	srv.AuthenticationService = authentication.NewAuthenticationServiceClient(authenticationConn)
	srv.PatientService = patient.NewPatientServiceClient(patientConn)
	srv.BookingService = booking.NewBookingServiceClient(bookingConn)
	srv.UseradminService = useradmin.NewUseradminServiceClient(useradminConn)
	srv.IotService = iot.NewIotServiceClient(iotConn)

	srv.StaticFileDir = "./static"

	srv.ServeHTTP()

	return nil
}
