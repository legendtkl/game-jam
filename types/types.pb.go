// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	User
	Pixel
	Point
	GameMap
	CreateUserRequest
	CreateUserResponse
	CreateGameMapRequest
	CreateGameMapResponse
	Challenge
	ChallengeRequest
	ChallengeResponse
	ListGameMapRequest
	MapList
	ListGameMapResponse
	UploadChallengeResultRequest
*/
package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Balance  int64  `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type Pixel struct {
	Point *Point `protobuf:"bytes,1,opt,name=point" json:"point,omitempty"`
	Type  int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Value int32  `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Pixel) Reset()                    { *m = Pixel{} }
func (m *Pixel) String() string            { return proto.CompactTextString(m) }
func (*Pixel) ProtoMessage()               {}
func (*Pixel) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *Pixel) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Pixel) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Pixel) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Point struct {
	XAxis int32 `protobuf:"varint,1,opt,name=x_axis,json=xAxis,proto3" json:"x_axis,omitempty"`
	YAxis int32 `protobuf:"varint,2,opt,name=y_axis,json=yAxis,proto3" json:"y_axis,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *Point) GetXAxis() int32 {
	if m != nil {
		return m.XAxis
	}
	return 0
}

func (m *Point) GetYAxis() int32 {
	if m != nil {
		return m.YAxis
	}
	return 0
}

type GameMap struct {
	MapId   string   `protobuf:"bytes,1,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	Creator *User    `protobuf:"bytes,2,opt,name=creator" json:"creator,omitempty"`
	Graph   []*Pixel `protobuf:"bytes,3,rep,name=graph" json:"graph,omitempty"`
	Reward  int64    `protobuf:"varint,4,opt,name=reward,proto3" json:"reward,omitempty"`
	Fee     int64    `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
	State   int32    `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	Solver  *User    `protobuf:"bytes,7,opt,name=solver" json:"solver,omitempty"`
	Fogs    []*Point `protobuf:"bytes,8,rep,name=fogs" json:"fogs,omitempty"`
}

func (m *GameMap) Reset()                    { *m = GameMap{} }
func (m *GameMap) String() string            { return proto.CompactTextString(m) }
func (*GameMap) ProtoMessage()               {}
func (*GameMap) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *GameMap) GetMapId() string {
	if m != nil {
		return m.MapId
	}
	return ""
}

func (m *GameMap) GetCreator() *User {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *GameMap) GetGraph() []*Pixel {
	if m != nil {
		return m.Graph
	}
	return nil
}

func (m *GameMap) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *GameMap) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *GameMap) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *GameMap) GetSolver() *User {
	if m != nil {
		return m.Solver
	}
	return nil
}

func (m *GameMap) GetFogs() []*Point {
	if m != nil {
		return m.Fogs
	}
	return nil
}

type CreateUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{4} }

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CreateUserResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	User    *User  `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (m *CreateUserResponse) Reset()                    { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()               {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{5} }

func (m *CreateUserResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateUserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateGameMapRequest struct {
	Creator *User    `protobuf:"bytes,1,opt,name=creator" json:"creator,omitempty"`
	Graph   []*Pixel `protobuf:"bytes,2,rep,name=graph" json:"graph,omitempty"`
	Reward  int64    `protobuf:"varint,3,opt,name=reward,proto3" json:"reward,omitempty"`
	Fee     int64    `protobuf:"varint,4,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *CreateGameMapRequest) Reset()                    { *m = CreateGameMapRequest{} }
func (m *CreateGameMapRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateGameMapRequest) ProtoMessage()               {}
func (*CreateGameMapRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{6} }

func (m *CreateGameMapRequest) GetCreator() *User {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *CreateGameMapRequest) GetGraph() []*Pixel {
	if m != nil {
		return m.Graph
	}
	return nil
}

func (m *CreateGameMapRequest) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *CreateGameMapRequest) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type CreateGameMapResponse struct {
	Code    int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Map     *GameMap `protobuf:"bytes,3,opt,name=map" json:"map,omitempty"`
}

func (m *CreateGameMapResponse) Reset()                    { *m = CreateGameMapResponse{} }
func (m *CreateGameMapResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateGameMapResponse) ProtoMessage()               {}
func (*CreateGameMapResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{7} }

func (m *CreateGameMapResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateGameMapResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateGameMapResponse) GetMap() *GameMap {
	if m != nil {
		return m.Map
	}
	return nil
}

type Challenge struct {
	ChanllengeId string `protobuf:"bytes,1,opt,name=chanllenge_id,json=chanllengeId,proto3" json:"chanllenge_id,omitempty"`
	Player       *User  `protobuf:"bytes,2,opt,name=player" json:"player,omitempty"`
	MapId        string `protobuf:"bytes,3,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
}

func (m *Challenge) Reset()                    { *m = Challenge{} }
func (m *Challenge) String() string            { return proto.CompactTextString(m) }
func (*Challenge) ProtoMessage()               {}
func (*Challenge) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{8} }

func (m *Challenge) GetChanllengeId() string {
	if m != nil {
		return m.ChanllengeId
	}
	return ""
}

func (m *Challenge) GetPlayer() *User {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *Challenge) GetMapId() string {
	if m != nil {
		return m.MapId
	}
	return ""
}

type ChallengeRequest struct {
	Player *User  `protobuf:"bytes,1,opt,name=player" json:"player,omitempty"`
	MapId  string `protobuf:"bytes,2,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
}

func (m *ChallengeRequest) Reset()                    { *m = ChallengeRequest{} }
func (m *ChallengeRequest) String() string            { return proto.CompactTextString(m) }
func (*ChallengeRequest) ProtoMessage()               {}
func (*ChallengeRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{9} }

func (m *ChallengeRequest) GetPlayer() *User {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *ChallengeRequest) GetMapId() string {
	if m != nil {
		return m.MapId
	}
	return ""
}

type ChallengeResponse struct {
	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ChanllengeId string `protobuf:"bytes,3,opt,name=chanllenge_id,json=chanllengeId,proto3" json:"chanllenge_id,omitempty"`
}

func (m *ChallengeResponse) Reset()                    { *m = ChallengeResponse{} }
func (m *ChallengeResponse) String() string            { return proto.CompactTextString(m) }
func (*ChallengeResponse) ProtoMessage()               {}
func (*ChallengeResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{10} }

func (m *ChallengeResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ChallengeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ChallengeResponse) GetChanllengeId() string {
	if m != nil {
		return m.ChanllengeId
	}
	return ""
}

type ListGameMapRequest struct {
}

func (m *ListGameMapRequest) Reset()                    { *m = ListGameMapRequest{} }
func (m *ListGameMapRequest) String() string            { return proto.CompactTextString(m) }
func (*ListGameMapRequest) ProtoMessage()               {}
func (*ListGameMapRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{11} }

type MapList struct {
	Maps []*GameMap `protobuf:"bytes,1,rep,name=maps" json:"maps,omitempty"`
}

func (m *MapList) Reset()                    { *m = MapList{} }
func (m *MapList) String() string            { return proto.CompactTextString(m) }
func (*MapList) ProtoMessage()               {}
func (*MapList) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{12} }

func (m *MapList) GetMaps() []*GameMap {
	if m != nil {
		return m.Maps
	}
	return nil
}

type ListGameMapResponse struct {
	Maps []*GameMap `protobuf:"bytes,1,rep,name=maps" json:"maps,omitempty"`
}

func (m *ListGameMapResponse) Reset()                    { *m = ListGameMapResponse{} }
func (m *ListGameMapResponse) String() string            { return proto.CompactTextString(m) }
func (*ListGameMapResponse) ProtoMessage()               {}
func (*ListGameMapResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{13} }

func (m *ListGameMapResponse) GetMaps() []*GameMap {
	if m != nil {
		return m.Maps
	}
	return nil
}

type UploadChallengeResultRequest struct {
	Result       bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Player       *User  `protobuf:"bytes,2,opt,name=player" json:"player,omitempty"`
	ChanllengeId string `protobuf:"bytes,3,opt,name=chanllenge_id,json=chanllengeId,proto3" json:"chanllenge_id,omitempty"`
	LastPoint    *Point `protobuf:"bytes,4,opt,name=last_point,json=lastPoint" json:"last_point,omitempty"`
}

func (m *UploadChallengeResultRequest) Reset()         { *m = UploadChallengeResultRequest{} }
func (m *UploadChallengeResultRequest) String() string { return proto.CompactTextString(m) }
func (*UploadChallengeResultRequest) ProtoMessage()    {}
func (*UploadChallengeResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTypes, []int{14}
}

func (m *UploadChallengeResultRequest) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *UploadChallengeResultRequest) GetPlayer() *User {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *UploadChallengeResultRequest) GetChanllengeId() string {
	if m != nil {
		return m.ChanllengeId
	}
	return ""
}

func (m *UploadChallengeResultRequest) GetLastPoint() *Point {
	if m != nil {
		return m.LastPoint
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*Pixel)(nil), "Pixel")
	proto.RegisterType((*Point)(nil), "Point")
	proto.RegisterType((*GameMap)(nil), "GameMap")
	proto.RegisterType((*CreateUserRequest)(nil), "CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "CreateUserResponse")
	proto.RegisterType((*CreateGameMapRequest)(nil), "CreateGameMapRequest")
	proto.RegisterType((*CreateGameMapResponse)(nil), "CreateGameMapResponse")
	proto.RegisterType((*Challenge)(nil), "Challenge")
	proto.RegisterType((*ChallengeRequest)(nil), "ChallengeRequest")
	proto.RegisterType((*ChallengeResponse)(nil), "ChallengeResponse")
	proto.RegisterType((*ListGameMapRequest)(nil), "ListGameMapRequest")
	proto.RegisterType((*MapList)(nil), "MapList")
	proto.RegisterType((*ListGameMapResponse)(nil), "ListGameMapResponse")
	proto.RegisterType((*UploadChallengeResultRequest)(nil), "UploadChallengeResultRequest")
}

func init() { proto.RegisterFile("types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x63, 0xaf, 0x93, 0x4c, 0x40, 0x6a, 0x97, 0x14, 0x99, 0x28, 0x88, 0x68, 0x11, 0xa2,
	0xa7, 0x20, 0xb5, 0xe2, 0xc6, 0x05, 0xf5, 0x00, 0x95, 0xa8, 0xa8, 0x2c, 0xf5, 0x88, 0xa2, 0x69,
	0x3c, 0x49, 0x2c, 0xf9, 0x63, 0xf1, 0x3a, 0x25, 0x39, 0xf1, 0x5b, 0xf8, 0x5b, 0xfc, 0x1a, 0xb4,
	0xbb, 0x76, 0xe2, 0xb4, 0xc8, 0xa0, 0xde, 0xf6, 0x3d, 0x67, 0xde, 0xce, 0x7b, 0x33, 0x59, 0x18,
	0x94, 0x5b, 0x49, 0x6a, 0x2a, 0x8b, 0xbc, 0xcc, 0xc5, 0x07, 0xf0, 0x6e, 0x14, 0x15, 0x7c, 0x04,
	0xbd, 0xb5, 0xa2, 0x22, 0xc3, 0x94, 0x02, 0x67, 0xe2, 0x9c, 0xf6, 0xc3, 0x1d, 0xe6, 0x01, 0x74,
	0x6f, 0x31, 0xc1, 0x6c, 0x4e, 0x41, 0x67, 0xe2, 0x9c, 0xba, 0x61, 0x0d, 0xc5, 0x57, 0x60, 0xd7,
	0xf1, 0x86, 0x12, 0x3e, 0x06, 0x26, 0xf3, 0x38, 0x2b, 0x4d, 0xed, 0xe0, 0xcc, 0x9f, 0x5e, 0x6b,
	0x14, 0x5a, 0x92, 0x73, 0xf0, 0xf4, 0x9d, 0xa6, 0x9a, 0x85, 0xe6, 0xcc, 0x87, 0xc0, 0xee, 0x30,
	0x59, 0x53, 0xe0, 0x1a, 0xd2, 0x02, 0xf1, 0x1e, 0x98, 0xa9, 0xe4, 0x27, 0xe0, 0x6f, 0x66, 0xb8,
	0x89, 0x95, 0x51, 0x64, 0x21, 0xdb, 0x7c, 0xdc, 0xc4, 0x4a, 0xd3, 0x5b, 0x4b, 0x5b, 0x2d, 0xb6,
	0xd5, 0xb4, 0xf8, 0xed, 0x40, 0xf7, 0x13, 0xa6, 0x74, 0x85, 0x52, 0xff, 0x24, 0x45, 0x39, 0x8b,
	0xa3, 0xca, 0x07, 0x4b, 0x51, 0x5e, 0x46, 0xfc, 0x15, 0x74, 0xe7, 0x05, 0x61, 0x99, 0x17, 0xa6,
	0x74, 0x70, 0xc6, 0xa6, 0xda, 0x78, 0x58, 0xb3, 0xda, 0xc2, 0xb2, 0x40, 0xb9, 0x0a, 0xdc, 0x89,
	0x6b, 0x2d, 0x68, 0x67, 0xa1, 0x25, 0xf9, 0x73, 0xf0, 0x0b, 0xfa, 0x81, 0x45, 0x14, 0x78, 0x26,
	0x82, 0x0a, 0xf1, 0x23, 0x70, 0x17, 0x44, 0x01, 0x33, 0xa4, 0x3e, 0x6a, 0x63, 0xaa, 0xc4, 0x92,
	0x02, 0xdf, 0x76, 0x68, 0x00, 0x7f, 0x09, 0xbe, 0xca, 0x93, 0x3b, 0x2a, 0x82, 0x6e, 0xf3, 0xf6,
	0x8a, 0xe4, 0x23, 0xf0, 0x16, 0xf9, 0x52, 0x05, 0xbd, 0xfa, 0x6e, 0x13, 0x9f, 0xe1, 0xc4, 0x3b,
	0x38, 0xbe, 0xd0, 0x3d, 0x92, 0xa9, 0xa0, 0xef, 0x6b, 0x52, 0x65, 0xdb, 0xbc, 0xc4, 0x37, 0xe0,
	0xcd, 0x02, 0x25, 0xf3, 0x4c, 0x91, 0x1e, 0xc2, 0x3c, 0x8f, 0xa8, 0xca, 0xd3, 0x9c, 0xf5, 0x64,
	0x53, 0x52, 0x0a, 0x97, 0x76, 0x36, 0xfd, 0xb0, 0x86, 0xfc, 0x05, 0x78, 0x5a, 0xcf, 0x4c, 0x67,
	0xd7, 0xad, 0xa1, 0xc4, 0x4f, 0x18, 0x5a, 0xf9, 0x2a, 0xf1, 0xba, 0xa5, 0x46, 0xc2, 0x4e, 0x7b,
	0xc2, 0x9d, 0xf6, 0x84, 0xdd, 0xbf, 0x25, 0xec, 0xed, 0x12, 0x16, 0x08, 0x27, 0xf7, 0x1a, 0x78,
	0x94, 0xc5, 0x11, 0xb8, 0x29, 0xca, 0xca, 0x61, 0x6f, 0x5a, 0x8b, 0x69, 0x52, 0x2c, 0xa0, 0x7f,
	0xb1, 0xc2, 0x24, 0xa1, 0x6c, 0x49, 0xfc, 0x35, 0x3c, 0x9d, 0xaf, 0x30, 0xb3, 0x68, 0xbf, 0x58,
	0x4f, 0xf6, 0xe4, 0x65, 0xa4, 0x07, 0x2c, 0x13, 0xdc, 0xd2, 0xbd, 0xf5, 0xaa, 0xc8, 0xc6, 0x56,
	0xba, 0x8d, 0xad, 0x14, 0x9f, 0xe1, 0x68, 0x77, 0x4f, 0x9d, 0xe3, 0x5e, 0xc9, 0x69, 0x57, 0xea,
	0x34, 0x95, 0x16, 0x70, 0xdc, 0x50, 0x7a, 0x54, 0x20, 0x0f, 0x7c, 0xba, 0x0f, 0x7d, 0x8a, 0x21,
	0xf0, 0x2f, 0xb1, 0x2a, 0x0f, 0x67, 0x2f, 0xde, 0x42, 0xf7, 0x0a, 0xa5, 0xfe, 0xc0, 0xc7, 0xe0,
	0xa5, 0x28, 0xf5, 0xff, 0xd6, 0x3d, 0xc8, 0xd5, 0xb0, 0xe2, 0x1c, 0x9e, 0x1d, 0x94, 0x57, 0x8d,
	0xb6, 0x17, 0xfd, 0x72, 0x60, 0x7c, 0x23, 0x93, 0x1c, 0xa3, 0xa6, 0xc5, 0x75, 0x52, 0xd6, 0x91,
	0x99, 0xdd, 0xd1, 0x84, 0x71, 0xda, 0x0b, 0x2b, 0xf4, 0xaf, 0xa1, 0xfc, 0x8f, 0x61, 0xfe, 0x06,
	0x20, 0x41, 0x55, 0xce, 0xec, 0xfb, 0xe6, 0x1d, 0xbc, 0x6f, 0x7d, 0xfd, 0xc5, 0x1c, 0x6f, 0x7d,
	0xf3, 0x9e, 0x9e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x74, 0x8e, 0xc9, 0x5e, 0x05, 0x00,
	0x00,
}
