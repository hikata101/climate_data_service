package handler

import (
	"log/slog"

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

func (s climateDataServer) Download(req *pb.DownloadDatasetRequest, stream grpc.ServerStreamingServer[pb.DownloadDatasetResponse]) error {
	slog.Info("Download request received")
	err := domain.DownloadDataset(req, stream)
	if err != nil {
		slog.Error("Error processing download request: " + err.Error())
		return err
	}
	slog.Info("Download request processed successfully")
	return nil
}

func (s climateDataServer) UploadDataset(stream grpc.ClientStreamingServer[pb.UploadChunk, pb.UploadResponse]) error {
	slog.Debug("Upload request received")
	slog.Debug("Upload functionality not implemented yet")
	return nil
}

func (s climateDataServer) ListDatasets(req *pb.ListDatasetsRequest, stream grpc.ServerStreamingServer[pb.DownloadDatasetResponse]) error {
	slog.Debug("List request received")
	err := domain.ListDatasets(req, stream)
	if err != nil {
		slog.Error("Error processing list request: " + err.Error())
		return err
	}
	slog.Debug("List functionality processed successfully")
	return nil
}
