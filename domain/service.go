package domain

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/hikata101/climate_data_service/infrastructure"

	pb "github.com/hikata101/climate_data_service/gen/github.com/hikata101/climate_data_service/v1"
	openmeteo "github.com/innotechdevops/openmeteo"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func DownloadDataset(req *pb.DownloadDatasetRequest, stream grpc.ServerStreamingServer[pb.DownloadDatasetResponse]) error {
	// Implement the logic to download the dataset here.
	// For now, just print a message.
	switch req.Request.(type) {
	case *pb.DownloadDatasetRequest_OpenMeteo:
		// Handle OpenMeteo download requests
		tmp := req.GetOpenMeteo()
		request_parameters := openmeteo.Parameter{
			Longitude:      openmeteo.Float32(tmp.GetLongitude()),
			Latitude:       openmeteo.Float32(tmp.GetLatitude()),
			Hourly:         convertHourlyParameters(tmp.GetHourly()),
			Elevation:      openmeteo.Float32(tmp.GetElevation()),
			CurrentWeather: openmeteo.Bool(tmp.CurrentWeather),
		}
		slog.Debug(fmt.Sprintf("Downloading OpenMeteo dataset with parameters: %+v\n", request_parameters))
		// Call the external API to get the data
		resp, err := infrastructure.Execute(request_parameters)
		if err != nil {
			slog.Error(fmt.Sprintf("Error executing OpenMeteo request: %v", err))
			stream.Send(&pb.DownloadDatasetResponse{
				Status: int32(codes.Unknown),
			})
			return err
		}
		// Parse resp (JSON) into pb.OpenMeteoResponse using jsonpb
		var parsed pb.OpenMeteoDataset

		// if err := json.Unmarshal([]byte(resp), &parsed); err != nil {
		if err := jsonpb.UnmarshalString(resp, &parsed); err != nil {
			slog.Error(fmt.Sprintf("Error unmarshalling OpenMeteo response: %v", err))
			stream.Send(&pb.DownloadDatasetResponse{
				Status: int32(codes.Internal),
			})
			return err
		}
		// Send the parsed protobuf message back to the client
		slog.Debug("Successfully parsed OpenMeteo response into protobuf")
		stream.Send(&pb.DownloadDatasetResponse{
			Status: int32(codes.OK),
			Response: &pb.DownloadDatasetResponse_OpenMeteoDataset{
				OpenMeteoDataset: &parsed,
			},
		})
	default:
		slog.Error("Unknown download request type")
		return errors.New("unknown download request type")
	}
	return nil
}

func ListDatasets(req *pb.ListDatasetsRequest, stream grpc.ServerStreamingServer[pb.DownloadDatasetResponse]) error {
	wg := sync.WaitGroup{}
	switch req.Request.(type) {
	case *pb.ListDatasetsRequest_OpenMeteo:
		// Handle OpenMeteo download requests
		tmp := req.GetOpenMeteo()
		request_parameters := openmeteo.Parameter{
			Hourly:         convertHourlyParameters(tmp.GetHourly()),
			Elevation:      openmeteo.Float32(tmp.GetElevation()),
			CurrentWeather: openmeteo.Bool(tmp.CurrentWeather),
		}
		// Call the external API to get the data
		locations := tmp.GetLocations()
		slog.Debug(fmt.Sprintf("Listing OpenMeteo dataset with parameters: %+v\n", request_parameters))
		for _, loc := range locations {
			request_parameters.Latitude = openmeteo.Float32(loc.GetLatitude())
			request_parameters.Longitude = openmeteo.Float32(loc.GetLongitude())
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp, err := infrastructure.Execute(request_parameters)
				if err != nil {
					slog.Error(fmt.Sprintf("Error executing OpenMeteo request: %v", err))
					stream.Send(&pb.DownloadDatasetResponse{
						Status: int32(codes.Unknown),
					})
				}
				// Parse resp (JSON) into pb.OpenMeteoResponse using jsonpb
				var parsed pb.OpenMeteoDataset

				// if err := json.Unmarshal([]byte(resp), &parsed); err != nil {
				if err := jsonpb.UnmarshalString(resp, &parsed); err != nil {
					slog.Error(fmt.Sprintf("Error unmarshalling OpenMeteo response: %v", err))
					stream.Send(&pb.DownloadDatasetResponse{
						Status: int32(codes.Internal),
					})
				}

				// Send the parsed protobuf message back to the client
				slog.Debug("Successfully parsed OpenMeteo response into protobuf")
				stream.Send(&pb.DownloadDatasetResponse{
					Status: int32(codes.OK),
					Response: &pb.DownloadDatasetResponse_OpenMeteoDataset{
						OpenMeteoDataset: &parsed,
					},
				})
			}()
		}
	default:
		slog.Error("Unknown download request type")
		return errors.New("unknown download request type")
	}
	wg.Wait()
	slog.Debug("All dataset listings processed successfully")
	return nil
}
