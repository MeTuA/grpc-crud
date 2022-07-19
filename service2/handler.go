package main

import (
	"context"
	datapb "service2/proto"
)

type dataService struct {
	datapb.UnimplementedDataServiceServer
	storage *Storage
}

func NewDataService(storage *Storage) *dataService {
	return &dataService{
		storage: storage,
	}
}

func (ds *dataService) DeleteData(ctx context.Context, in *datapb.Data) (*datapb.Data, error) {
	err := ds.storage.Delete(int(in.GetId()))
	if err != nil{
		return nil, err
	}

	return &datapb.Data{Id: in.GetId()}, nil
}

func (ds *dataService) UpdateData(ctx context.Context, in *datapb.Data) (*datapb.Data, error) {
	err := ds.storage.Update(in)
	if err != nil{
		return nil, err
	}

	return in, nil
}

func (ds *dataService) GetData(ctx context.Context, in *datapb.Data) (*datapb.Data, error) {
	data, err := ds.storage.Get(int(in.GetId()))
	if err != nil{
		return nil, err
	}

	return data, nil
}

func (ds *dataService) GetAllData(ctx context.Context, in *datapb.EmptyParams) (*datapb.ListData, error) {
	data, err := ds.storage.GetAll()
	if err != nil{
		return nil, err
	}

	return data, nil
}
