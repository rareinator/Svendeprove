package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
	"github.com/rareinator/Svendeprove/Backend/packages/protocol"
	"github.com/rareinator/Svendeprove/Backend/services/useradminService/useradmin"
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

	fmt.Println("Useradmin service listening on")
	fmt.Println(os.Getenv("USERADMIN_SERVICE_ADDR"))

	lis, err := net.Listen("tcp", os.Getenv("USERADMIN_SERVICE_ADDR"))
	if err != nil {
		return err
	}

	fmt.Println(os.Getenv("MSSQL_URI"))

	sql, err := mssql.NewConnection(os.Getenv("MSSQL_URI"))
	if err != nil {
		return err
	}
	ps := useradmin.UseradminServer{
		DB:            &sql,
		ListenAddress: os.Getenv("USERADMIN_SERVICE_ADDR"),
	}

	grpcServer := grpc.NewServer()

	protocol.RegisterUseradminServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Faild to start gRPC server over addr: %v err: %v", os.Getenv("MSSQL_URI"), err)
	}

	return nil
}
