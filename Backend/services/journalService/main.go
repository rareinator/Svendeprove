package main

import (
	"fmt"
	"net"
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
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}

	js := journal.JournalServer{}

	grpcServer := grpc.NewServer()

	journal.RegisterJournalServiceServer(grpcServer, &js)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Faild to start gRPC server over port 9000: %v", err)
	}

	return nil
}
