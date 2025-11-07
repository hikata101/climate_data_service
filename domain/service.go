package domain

import (
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/hikata101/climate_data_service/infrastructure"
	"github.com/hikata101/climate_data_service/logger"

	pb "github.com/hikata101/climate_data_service/gen/github.com/hikata101/climate_data_service/v1"
	openmeteo "github.com/innotechdevops/openmeteo"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func DownloadDataset(req *pb.DownloadRequest, stream grpc.ServerStreamingServer[pb.DownloadReply]) error {
	// Implement the logic to download the dataset here.
	// For now, just print a message.
	switch req.Request.(type) {
	case *pb.DownloadRequest_OpenMeteo:
		// Handle OpenMeteo download requests
		tmp := req.GetOpenMeteo()
		request_parameters := openmeteo.Parameter{
			Longitude:      openmeteo.Float32(tmp.GetLongitude()),
			Latitude:       openmeteo.Float32(tmp.GetLatitude()),
			Hourly:         convertHourlyParameters(tmp.GetHourly()),
			Elevation:      openmeteo.Float32(tmp.GetElevation()),
			CurrentWeather: openmeteo.Bool(tmp.CurrentWeather),
		}
		logger.Logger.Debug(fmt.Sprintf("Downloading OpenMeteo dataset with parameters: %+v\n", request_parameters))
		// Call the external API to get the data
		resp, err := infrastructure.Execute(request_parameters)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("Error executing OpenMeteo request: %v", err))
			stream.Send(&pb.DownloadReply{
				Status: int32(codes.Unknown),
			})
			return err
		}
		// print(resp)
		// Parse resp (JSON) into pb.OpenMeteoReply using jsonpb
		// var parsed OpenMeteo_Response
		var parsed pb.DownloadReply_OpenMeteo

		// if err := json.Unmarshal([]byte(resp), &parsed); err != nil {
		if err := jsonpb.UnmarshalString(resp, &parsed); err != nil {
			logger.Logger.Error(fmt.Sprintf("Error unmarshalling OpenMeteo response: %v", err))
			stream.Send(&pb.DownloadReply{
				Status: int32(codes.Internal),
			})
			return err
		}
		// Send the parsed protobuf message back to the client
		logger.Logger.Debug("Successfully parsed OpenMeteo response into protobuf")
		stream.Send(&pb.DownloadReply{
			Status: int32(codes.OK),
			Reply: &pb.DownloadReply_OpenMeteoReply{
				OpenMeteoReply: &parsed,
			},
		})
	default:
		logger.Logger.Error("Unknown download request type")
		return errors.New("unknown download request type")
	}
	return nil
}
