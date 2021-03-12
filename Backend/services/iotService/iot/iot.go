package iot

import (
	"context"
	"fmt"
)

type IotServer struct {
	UnimplementedIotServiceServer
	ListenAddress string
}

func (i *IotServer) GetHealth(ctx context.Context, e *IOTEmpty) (*IOTHealth, error) {
	return &IOTHealth{Message: fmt.Sprintf("IOT service is up and running on: %v 🚀", i.ListenAddress)}, nil
}
