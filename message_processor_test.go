package main

import (
	"testing"

	"openprio_proxy/openprio_pt_position_data"

	"github.com/go-test/deep"
)

func TestConvertCSV(t *testing.T) {
	got, err := ConvertCSV("QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,1613652529565,0,1,")
	if err == nil {
		t.Error("Expected error.", got)
	}
}

func TestConvertCSV_2(t *testing.T) {
	got, err := ConvertCSV("QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,1613652529565,1,2")
	if err != nil {
		t.Error("Expected no error")
	}
	expected := openprio_pt_position_data.LocationMessage{
		VehicleDescriptor: &openprio_pt_position_data.VehicleDescriptor{
			DataOwnerCode: "QBUZZ",
			BlockCode:     40,
			VehicleNumber: 6015,
		},
		Position: &openprio_pt_position_data.Position{
			Latitude:  52.0,
			Longitude: 5.0,
			Accuracy:  4.0,
			Speed:     20,
			Bearing:   194.3,
			Odometer:  13242,
			Hdop:      4,
		},
		Timestamp:        1613652529565,
		DoorStatus:       openprio_pt_position_data.DoorOpeningStatus(openprio_pt_position_data.DoorOpeningStatus_DOOR_OPENING_STATUS_OPEN),
		StopButtonStatus: openprio_pt_position_data.StopButtonStatus(openprio_pt_position_data.StopButtonStatus_STOP_BUTTON_STATUS_NO_DATA),
	}

	if diff := deep.Equal(got, expected); diff != nil {
		t.Error(diff)
	}
}

func TestConvertCSV_3(t *testing.T) {
	got, err := ConvertCSV("QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,1613652529565,1,2\n")
	if err != nil {
		t.Error("Expected no error")
	}
	expected := openprio_pt_position_data.LocationMessage{
		VehicleDescriptor: &openprio_pt_position_data.VehicleDescriptor{
			DataOwnerCode: "QBUZZ",
			BlockCode:     40,
			VehicleNumber: 6015,
		},
		Position: &openprio_pt_position_data.Position{
			Latitude:  52.0,
			Longitude: 5.0,
			Accuracy:  4.0,
			Speed:     20,
			Bearing:   194.3,
			Odometer:  13242,
			Hdop:      4,
		},
		Timestamp:        1613652529565,
		DoorStatus:       openprio_pt_position_data.DoorOpeningStatus(openprio_pt_position_data.DoorOpeningStatus_DOOR_OPENING_STATUS_OPEN),
		StopButtonStatus: openprio_pt_position_data.StopButtonStatus(openprio_pt_position_data.StopButtonStatus_STOP_BUTTON_STATUS_NO_DATA),
	}

	if diff := deep.Equal(got, expected); diff != nil {
		t.Error(diff)
	}
}

func TestConvertCSV_4(t *testing.T) {
	got, err := ConvertCSV("QBUZZ,40,6015,52.0,5.0,4.0,20,194.3,13242,4,1613652529565,1,2\r\n")
	if err != nil {
		t.Error("Expected no error")
	}
	expected := openprio_pt_position_data.LocationMessage{
		VehicleDescriptor: &openprio_pt_position_data.VehicleDescriptor{
			DataOwnerCode: "QBUZZ",
			BlockCode:     40,
			VehicleNumber: 6015,
		},
		Position: &openprio_pt_position_data.Position{
			Latitude:  52.0,
			Longitude: 5.0,
			Accuracy:  4.0,
			Speed:     20,
			Bearing:   194.3,
			Odometer:  13242,
			Hdop:      4,
		},
		Timestamp:        1613652529565,
		DoorStatus:       openprio_pt_position_data.DoorOpeningStatus(openprio_pt_position_data.DoorOpeningStatus_DOOR_OPENING_STATUS_OPEN),
		StopButtonStatus: openprio_pt_position_data.StopButtonStatus(openprio_pt_position_data.StopButtonStatus_STOP_BUTTON_STATUS_NO_DATA),
	}

	if diff := deep.Equal(got, expected); diff != nil {
		t.Error(diff)
	}
}
