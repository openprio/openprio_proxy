// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: openprio_pt_position_data.proto

package openprio_pt_position_data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This enum describes the status of the (unlocking of the) doors of the vehicle.
// DoorOpenStatus is OPEN when the doors are open. When the doors are locked the status is CLOSED. The DoorOpenStatus is RELEASED  when the passengers can open the doors but de the doors are physically closed.
type DoorOpeningStatus int32

const (
	DoorOpeningStatus_DOOR_OPENING_STATUS_NO_DATA  DoorOpeningStatus = 0
	DoorOpeningStatus_DOOR_OPENING_STATUS_CLOSED   DoorOpeningStatus = 1
	DoorOpeningStatus_DOOR_OPENING_STATUS_OPEN     DoorOpeningStatus = 2
	DoorOpeningStatus_DOOR_OPENING_STATUS_RELEASED DoorOpeningStatus = 3
)

// Enum value maps for DoorOpeningStatus.
var (
	DoorOpeningStatus_name = map[int32]string{
		0: "DOOR_OPENING_STATUS_NO_DATA",
		1: "DOOR_OPENING_STATUS_CLOSED",
		2: "DOOR_OPENING_STATUS_OPEN",
		3: "DOOR_OPENING_STATUS_RELEASED",
	}
	DoorOpeningStatus_value = map[string]int32{
		"DOOR_OPENING_STATUS_NO_DATA":  0,
		"DOOR_OPENING_STATUS_CLOSED":   1,
		"DOOR_OPENING_STATUS_OPEN":     2,
		"DOOR_OPENING_STATUS_RELEASED": 3,
	}
)

func (x DoorOpeningStatus) Enum() *DoorOpeningStatus {
	p := new(DoorOpeningStatus)
	*p = x
	return p
}

func (x DoorOpeningStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DoorOpeningStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_openprio_pt_position_data_proto_enumTypes[0].Descriptor()
}

func (DoorOpeningStatus) Type() protoreflect.EnumType {
	return &file_openprio_pt_position_data_proto_enumTypes[0]
}

func (x DoorOpeningStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DoorOpeningStatus.Descriptor instead.
func (DoorOpeningStatus) EnumDescriptor() ([]byte, []int) {
	return file_openprio_pt_position_data_proto_rawDescGZIP(), []int{0}
}

// This enum describes if the stop button is activated
// indicating that passengers want tot get out at it's next stop
type StopButtonStatus int32

const (
	StopButtonStatus_STOP_BUTTON_STATUS_NO_DATA       StopButtonStatus = 0
	StopButtonStatus_STOP_BUTTON_STATUS_NOT_ACTIVATED StopButtonStatus = 1
	StopButtonStatus_STOP_BUTTON_STATUS_ACTIVATED     StopButtonStatus = 2
)

// Enum value maps for StopButtonStatus.
var (
	StopButtonStatus_name = map[int32]string{
		0: "STOP_BUTTON_STATUS_NO_DATA",
		1: "STOP_BUTTON_STATUS_NOT_ACTIVATED",
		2: "STOP_BUTTON_STATUS_ACTIVATED",
	}
	StopButtonStatus_value = map[string]int32{
		"STOP_BUTTON_STATUS_NO_DATA":       0,
		"STOP_BUTTON_STATUS_NOT_ACTIVATED": 1,
		"STOP_BUTTON_STATUS_ACTIVATED":     2,
	}
)

func (x StopButtonStatus) Enum() *StopButtonStatus {
	p := new(StopButtonStatus)
	*p = x
	return p
}

func (x StopButtonStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StopButtonStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_openprio_pt_position_data_proto_enumTypes[1].Descriptor()
}

func (StopButtonStatus) Type() protoreflect.EnumType {
	return &file_openprio_pt_position_data_proto_enumTypes[1]
}

func (x StopButtonStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StopButtonStatus.Descriptor instead.
func (StopButtonStatus) EnumDescriptor() ([]byte, []int) {
	return file_openprio_pt_position_data_proto_rawDescGZIP(), []int{1}
}

// This enum describes the bus/tram will pass the stop without stopping.
// The status is used in the view area of the stop, where the driver can see that no passengers want to get on or off.
// or the driver wants to trigger the departure standing at a TimingStop with closed Doors.
// PASSAGE_STOP_STATUS_TRUE means that the driver triggered an action
// PASSAGE_STOP_STATUS_UNKNOWN means that the functionality is not implemented
// PASSAGE_STATUS_FALSE should be the default when this functionality is implemented.
type PassageStopStatus int32

const (
	PassageStopStatus_PASSAGE_STOP_STATUS_UNKNOWN PassageStopStatus = 0
	PassageStopStatus_PASSAGE_STOP_STATUS_FALSE   PassageStopStatus = 1
	PassageStopStatus_PASSAGE_STOP_STATUS_TRUE    PassageStopStatus = 2
)

// Enum value maps for PassageStopStatus.
var (
	PassageStopStatus_name = map[int32]string{
		0: "PASSAGE_STOP_STATUS_UNKNOWN",
		1: "PASSAGE_STOP_STATUS_FALSE",
		2: "PASSAGE_STOP_STATUS_TRUE",
	}
	PassageStopStatus_value = map[string]int32{
		"PASSAGE_STOP_STATUS_UNKNOWN": 0,
		"PASSAGE_STOP_STATUS_FALSE":   1,
		"PASSAGE_STOP_STATUS_TRUE":    2,
	}
)

func (x PassageStopStatus) Enum() *PassageStopStatus {
	p := new(PassageStopStatus)
	*p = x
	return p
}

func (x PassageStopStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PassageStopStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_openprio_pt_position_data_proto_enumTypes[2].Descriptor()
}

func (PassageStopStatus) Type() protoreflect.EnumType {
	return &file_openprio_pt_position_data_proto_enumTypes[2]
}

func (x PassageStopStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PassageStopStatus.Descriptor instead.
func (PassageStopStatus) EnumDescriptor() ([]byte, []int) {
	return file_openprio_pt_position_data_proto_rawDescGZIP(), []int{2}
}

// This enum describes which side of a rail vehicle is in front.
type DrivingDirection int32

const (
	DrivingDirection_UNDEFINED DrivingDirection = 0
	DrivingDirection_A_SIDE    DrivingDirection = 1
	DrivingDirection_B_SIDE    DrivingDirection = 2
)

// Enum value maps for DrivingDirection.
var (
	DrivingDirection_name = map[int32]string{
		0: "UNDEFINED",
		1: "A_SIDE",
		2: "B_SIDE",
	}
	DrivingDirection_value = map[string]int32{
		"UNDEFINED": 0,
		"A_SIDE":    1,
		"B_SIDE":    2,
	}
)

func (x DrivingDirection) Enum() *DrivingDirection {
	p := new(DrivingDirection)
	*p = x
	return p
}

func (x DrivingDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DrivingDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_openprio_pt_position_data_proto_enumTypes[3].Descriptor()
}

func (DrivingDirection) Type() protoreflect.EnumType {
	return &file_openprio_pt_position_data_proto_enumTypes[3]
}

func (x DrivingDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DrivingDirection.Descriptor instead.
func (DrivingDirection) EnumDescriptor() ([]byte, []int) {
	return file_openprio_pt_position_data_proto_rawDescGZIP(), []int{3}
}

// Message that vehicle sends every second.
type LocationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleDescriptor *VehicleDescriptor `protobuf:"bytes,1,opt,name=vehicle_descriptor,json=vehicleDescriptor,proto3" json:"vehicle_descriptor,omitempty"`
	Position          *Position          `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Timestamp         int64              `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // milliseconds since epoch
	DoorStatus        DoorOpeningStatus  `protobuf:"varint,4,opt,name=door_status,json=doorStatus,proto3,enum=DoorOpeningStatus" json:"door_status,omitempty"`
	StopButtonStatus  StopButtonStatus   `protobuf:"varint,5,opt,name=stop_button_status,json=stopButtonStatus,proto3,enum=StopButtonStatus" json:"stop_button_status,omitempty"`
	PassageStopStatus PassageStopStatus  `protobuf:"varint,6,opt,name=passage_stop_status,json=passageStopStatus,proto3,enum=PassageStopStatus" json:"passage_stop_status,omitempty"`
}

func (x *LocationMessage) Reset() {
	*x = LocationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openprio_pt_position_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationMessage) ProtoMessage() {}

func (x *LocationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_openprio_pt_position_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationMessage.ProtoReflect.Descriptor instead.
func (*LocationMessage) Descriptor() ([]byte, []int) {
	return file_openprio_pt_position_data_proto_rawDescGZIP(), []int{0}
}

func (x *LocationMessage) GetVehicleDescriptor() *VehicleDescriptor {
	if x != nil {
		return x.VehicleDescriptor
	}
	return nil
}

func (x *LocationMessage) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *LocationMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LocationMessage) GetDoorStatus() DoorOpeningStatus {
	if x != nil {
		return x.DoorStatus
	}
	return DoorOpeningStatus_DOOR_OPENING_STATUS_NO_DATA
}

func (x *LocationMessage) GetStopButtonStatus() StopButtonStatus {
	if x != nil {
		return x.StopButtonStatus
	}
	return StopButtonStatus_STOP_BUTTON_STATUS_NO_DATA
}

func (x *LocationMessage) GetPassageStopStatus() PassageStopStatus {
	if x != nil {
		return x.PassageStopStatus
	}
	return PassageStopStatus_PASSAGE_STOP_STATUS_UNKNOWN
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`   // WGS 84
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"` // WGS 84
	Accuracy  float32 `protobuf:"fixed32,3,opt,name=accuracy,proto3" json:"accuracy,omitempty"`   // in m
	Speed     float32 `protobuf:"fixed32,4,opt,name=speed,proto3" json:"speed,omitempty"`         // speed in m/s
	Bearing   float32 `protobuf:"fixed32,5,opt,name=bearing,proto3" json:"bearing,omitempty"`     // bearing in degrees
	// the odometer is always used relatively,
	// the odometer may be either the total distance a vehicle has traveled during its lifetime,
	// today or during the current trip (resetting more often is not allowed)
	Odometer int64   `protobuf:"varint,6,opt,name=odometer,proto3" json:"odometer,omitempty"` // in meters
	Hdop     float32 `protobuf:"fixed32,7,opt,name=hdop,proto3" json:"hdop,omitempty"`        // horizontal dilution of precision. https://en.m.wikipedia.org/wiki/Dilution_of_precision_(navigation)
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openprio_pt_position_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_openprio_pt_position_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_openprio_pt_position_data_proto_rawDescGZIP(), []int{1}
}

func (x *Position) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Position) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Position) GetAccuracy() float32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

func (x *Position) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Position) GetBearing() float32 {
	if x != nil {
		return x.Bearing
	}
	return 0
}

func (x *Position) GetOdometer() int64 {
	if x != nil {
		return x.Odometer
	}
	return 0
}

func (x *Position) GetHdop() float32 {
	if x != nil {
		return x.Hdop
	}
	return 0
}

// Fields to identify vehicle
type VehicleDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataOwnerCode string `protobuf:"bytes,1,opt,name=data_owner_code,json=dataOwnerCode,proto3" json:"data_owner_code,omitempty"` // required
	// One of block_code and vehicle_number is required
	BlockCode     int32 `protobuf:"varint,2,opt,name=block_code,json=blockCode,proto3" json:"block_code,omitempty"`
	VehicleNumber int32 `protobuf:"varint,3,opt,name=vehicle_number,json=vehicleNumber,proto3" json:"vehicle_number,omitempty"`
	// Optional, only relevant for vehicles that can run in two directions.
	DrivingDirection DrivingDirection `protobuf:"varint,5,opt,name=driving_direction,json=drivingDirection,proto3,enum=DrivingDirection" json:"driving_direction,omitempty"`
	// Optional, number of vehicles being coupled in total including the reporting vehicle.
	// For a single tram this is one, for two trams combined this value is two. When not reported (=0) 1 can be assumed.
	NumberOfVehiclesCoupled int32              `protobuf:"varint,6,opt,name=number_of_vehicles_coupled,json=numberOfVehiclesCoupled,proto3" json:"number_of_vehicles_coupled,omitempty"`
	JourneyDescriptor       *JourneyDescriptor `protobuf:"bytes,4,opt,name=journey_descriptor,json=journeyDescriptor,proto3" json:"journey_descriptor,omitempty"` // optional
}

func (x *VehicleDescriptor) Reset() {
	*x = VehicleDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openprio_pt_position_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleDescriptor) ProtoMessage() {}

func (x *VehicleDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_openprio_pt_position_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleDescriptor.ProtoReflect.Descriptor instead.
func (*VehicleDescriptor) Descriptor() ([]byte, []int) {
	return file_openprio_pt_position_data_proto_rawDescGZIP(), []int{2}
}

func (x *VehicleDescriptor) GetDataOwnerCode() string {
	if x != nil {
		return x.DataOwnerCode
	}
	return ""
}

func (x *VehicleDescriptor) GetBlockCode() int32 {
	if x != nil {
		return x.BlockCode
	}
	return 0
}

func (x *VehicleDescriptor) GetVehicleNumber() int32 {
	if x != nil {
		return x.VehicleNumber
	}
	return 0
}

func (x *VehicleDescriptor) GetDrivingDirection() DrivingDirection {
	if x != nil {
		return x.DrivingDirection
	}
	return DrivingDirection_UNDEFINED
}

func (x *VehicleDescriptor) GetNumberOfVehiclesCoupled() int32 {
	if x != nil {
		return x.NumberOfVehiclesCoupled
	}
	return 0
}

func (x *VehicleDescriptor) GetJourneyDescriptor() *JourneyDescriptor {
	if x != nil {
		return x.JourneyDescriptor
	}
	return nil
}

// JourneyDescriptor are optional parameters to describe a journey
// When JourneyDescriptor data is filled OpenPrio data can be combined with a static timetable without the use of
// other realtime sources like for example KV6 or SIRI-VM (those are used to link the vehicle_number / block_code with a journey_number).
// To fill in the JourneyDescriptor you need more inteligence in the system that creates the OpenPrio data, because the system
// should be aware what Journey the vehicle is running on.
// If you want to use journey_descriptor you need to fill in all the parameters within the JourneyDescriptor.
type JourneyDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JourneyNumber int32 `protobuf:"varint,1,opt,name=journey_number,json=journeyNumber,proto3" json:"journey_number,omitempty"`
	// Unique line number within a data_owner (operator).
	LinePlanningNumber string `protobuf:"bytes,2,opt,name=line_planning_number,json=linePlanningNumber,proto3" json:"line_planning_number,omitempty"`
	// The operating_day this journey is scheduled, be aware this is not necessarily the current date.
	OperatingDay string `protobuf:"bytes,3,opt,name=operating_day,json=operatingDay,proto3" json:"operating_day,omitempty"`
}

func (x *JourneyDescriptor) Reset() {
	*x = JourneyDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openprio_pt_position_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JourneyDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JourneyDescriptor) ProtoMessage() {}

func (x *JourneyDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_openprio_pt_position_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JourneyDescriptor.ProtoReflect.Descriptor instead.
func (*JourneyDescriptor) Descriptor() ([]byte, []int) {
	return file_openprio_pt_position_data_proto_rawDescGZIP(), []int{3}
}

func (x *JourneyDescriptor) GetJourneyNumber() int32 {
	if x != nil {
		return x.JourneyNumber
	}
	return 0
}

func (x *JourneyDescriptor) GetLinePlanningNumber() string {
	if x != nil {
		return x.LinePlanningNumber
	}
	return ""
}

func (x *JourneyDescriptor) GetOperatingDay() string {
	if x != nil {
		return x.OperatingDay
	}
	return ""
}

var File_openprio_pt_position_data_proto protoreflect.FileDescriptor

var file_openprio_pt_position_data_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x70, 0x65, 0x6e, 0x70, 0x72, 0x69, 0x6f, 0x5f, 0x70, 0x74, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x12, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x11, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x33, 0x0a,
	0x0b, 0x64, 0x6f, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x44, 0x6f, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x64, 0x6f, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3f, 0x0a, 0x12, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x62, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x10, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x6f,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x62, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x64,
	0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x64,
	0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x64, 0x6f, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x68, 0x64, 0x6f, 0x70, 0x22, 0xc1, 0x02, 0x0a, 0x11, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x12, 0x26, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3e,
	0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x72,
	0x69, 0x76, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x1a, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x70, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x17, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x70, 0x6c, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x12, 0x6a,
	0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x11, 0x6a, 0x6f, 0x75,
	0x72, 0x6e, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x22, 0x91,
	0x01, 0x0a, 0x11, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x65, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x69, 0x6e, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x79, 0x2a, 0x94, 0x01, 0x0a, 0x11, 0x44, 0x6f, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x4f, 0x4f, 0x52,
	0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x4f, 0x4f,
	0x52, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x4f, 0x4f,
	0x52, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x4f, 0x4f, 0x52, 0x5f,
	0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x7a, 0x0a, 0x10, 0x53, 0x74, 0x6f,
	0x70, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a,
	0x1a, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x00, 0x12, 0x24, 0x0a,
	0x20, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x42, 0x55, 0x54, 0x54,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x71, 0x0a, 0x11, 0x50, 0x61, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x41,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x50,
	0x41, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x41,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x10, 0x02, 0x2a, 0x39, 0x0a, 0x10, 0x44, 0x72, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x5f, 0x53, 0x49, 0x44, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x5f, 0x53, 0x49, 0x44,
	0x45, 0x10, 0x02, 0x42, 0x1c, 0x5a, 0x1a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x70, 0x72, 0x69, 0x6f,
	0x5f, 0x70, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_openprio_pt_position_data_proto_rawDescOnce sync.Once
	file_openprio_pt_position_data_proto_rawDescData = file_openprio_pt_position_data_proto_rawDesc
)

func file_openprio_pt_position_data_proto_rawDescGZIP() []byte {
	file_openprio_pt_position_data_proto_rawDescOnce.Do(func() {
		file_openprio_pt_position_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_openprio_pt_position_data_proto_rawDescData)
	})
	return file_openprio_pt_position_data_proto_rawDescData
}

var file_openprio_pt_position_data_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_openprio_pt_position_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_openprio_pt_position_data_proto_goTypes = []interface{}{
	(DoorOpeningStatus)(0),    // 0: DoorOpeningStatus
	(StopButtonStatus)(0),     // 1: StopButtonStatus
	(PassageStopStatus)(0),    // 2: PassageStopStatus
	(DrivingDirection)(0),     // 3: DrivingDirection
	(*LocationMessage)(nil),   // 4: LocationMessage
	(*Position)(nil),          // 5: Position
	(*VehicleDescriptor)(nil), // 6: VehicleDescriptor
	(*JourneyDescriptor)(nil), // 7: JourneyDescriptor
}
var file_openprio_pt_position_data_proto_depIdxs = []int32{
	6, // 0: LocationMessage.vehicle_descriptor:type_name -> VehicleDescriptor
	5, // 1: LocationMessage.position:type_name -> Position
	0, // 2: LocationMessage.door_status:type_name -> DoorOpeningStatus
	1, // 3: LocationMessage.stop_button_status:type_name -> StopButtonStatus
	2, // 4: LocationMessage.passage_stop_status:type_name -> PassageStopStatus
	3, // 5: VehicleDescriptor.driving_direction:type_name -> DrivingDirection
	7, // 6: VehicleDescriptor.journey_descriptor:type_name -> JourneyDescriptor
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_openprio_pt_position_data_proto_init() }
func file_openprio_pt_position_data_proto_init() {
	if File_openprio_pt_position_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_openprio_pt_position_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openprio_pt_position_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openprio_pt_position_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleDescriptor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openprio_pt_position_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JourneyDescriptor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openprio_pt_position_data_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_openprio_pt_position_data_proto_goTypes,
		DependencyIndexes: file_openprio_pt_position_data_proto_depIdxs,
		EnumInfos:         file_openprio_pt_position_data_proto_enumTypes,
		MessageInfos:      file_openprio_pt_position_data_proto_msgTypes,
	}.Build()
	File_openprio_pt_position_data_proto = out.File
	file_openprio_pt_position_data_proto_rawDesc = nil
	file_openprio_pt_position_data_proto_goTypes = nil
	file_openprio_pt_position_data_proto_depIdxs = nil
}
