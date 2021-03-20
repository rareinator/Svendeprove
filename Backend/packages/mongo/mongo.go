package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

	svendeProveDB := client.Database("Svendeprove")
	devicesCollection := svendeProveDB.Collection("Devices")
	fmt.Println("huh")

	_, err = devicesCollection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	if err := client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) ReadData(ctx context.Context, deviceID int32) ([]*Device, error) {
	client, err := m.newConnection(ctx)
	if err != nil {
		return nil, err
	}

	svendeProveDB := client.Database("Svendeprove")
	devicesCollection := svendeProveDB.Collection("Devices")
	fmt.Println("huh")

	var results []*Device

	cur, err := devicesCollection.Find(ctx, bson.M{"sensorID": deviceID})
	if err != nil {
		return nil, err
	}

	fmt.Println("read data")

	for cur.Next(ctx) {
		fmt.Println("cursor.Next")
		var device Device
		if err := cur.Decode(&device); err != nil {
			return nil, err
		}

		results = append(results, &device)
	}

	cur.Close(ctx)

	if err := client.Disconnect(ctx); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *MongoDB) ReadDataInTimeFrame(ctx context.Context, timeStart, timeEnd time.Time) ([]*Device, error) {
	client, err := m.newConnection(ctx)
	if err != nil {
		return nil, err
	}

	svendeProveDB := client.Database("Svendeprove")
	devicesCollection := svendeProveDB.Collection("Devices")
	fmt.Println("huh")
	var results []*Device

	cur, err := devicesCollection.Find(ctx, bson.M{"date": bson.M{"$gte": timeStart, "$lte": timeEnd}})
	if err != nil {
		return nil, err
	}

	fmt.Println("read data")

	for cur.Next(ctx) {
		fmt.Println("cursor.Next")
		var device Device
		if err := cur.Decode(&device); err != nil {
			return nil, err
		}

		results = append(results, &device)
	}

	cur.Close(ctx)

	if err := client.Disconnect(ctx); err != nil {
		return nil, err
	}

	return results, nil
}
