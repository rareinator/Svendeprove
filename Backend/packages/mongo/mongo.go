package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Addr string
}

func (m *MongoDB) newConnection(ctx context.Context) (*mongo.Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(m.Addr))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *MongoDB) UploadData(ctx context.Context, data *Device) error {
	client, err := m.newConnection(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	svendeProveDB := client.Database("Svendeprove")
	devicesCollection := svendeProveDB.Collection("Devices")
	fmt.Println("huh")

	_, err = devicesCollection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	fmt.Println("aha")

	return nil
}
