package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"openprio_proxy/openprio_pt_position_data"

	"github.com/golang/protobuf/proto"
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

	log.Print(positionData)
	return positionData, err
}

// ConvertCSV converts a CSV openprio message into a proto message.
// format:
// data_ownercode, block_code, vehicle_number, latitude, longitude, accuracy, speed, bearing, odometer, hdop, timestamp, dooropeningstatus, stopbuttonstatus\n
// QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,1613652529565,0,1
func ConvertCSV(data string) (openprio_pt_position_data.LocationMessage, error) {
	positionData := openprio_pt_position_data.LocationMessage{}
	data = strings.TrimRight(data, "\r\n")
	values := strings.Split(data, ",")
	if len(values) != 13 {
		return positionData, fmt.Errorf("Number of fields in openprio CSV messages incorrect, it should be 13 but is %d", len(values))
	}

	blockCode, _ := strconv.ParseInt(values[1], 10, 32)
	vehicleNumber, _ := strconv.ParseInt(values[2], 10, 32)
	positionData.VehicleDescriptor = &openprio_pt_position_data.VehicleDescriptor{
		DataOwnerCode: values[0],
		BlockCode:     int32(blockCode),
		VehicleNumber: int32(vehicleNumber),
	}

	latitude, _ := strconv.ParseFloat(values[3], 32)
	longitude, _ := strconv.ParseFloat(values[4], 32)
	accuracy, _ := strconv.ParseFloat(values[5], 32)
	speed, _ := strconv.ParseFloat(values[6], 32)
	bearing, _ := strconv.ParseFloat(values[7], 32)
	odometer, _ := strconv.ParseInt(values[8], 10, 64)
	hdop, _ := strconv.ParseFloat(values[9], 32)
	positionData.Position = &openprio_pt_position_data.Position{
		Latitude:  float32(latitude),
		Longitude: float32(longitude),
		Accuracy:  float32(accuracy),
		Speed:     float32(speed),
		Bearing:   float32(bearing),
		Odometer:  odometer,
		Hdop:      float32(hdop),
	}

	positionData.Timestamp, _ = strconv.ParseInt(values[10], 10, 64)
	doorOpeningStatus, _ := strconv.ParseInt(values[11], 10, 64)
	positionData.DoorStatus = openprio_pt_position_data.DoorOpeningStatus(doorOpeningStatus)
	stopButtonStatus, _ := strconv.ParseInt(values[12], 10, 64)
	positionData.StopButtonStatus = openprio_pt_position_data.StopButtonStatus(stopButtonStatus)

	return positionData, nil
}
