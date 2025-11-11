package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"

	"github.com/hikata101/climate_data_service/config"
	pb "github.com/hikata101/climate_data_service/gen/github.com/hikata101/climate_data_service/v1"
	"github.com/hikata101/climate_data_service/logger"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	handler "github.com/hikata101/climate_data_service/interface"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = logger.InitializeLogger()
	if err != nil {
		panic(err)
	}
	slog.Info("Starting Climate Data Service...")
	config := config.Load()
	host := config.Host
	port, err := strconv.Atoi(config.Port)
	if err != nil {
		slog.Error("Invalid port number", "error", err)
		panic(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	slog.Info(fmt.Sprintf("gRPC server is starting on %s:%d...", host, port))

	s := grpc.NewServer()
	server := handler.NewClimateDataServer()
	pb.RegisterDatasetServer(s, &server)

	go func() {
		slog.Info(fmt.Sprintf("start gRPC server on %s:%d", host, port))
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("stopping gRPC server...")
	s.GracefulStop()
}
