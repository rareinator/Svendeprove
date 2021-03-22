package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
	"github.com/rareinator/Svendeprove/Backend/packages/protocol"
	"github.com/rareinator/Svendeprove/Backend/services/adminService/admin"
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

	fmt.Println("Admin service listening on")
	fmt.Println(os.Getenv("ADMIN_SERVICE_ADDR"))

	lis, err := net.Listen("tcp", os.Getenv("ADMIN_SERVICE_ADDR"))
	if err != nil {
		return err
	}

	sql, err := mssql.NewConnection(os.Getenv("MSSQL_URI"))
	if err != nil {
		return err
	}
	ps := admin.AdminServer{
		DB:            sql,
		ListenAddress: os.Getenv("ADMIN_SERVICE_ADDR"),
	}

	grpcServer := grpc.NewServer()

	protocol.RegisterAdminServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("faild to start gRPC server over addr: %v err: %w", os.Getenv("MSSQL_URI"), err)
	}

	return nil
}
