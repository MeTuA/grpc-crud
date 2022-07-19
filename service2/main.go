package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	datapb "service2/proto"
	"time"
)

func main() {
	// connecting to DB
	dbConn, err := ConnectToDB()
	if err != nil {
		fmt.Println("cannot connect to db", err.Error())
		return
	}

	// migration to db
	err = MigrateDB()
	if err != nil {
		fmt.Println("cannot migrate to db", err.Error())
		return
	}

	// starting listener
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	// starting grpc server and adding service
	s := grpc.NewServer()

	dataStorage := NewStorage(dbConn)
	dataService := NewDataService(dataStorage)
	datapb.RegisterDataServiceServer(s, dataService)


	go func() {
		s.Serve(lis)
	}()

	// gracefull shutdown
	stop := make(chan os.Signal)
	go func() {
		signal.Notify(stop, os.Interrupt, os.Kill)
	}()
	<-stop

	_, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	fmt.Println("service2 shut down")
}
