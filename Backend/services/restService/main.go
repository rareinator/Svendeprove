package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	"github.com/rareinator/Svendeprove/Backend/services/journalService/journal"
	"google.golang.org/grpc"
)

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func execute() error {
	godotenv.Load("../.env")

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

	srv.journalService = journal.NewJournalServiceClient(journalConn)
	srv.authenticationService = authentication.NewAuthenticationServiceClient(authenticationConn)

	srv.ServeHTTP()

	return nil
}
