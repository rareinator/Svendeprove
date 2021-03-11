package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	goldap "github.com/rareinator/Svendeprove/Backend/packages/ldap"
	"github.com/rareinator/Svendeprove/Backend/packages/mssql"
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
	godotenv.Load("../../.env")

	lis, err := net.Listen("tcp", os.Getenv("AUTHENTICATION_SERVICE_ADDR"))
	if err != nil {
		return err
	}

	ldap := goldap.LDAP{
		Uri:           os.Getenv("LDAP_URI"),
		AdminUsername: os.Getenv("LDAP_READONLY_USER"),
		AdminPassword: os.Getenv("LDAP_READONLY_USER_PASSWORD"),
	}

	sql, err := mssql.NewConnection(os.Getenv("MSSQL_URI"))
	if err != nil {
		return err
	}

	as := authentication.AuthenticationServer{}
	as.ListenAddress = os.Getenv("AUTHENTICATION_SERVICE_ADDR")
	as.Ldap = &ldap
	as.DB = &sql

	grpcServer := grpc.NewServer()

	authentication.RegisterAuthenticationServiceServer(grpcServer, &as)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Failed to start gRPC server over port %v: %v", os.Getenv("AUTHENTICATION_SERVICE_ADDR"), err)
	}

	return nil

}
