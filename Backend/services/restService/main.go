package main

import (
	"fmt"
	"os"

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
	srv := newServer()

	var journalConn *grpc.ClientConn
	journalConn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer journalConn.Close()

	srv.journalService = journal.NewJournalServiceClient(journalConn)

	srv.ServeHTTP()

	return nil
}
