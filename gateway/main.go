package main

import (
	"context"
	"fmt"
	datapb "gateway/proto"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	// connection to DataService
	conn, err := grpc.Dial("service2:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connected")
	}
	fmt.Println("connected to DataService grpc server")

	defer conn.Close()

	// connection to DataLoader
	connDataLoader, err := grpc.Dial("service1:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connected")
	}
	defer connDataLoader.Close()
	fmt.Println("connected to DataLoader grpc server")



	// setting clients
	dataServiceClient := datapb.NewDataServiceClient(conn)
	dataLoaderClient := datapb.NewDataLoaderClient(connDataLoader)

	// starting gateway server
	app := NewHandler(dataServiceClient, dataLoaderClient)
	app.Start(":8085")
	fmt.Println("server started")

	// gracefull shutdown
	stop := make(chan os.Signal)
	go func() {
		signal.Notify(stop, os.Interrupt, os.Kill)
	}()
	<-stop

	_, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	fmt.Println("gateway stopped")
}
