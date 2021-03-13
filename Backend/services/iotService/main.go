package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rareinator/Svendeprove/Backend/packages/mongo"
	"github.com/rareinator/Svendeprove/Backend/services/iotService/iot"
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

	fmt.Println("iot service listening on")
	fmt.Println(os.Getenv("IOT_SERVICE_ADDR"))

	lis, err := net.Listen("tcp", os.Getenv("IOT_SERVICE_ADDR"))
	if err != nil {
		return err
	}

	fmt.Println(os.Getenv("MONGO_URI"))

	mongo := mongo.MongoDB{
		Addr: os.Getenv("MONGO_URI"),
	}

	is := iot.IotServer{
		DB:            &mongo,
		ListenAddress: os.Getenv("IOT_SERVICE_ADDR"),
	}

	grpcServer := grpc.NewServer()

	iot.RegisterIotServiceServer(grpcServer, &is)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Faild to start gRPC server over addr: %v err: %v", os.Getenv("IOT_SERVICE_ADDR"), err)
	}

	return nil
}
