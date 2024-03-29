package main

import (
	"fmt"
	"strconv"
	"strings"

	"openprio_proxy/openprio_pt_position_data"

	"google.golang.org/protobuf/proto"
)

func processMessage(bytes []byte, dataFormat string) (openprio_pt_position_data.LocationMessage, error) {
	positionData := openprio_pt_position_data.LocationMessage{}
	var err error
	switch dataFormat {
	case "csv":
		positionData, err = ConvertCSV(string(bytes))
	case "proto":
		err = proto.Unmarshal(bytes, &positionData)
	}

	return positionData, err
}

// ConvertCSV converts a CSV openprio message into a proto message.
// format:
// data_ownercode, block_code, vehicle_number, latitude (wsg84), longitude (wsg84), accuracy (m), speed (m/s), bearing (degrees), odometer (m), hdop, number_of_received_satellites, timestamp ( milliseconds since epoch), dooropeningstatus (see .proto), stopbuttonstatus (see .proto), number_of_vehicles_coupled (see .proto), driving_direction (nothing, A or B) passage_stop_status (see .proto) \n
// QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,1613652529565,0,1,1,A,0
func ConvertCSV(data string) (openprio_pt_position_data.LocationMessage, error) {
	positionData := openprio_pt_position_data.LocationMessage{}
	data = strings.TrimRight(data, "\r\n")
	values := strings.Split(data, ",")
	if len(values) != 17 {
		return positionData, fmt.Errorf("Number of fields in openprio CSV messages incorrect, it should be 16 but is %d", len(values))
	}

	blockCode, _ := strconv.ParseInt(values[1], 10, 32)
	vehicleNumber, _ := strconv.ParseInt(values[2], 10, 32)
	drivingDirection := openprio_pt_position_data.DrivingDirection(openprio_pt_position_data.DrivingDirection_UNDEFINED)
	if values[15] == "A" {
		drivingDirection = openprio_pt_position_data.DrivingDirection(openprio_pt_position_data.DrivingDirection_A_SIDE)
	} else if values[15] == "B" {
		drivingDirection = openprio_pt_position_data.DrivingDirection(openprio_pt_position_data.DrivingDirection_B_SIDE)

	}
	numberOfVehiclesCoupled, _ := strconv.ParseInt(values[14], 10, 32)
	positionData.VehicleDescriptor = &openprio_pt_position_data.VehicleDescriptor{
		DataOwnerCode:           values[0],
		BlockCode:               int32(blockCode),
		VehicleNumber:           int32(vehicleNumber),
		DrivingDirection:        drivingDirection,
		NumberOfVehiclesCoupled: int32(numberOfVehiclesCoupled),
	}

	latitude, _ := strconv.ParseFloat(values[3], 32)
	longitude, _ := strconv.ParseFloat(values[4], 32)
	accuracy, _ := strconv.ParseFloat(values[5], 32)
	speed, _ := strconv.ParseFloat(values[6], 32)
	bearing, _ := strconv.ParseFloat(values[7], 32)
	odometer, _ := strconv.ParseInt(values[8], 10, 64)
	hdop, _ := strconv.ParseFloat(values[9], 32)
	numberOfReceivedSatellites, _ := strconv.ParseInt(values[10], 10, 32)
	positionData.Position = &openprio_pt_position_data.Position{
		Latitude:                   float32(latitude),
		Longitude:                  float32(longitude),
		Accuracy:                   float32(accuracy),
		Speed:                      float32(speed),
		Bearing:                    float32(bearing),
		Odometer:                   odometer,
		Hdop:                       float32(hdop),
		NumberOfReceivedSatellites: int32(numberOfReceivedSatellites),
	}

	positionData.Timestamp, _ = strconv.ParseInt(values[11], 10, 64)
	doorOpeningStatus, _ := strconv.ParseInt(values[12], 10, 64)
	positionData.DoorStatus = openprio_pt_position_data.DoorOpeningStatus(doorOpeningStatus)
	stopButtonStatus, _ := strconv.ParseInt(values[13], 10, 64)
	positionData.StopButtonStatus = openprio_pt_position_data.StopButtonStatus(stopButtonStatus)
	passageStopStatus, _ := strconv.ParseInt(values[16], 10, 64)
	positionData.PassageStopStatus = openprio_pt_position_data.PassageStopStatus(passageStopStatus)

	return positionData, nil
}
