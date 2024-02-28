package main

import (
	"testing"

	"openprio_proxy/openprio_pt_position_data"

	"github.com/go-test/deep"
)

func TestConvertCSV(t *testing.T) {
	got, err := ConvertCSV("QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,2,1613652529565,0,1,")
	if err == nil {
		t.Error("Expected error.", got)
	}
}

func TestConvertCSV_2(t *testing.T) {
	got, err := ConvertCSV("QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,-1,1613652529565,0,0,0,,0")
	if err != nil {
		t.Error("Expected no error")
		return
	}
	expected := openprio_pt_position_data.LocationMessage{
		VehicleDescriptor: &openprio_pt_position_data.VehicleDescriptor{
			DataOwnerCode:           "QBUZZ",
			BlockCode:               40,
			VehicleNumber:           6015,
			NumberOfVehiclesCoupled: 0,
			DrivingDirection:        openprio_pt_position_data.DrivingDirection(openprio_pt_position_data.DrivingDirection_UNDEFINED),
		},
		Position: &openprio_pt_position_data.Position{
			Latitude:                   52.0,
			Longitude:                  5.0,
			Accuracy:                   4.0,
			Speed:                      20,
			Bearing:                    194.3,
			Odometer:                   13242,
			Hdop:                       4,
			NumberOfReceivedSatellites: -1,
		},
		Timestamp:         1613652529565,
		DoorStatus:        openprio_pt_position_data.DoorOpeningStatus(openprio_pt_position_data.DoorOpeningStatus_DOOR_OPENING_STATUS_NO_DATA),
		StopButtonStatus:  openprio_pt_position_data.StopButtonStatus(openprio_pt_position_data.StopButtonStatus_STOP_BUTTON_STATUS_NO_DATA),
		PassageStopStatus: openprio_pt_position_data.PassageStopStatus(openprio_pt_position_data.PassageStopStatus_PASSAGE_STOP_STATUS_UNKNOWN),
	}

	if diff := deep.Equal(got, expected); diff != nil {
		t.Error(diff)
	}
}

func TestConvertCSV_3(t *testing.T) {
	got, err := ConvertCSV("QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,6,1613652529565,1,0,1,B,1\n")
	if err != nil {
		t.Error("Expected no error")
	}
	expected := openprio_pt_position_data.LocationMessage{
		VehicleDescriptor: &openprio_pt_position_data.VehicleDescriptor{
			DataOwnerCode:           "QBUZZ",
			BlockCode:               40,
			VehicleNumber:           6015,
			DrivingDirection:        openprio_pt_position_data.DrivingDirection(openprio_pt_position_data.DrivingDirection_B_SIDE),
			NumberOfVehiclesCoupled: 1,
		},
		Position: &openprio_pt_position_data.Position{
			Latitude:                   52.0,
			Longitude:                  5.0,
			Accuracy:                   4.0,
			Speed:                      20,
			Bearing:                    194.3,
			Odometer:                   13242,
			Hdop:                       4,
			NumberOfReceivedSatellites: 6,
		},
		Timestamp:         1613652529565,
		DoorStatus:        openprio_pt_position_data.DoorOpeningStatus(openprio_pt_position_data.DoorOpeningStatus_DOOR_OPENING_STATUS_CLOSED),
		StopButtonStatus:  openprio_pt_position_data.StopButtonStatus(openprio_pt_position_data.StopButtonStatus_STOP_BUTTON_STATUS_NO_DATA),
		PassageStopStatus: openprio_pt_position_data.PassageStopStatus(openprio_pt_position_data.PassageStopStatus_PASSAGE_STOP_STATUS_FALSE),
	}

	if diff := deep.Equal(got, expected); diff != nil {
		t.Error(diff)
	}
}

func TestConvertCSV_4(t *testing.T) {
	got, err := ConvertCSV("QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,8,1613652529565,3,2,2,A,2\r\n")
	if err != nil {
		t.Error("Expected no error")
		return
	}
	expected := openprio_pt_position_data.LocationMessage{
		VehicleDescriptor: &openprio_pt_position_data.VehicleDescriptor{
			DataOwnerCode:           "QBUZZ",
			BlockCode:               40,
			VehicleNumber:           6015,
			NumberOfVehiclesCoupled: 2,
			DrivingDirection:        openprio_pt_position_data.DrivingDirection(openprio_pt_position_data.DrivingDirection_A_SIDE),
		},
		Position: &openprio_pt_position_data.Position{
			Latitude:                   52.0,
			Longitude:                  5.0,
			Accuracy:                   4.0,
			Speed:                      20,
			Bearing:                    194.3,
			Odometer:                   13242,
			Hdop:                       4,
			NumberOfReceivedSatellites: 8,
		},
		Timestamp:         1613652529565,
		DoorStatus:        openprio_pt_position_data.DoorOpeningStatus(openprio_pt_position_data.DoorOpeningStatus_DOOR_OPENING_STATUS_RELEASED),
		StopButtonStatus:  openprio_pt_position_data.StopButtonStatus(openprio_pt_position_data.StopButtonStatus_STOP_BUTTON_STATUS_ACTIVATED),
		PassageStopStatus: openprio_pt_position_data.PassageStopStatus(openprio_pt_position_data.PassageStopStatus_PASSAGE_STOP_STATUS_TRUE),
	}

	if diff := deep.Equal(got, expected); diff != nil {
		t.Error(diff)
	}
}
