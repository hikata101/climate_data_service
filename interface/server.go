package handler

import (
	"context"
	"fmt"

	"github.com/hikata101/climate_data_service/domain"
	pb "github.com/hikata101/climate_data_service/gen/github.com/hikata101/climate_data_service/v1"
	"google.golang.org/grpc"
)

type climateDataServer struct {
	pb.UnimplementedDatasetServer
}

func NewClimateDataServer() climateDataServer {
	return climateDataServer{}
}

func (s climateDataServer) Download(req *pb.DownloadRequest, stream grpc.ServerStreamingServer[pb.DownloadReply]) error {
	fmt.Println("Download request received")
	domain.DownloadDataset(req, stream)
	return nil
}

func (s climateDataServer) Upload(stream grpc.ClientStreamingServer[pb.UploadChunk, pb.UploadReply]) error {
	fmt.Println("Upload request received")
	return nil
}

func (s climateDataServer) List(ctx context.Context, req *pb.ListDatasetsRequest) (*pb.ListDatasetsReply, error) {
	fmt.Println("List request received")
	return nil, nil
}
