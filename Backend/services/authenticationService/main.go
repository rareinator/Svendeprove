package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/services/authenticationService/authentication"
	"google.golang.org/grpc"
)

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func execute() error {
	godotenv.Load()

	lis, err := net.Listen("tcp", os.Getenv("LISTEN_ADDR"))
	if err != nil {
		return err
	}

	as := authentication.AuthenticationServer{}

	grpcServer := grpc.NewServer()

	authentication.RegisterAuthenticationServiceServer(grpcServer, &as)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Failed to start gRPC server over port %v: %v", os.Getenv("LISTEN_ADDR"), err)
	}

	return nil

}
