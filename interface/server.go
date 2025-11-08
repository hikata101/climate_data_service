package handler

import (
	"github.com/hikata101/climate_data_service/domain"
	pb "github.com/hikata101/climate_data_service/gen/github.com/hikata101/climate_data_service/v1"
	"github.com/hikata101/climate_data_service/logger"
	"google.golang.org/grpc"
)

type climateDataServer struct {
	pb.UnimplementedDatasetServer
}

func NewClimateDataServer() climateDataServer {
	return climateDataServer{}
}

func (s climateDataServer) Download(req *pb.DownloadDatasetRequest, stream grpc.ServerStreamingServer[pb.DownloadDatasetResponse]) error {
	logger.Logger.Info("Download request received")
	err := domain.DownloadDataset(req, stream)
	if err != nil {
		logger.Logger.Error("Error processing download request: " + err.Error())
		return err
	}
	logger.Logger.Info("Download request processed successfully")
	return nil
}

func (s climateDataServer) UploadDataset(stream grpc.ClientStreamingServer[pb.UploadChunk, pb.UploadResponse]) error {
	logger.Logger.Debug("Upload request received")
	logger.Logger.Debug("Upload functionality not implemented yet")
	return nil
}

func (s climateDataServer) ListDatasets(req *pb.ListDatasetsRequest, stream grpc.ServerStreamingServer[pb.DownloadDatasetResponse]) error {
	logger.Logger.Debug("List request received")
	err := domain.ListDatasets(req, stream)
	if err != nil {
		logger.Logger.Error("Error processing list request: " + err.Error())
		return err
	}
	logger.Logger.Debug("List functionality processed successfully")
	return nil
}
