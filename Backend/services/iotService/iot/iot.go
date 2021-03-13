package iot

import (
	"context"
	"fmt"

	"github.com/rareinator/Svendeprove/Backend/packages/mongo"
)

type IotServer struct {
	UnimplementedIotServiceServer
	ListenAddress string
	DB            *mongo.MongoDB
}

func (i *IotServer) GetHealth(ctx context.Context, e *IOTEmpty) (*IOTHealth, error) {
	return &IOTHealth{Message: fmt.Sprintf("IOT service is up and running on: %v 🚀", i.ListenAddress)}, nil
}

func (i *IotServer) UploadData(ctx context.Context, input *IOTData) (*IOTData, error) {
	fmt.Println("Saving data to mongoDB")

	fmt.Printf("saving %v for %v for device %v\n\r", input.Data, input.SensorID, input.Name)

	data := mongo.Device{
		Name:      input.Name,
		SensorID:  input.SensorID,
		Data:      input.Data,
		Timestamp: input.Timestamp,
	}

	err := i.DB.UploadData(context.Background(), &data)
	if err != nil {
		return nil, err
	}

	return input, nil

}

func (i *IotServer) ReadData(ctx context.Context, request *IOTRequest) (*IOTDatas, error) {
	datas, err := i.DB.ReadData(context.Background(), request.ID)
	if err != nil {
		return nil, err
	}

	fmt.Println("Read IOT data")

	response := IOTDatas{
		IOTDatas: make([]*IOTData, 0),
	}

	for _, data := range datas {
		iotData := IOTData{
			ID:        data.ID.String(),
			Name:      data.Name,
			SensorID:  data.SensorID,
			Data:      data.Data,
			Timestamp: data.Timestamp,
		}

		response.IOTDatas = append(response.IOTDatas, &iotData)
	}

	return &response, nil

}
