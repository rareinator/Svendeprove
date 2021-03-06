package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
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
	godotenv.Load("../../.env")

	fmt.Println("journal service listening on")
	fmt.Println(os.Getenv("JOURNAL_SERVICE_ADDR"))

	lis, err := net.Listen("tcp", os.Getenv("JOURNAL_SERVICE_ADDR"))
	if err != nil {
		return err
	}

	fmt.Println(os.Getenv("MSSQL_URI"))

	sql, err := mssql.NewConnection(os.Getenv("MSSQL_URI"))
	if err != nil {
		return err
	}
	js := journal.JournalServer{
		DB: sql,
	}

	grpcServer := grpc.NewServer()

	journal.RegisterJournalServiceServer(grpcServer, &js)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Faild to start gRPC server over addr: %v err: %v", os.Getenv("MSSQL_URI"), err)
	}

	return nil
}
