package iot

import (
	"context"
	"fmt"
	"time"

	"github.com/rareinator/Svendeprove/Backend/packages/mongo"
	. "github.com/rareinator/Svendeprove/Backend/packages/protocol"
)

type IotServer struct {
	UnimplementedIotServiceServer
	ListenAddress string
	DB            *mongo.MongoDB
}

func (i *IotServer) GetHealth(ctx context.Context, e *Empty) (*Health, error) {
	return &Health{Message: fmt.Sprintf("IOT service is up and running on: %v 🚀", i.ListenAddress)}, nil
}

func (i *IotServer) UploadData(ctx context.Context, input *IOTData) (*IOTData, error) {
	fmt.Println("Saving data to mongoDB")

	fmt.Printf("saving %v for %v for device %v\n\r", input.Data, input.SensorID, input.Name)

	data := mongo.Device{
		Name:     input.Name,
		SensorID: input.SensorID,
		Data:     input.Data,
		Date:     time.Now(),
	}

	err := i.DB.UploadData(context.Background(), &data)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (i *IotServer) ReadData(ctx context.Context, request *Request) (*IOTDatas, error) {
	datas, err := i.DB.ReadData(context.Background(), request.Id)
	if err != nil {
		return nil, err
	}

	response := IOTDatas{
		IOTDatas: make([]*IOTData, 0),
	}

	for _, data := range datas {
		iotData := IOTData{
			ID:       data.ID.String(),
			Name:     data.Name,
			SensorID: data.SensorID,
			Data:     data.Data,
			Date:     data.Date.Format("02/01/2006 15:04:05"),
		}

		response.IOTDatas = append(response.IOTDatas, &iotData)
	}

	return &response, nil
}

func (i *IotServer) ReadDataInTimeFrame(ctx context.Context, request *IOTTimeframeRequest) (*IOTDatas, error) {
	parsedStartTime, err := time.Parse("02/01/2006 15:04:05", request.TimeStart)
	if err != nil {
		return nil, err
	}

	parsedEndTime, err := time.Parse("02/01/2006 15:04:05", request.TimeEnd)
	if err != nil {
		return nil, err
	}

	datas, err := i.DB.ReadDataInTimeFrame(context.Background(), parsedStartTime, parsedEndTime)
	if err != nil {
		return nil, err
	}

	response := IOTDatas{
		IOTDatas: make([]*IOTData, 0),
	}

	for _, data := range datas {
		iotData := IOTData{
			ID:       data.ID.String(),
			Name:     data.Name,
			SensorID: data.SensorID,
			Data:     data.Data,
			Date:     data.Date.Format("02/01/2006 15:04:05"),
		}

		response.IOTDatas = append(response.IOTDatas, &iotData)
	}

	return &response, nil
}
