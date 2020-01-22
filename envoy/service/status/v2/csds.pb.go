// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/status/v2/csds.proto

package envoy_service_status_v2

import (
	context "context"
	fmt "fmt"
	v2alpha "github.com/envoyproxy/go-control-plane/envoy/admin/v2alpha"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConfigStatus int32

const (
	ConfigStatus_UNKNOWN  ConfigStatus = 0
	ConfigStatus_SYNCED   ConfigStatus = 1
	ConfigStatus_NOT_SENT ConfigStatus = 2
	ConfigStatus_STALE    ConfigStatus = 3
	ConfigStatus_ERROR    ConfigStatus = 4
)

var ConfigStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SYNCED",
	2: "NOT_SENT",
	3: "STALE",
	4: "ERROR",
}

var ConfigStatus_value = map[string]int32{
	"UNKNOWN":  0,
	"SYNCED":   1,
	"NOT_SENT": 2,
	"STALE":    3,
	"ERROR":    4,
}

func (x ConfigStatus) String() string {
	return proto.EnumName(ConfigStatus_name, int32(x))
}

func (ConfigStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8273221d61b8a20, []int{0}
}

type ClientStatusRequest struct {
	NodeMatchers         []*matcher.NodeMatcher `protobuf:"bytes,1,rep,name=node_matchers,json=nodeMatchers,proto3" json:"node_matchers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ClientStatusRequest) Reset()         { *m = ClientStatusRequest{} }
func (m *ClientStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ClientStatusRequest) ProtoMessage()    {}
func (*ClientStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8273221d61b8a20, []int{0}
}

func (m *ClientStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatusRequest.Unmarshal(m, b)
}
func (m *ClientStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatusRequest.Marshal(b, m, deterministic)
}
func (m *ClientStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatusRequest.Merge(m, src)
}
func (m *ClientStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ClientStatusRequest.Size(m)
}
func (m *ClientStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatusRequest proto.InternalMessageInfo

func (m *ClientStatusRequest) GetNodeMatchers() []*matcher.NodeMatcher {
	if m != nil {
		return m.NodeMatchers
	}
	return nil
}

type PerXdsConfig struct {
	Status ConfigStatus `protobuf:"varint,1,opt,name=status,proto3,enum=envoy.service.status.v2.ConfigStatus" json:"status,omitempty"`
	// Types that are valid to be assigned to PerXdsConfig:
	//	*PerXdsConfig_ListenerConfig
	//	*PerXdsConfig_ClusterConfig
	//	*PerXdsConfig_RouteConfig
	//	*PerXdsConfig_ScopedRouteConfig
	PerXdsConfig         isPerXdsConfig_PerXdsConfig `protobuf_oneof:"per_xds_config"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PerXdsConfig) Reset()         { *m = PerXdsConfig{} }
func (m *PerXdsConfig) String() string { return proto.CompactTextString(m) }
func (*PerXdsConfig) ProtoMessage()    {}
func (*PerXdsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8273221d61b8a20, []int{1}
}

func (m *PerXdsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerXdsConfig.Unmarshal(m, b)
}
func (m *PerXdsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerXdsConfig.Marshal(b, m, deterministic)
}
func (m *PerXdsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerXdsConfig.Merge(m, src)
}
func (m *PerXdsConfig) XXX_Size() int {
	return xxx_messageInfo_PerXdsConfig.Size(m)
}
func (m *PerXdsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerXdsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerXdsConfig proto.InternalMessageInfo

func (m *PerXdsConfig) GetStatus() ConfigStatus {
	if m != nil {
		return m.Status
	}
	return ConfigStatus_UNKNOWN
}

type isPerXdsConfig_PerXdsConfig interface {
	isPerXdsConfig_PerXdsConfig()
}

type PerXdsConfig_ListenerConfig struct {
	ListenerConfig *v2alpha.ListenersConfigDump `protobuf:"bytes,2,opt,name=listener_config,json=listenerConfig,proto3,oneof"`
}

type PerXdsConfig_ClusterConfig struct {
	ClusterConfig *v2alpha.ClustersConfigDump `protobuf:"bytes,3,opt,name=cluster_config,json=clusterConfig,proto3,oneof"`
}

type PerXdsConfig_RouteConfig struct {
	RouteConfig *v2alpha.RoutesConfigDump `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3,oneof"`
}

type PerXdsConfig_ScopedRouteConfig struct {
	ScopedRouteConfig *v2alpha.ScopedRoutesConfigDump `protobuf:"bytes,5,opt,name=scoped_route_config,json=scopedRouteConfig,proto3,oneof"`
}

func (*PerXdsConfig_ListenerConfig) isPerXdsConfig_PerXdsConfig() {}

func (*PerXdsConfig_ClusterConfig) isPerXdsConfig_PerXdsConfig() {}

func (*PerXdsConfig_RouteConfig) isPerXdsConfig_PerXdsConfig() {}

func (*PerXdsConfig_ScopedRouteConfig) isPerXdsConfig_PerXdsConfig() {}

func (m *PerXdsConfig) GetPerXdsConfig() isPerXdsConfig_PerXdsConfig {
	if m != nil {
		return m.PerXdsConfig
	}
	return nil
}

func (m *PerXdsConfig) GetListenerConfig() *v2alpha.ListenersConfigDump {
	if x, ok := m.GetPerXdsConfig().(*PerXdsConfig_ListenerConfig); ok {
		return x.ListenerConfig
	}
	return nil
}

func (m *PerXdsConfig) GetClusterConfig() *v2alpha.ClustersConfigDump {
	if x, ok := m.GetPerXdsConfig().(*PerXdsConfig_ClusterConfig); ok {
		return x.ClusterConfig
	}
	return nil
}

func (m *PerXdsConfig) GetRouteConfig() *v2alpha.RoutesConfigDump {
	if x, ok := m.GetPerXdsConfig().(*PerXdsConfig_RouteConfig); ok {
		return x.RouteConfig
	}
	return nil
}

func (m *PerXdsConfig) GetScopedRouteConfig() *v2alpha.ScopedRoutesConfigDump {
	if x, ok := m.GetPerXdsConfig().(*PerXdsConfig_ScopedRouteConfig); ok {
		return x.ScopedRouteConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PerXdsConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PerXdsConfig_ListenerConfig)(nil),
		(*PerXdsConfig_ClusterConfig)(nil),
		(*PerXdsConfig_RouteConfig)(nil),
		(*PerXdsConfig_ScopedRouteConfig)(nil),
	}
}

type ClientConfig struct {
	Node                 *core.Node      `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XdsConfig            []*PerXdsConfig `protobuf:"bytes,2,rep,name=xds_config,json=xdsConfig,proto3" json:"xds_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ClientConfig) Reset()         { *m = ClientConfig{} }
func (m *ClientConfig) String() string { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()    {}
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8273221d61b8a20, []int{2}
}

func (m *ClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConfig.Unmarshal(m, b)
}
func (m *ClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConfig.Marshal(b, m, deterministic)
}
func (m *ClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConfig.Merge(m, src)
}
func (m *ClientConfig) XXX_Size() int {
	return xxx_messageInfo_ClientConfig.Size(m)
}
func (m *ClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConfig proto.InternalMessageInfo

func (m *ClientConfig) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *ClientConfig) GetXdsConfig() []*PerXdsConfig {
	if m != nil {
		return m.XdsConfig
	}
	return nil
}

type ClientStatusResponse struct {
	Config               []*ClientConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ClientStatusResponse) Reset()         { *m = ClientStatusResponse{} }
func (m *ClientStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ClientStatusResponse) ProtoMessage()    {}
func (*ClientStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8273221d61b8a20, []int{3}
}

func (m *ClientStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatusResponse.Unmarshal(m, b)
}
func (m *ClientStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatusResponse.Marshal(b, m, deterministic)
}
func (m *ClientStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatusResponse.Merge(m, src)
}
func (m *ClientStatusResponse) XXX_Size() int {
	return xxx_messageInfo_ClientStatusResponse.Size(m)
}
func (m *ClientStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatusResponse proto.InternalMessageInfo

func (m *ClientStatusResponse) GetConfig() []*ClientConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.service.status.v2.ConfigStatus", ConfigStatus_name, ConfigStatus_value)
	proto.RegisterType((*ClientStatusRequest)(nil), "envoy.service.status.v2.ClientStatusRequest")
	proto.RegisterType((*PerXdsConfig)(nil), "envoy.service.status.v2.PerXdsConfig")
	proto.RegisterType((*ClientConfig)(nil), "envoy.service.status.v2.ClientConfig")
	proto.RegisterType((*ClientStatusResponse)(nil), "envoy.service.status.v2.ClientStatusResponse")
}

func init() { proto.RegisterFile("envoy/service/status/v2/csds.proto", fileDescriptor_b8273221d61b8a20) }

var fileDescriptor_b8273221d61b8a20 = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xbb, 0x49, 0xdb, 0xef, 0xeb, 0x24, 0x0d, 0xe9, 0x16, 0xa9, 0x51, 0x69, 0x45, 0x65,
	0x29, 0x22, 0x6a, 0x8b, 0x8d, 0xcc, 0x01, 0xa9, 0x12, 0x07, 0x9a, 0x04, 0x21, 0x5a, 0xdc, 0xc8,
	0x6e, 0x05, 0x08, 0x21, 0xcb, 0xb5, 0xb7, 0xad, 0xa5, 0xc4, 0x6b, 0x76, 0xd7, 0xa6, 0x39, 0xc2,
	0x89, 0x3b, 0x77, 0xde, 0x81, 0x1b, 0xef, 0xc1, 0x2b, 0xf0, 0x20, 0xc8, 0xbb, 0x6b, 0x25, 0x41,
	0xa9, 0xca, 0x89, 0xdb, 0x5a, 0xf3, 0x9f, 0xdf, 0xcc, 0xce, 0xfc, 0xbd, 0x60, 0x90, 0x24, 0xa7,
	0x63, 0x8b, 0x13, 0x96, 0xc7, 0x21, 0xb1, 0xb8, 0x08, 0x44, 0xc6, 0xad, 0xdc, 0xb6, 0x42, 0x1e,
	0x71, 0x33, 0x65, 0x54, 0x50, 0xbc, 0x21, 0x35, 0xa6, 0xd6, 0x98, 0x4a, 0x63, 0xe6, 0xf6, 0x66,
	0x5b, 0x25, 0x07, 0xd1, 0x28, 0x4e, 0xac, 0xdc, 0x0e, 0x86, 0xe9, 0x55, 0x60, 0x85, 0x34, 0xb9,
	0x88, 0x2f, 0xfd, 0x28, 0x1b, 0xa5, 0x2a, 0x7f, 0x73, 0x4b, 0xcb, 0xd2, 0x58, 0x82, 0x29, 0x23,
	0xd6, 0x79, 0xc0, 0x89, 0x8e, 0x6e, 0xab, 0xa8, 0x18, 0xa7, 0xc4, 0x1a, 0x05, 0x22, 0xbc, 0x22,
	0xcc, 0x4a, 0x68, 0x54, 0x86, 0xb7, 0x2e, 0x29, 0xbd, 0x1c, 0x12, 0x99, 0x1d, 0x24, 0x09, 0x15,
	0x81, 0x88, 0x69, 0xc2, 0xff, 0x88, 0xca, 0xaf, 0xf3, 0xec, 0xc2, 0xe2, 0x82, 0x65, 0xa1, 0x50,
	0x51, 0xe3, 0x1d, 0xac, 0x77, 0x87, 0x31, 0x49, 0x84, 0x27, 0x5b, 0x76, 0xc9, 0x87, 0x8c, 0x70,
	0x81, 0x7b, 0xb0, 0x5a, 0x14, 0xf0, 0x75, 0x35, 0xde, 0x42, 0x3b, 0xd5, 0x4e, 0xcd, 0xbe, 0x6f,
	0xaa, 0x7b, 0x16, 0x9d, 0x98, 0x3a, 0x66, 0x3a, 0x34, 0x22, 0xaf, 0xd4, 0xd9, 0xad, 0x27, 0x93,
	0x0f, 0x6e, 0x7c, 0xaf, 0x42, 0x7d, 0x40, 0xd8, 0x9b, 0x88, 0x77, 0xe5, 0x8d, 0xf1, 0x53, 0x58,
	0x56, 0xa3, 0x69, 0xa1, 0x1d, 0xd4, 0x69, 0xd8, 0x6d, 0xf3, 0x86, 0xb9, 0x99, 0x2a, 0x41, 0x37,
	0xa5, 0x93, 0xb0, 0x07, 0x77, 0x86, 0x31, 0x17, 0x24, 0x21, 0xcc, 0x57, 0x33, 0x6c, 0x55, 0x76,
	0x50, 0xa7, 0x66, 0x77, 0x34, 0x47, 0x8e, 0xd9, 0xd4, 0x63, 0x36, 0x8f, 0xb5, 0x56, 0x57, 0xef,
	0x65, 0xa3, 0xf4, 0xc5, 0x82, 0xdb, 0x28, 0x11, 0xba, 0xa7, 0x01, 0x34, 0xc2, 0x61, 0xc6, 0xc5,
	0x84, 0x59, 0x95, 0xcc, 0x07, 0x73, 0x99, 0x5d, 0x25, 0x9d, 0x45, 0xae, 0x6a, 0x80, 0x26, 0xbe,
	0x84, 0x3a, 0xa3, 0x99, 0x20, 0x25, 0x6f, 0x51, 0xf2, 0xda, 0x73, 0x79, 0x6e, 0x21, 0x9c, 0xa5,
	0xd5, 0x64, 0xb2, 0x66, 0xbd, 0x87, 0x75, 0x1e, 0xd2, 0x94, 0x44, 0xfe, 0x0c, 0x72, 0x49, 0x22,
	0xf7, 0xe6, 0x22, 0x3d, 0xa9, 0x9f, 0x03, 0x5e, 0xe3, 0x93, 0x88, 0x0a, 0x1c, 0x36, 0xa1, 0x91,
	0x12, 0xe6, 0x5f, 0x47, 0x5c, 0x93, 0x8d, 0x4f, 0x08, 0xea, 0xca, 0x11, 0xba, 0x83, 0x3d, 0x58,
	0x2c, 0x96, 0x2a, 0x37, 0x56, 0xb3, 0x37, 0xca, 0x92, 0x69, 0x5c, 0xac, 0xa9, 0x70, 0xaa, 0x34,
	0x80, 0x2b, 0x45, 0xb8, 0x07, 0x30, 0x61, 0xb5, 0x2a, 0xd2, 0x34, 0x37, 0x2f, 0x79, 0xda, 0x1b,
	0xee, 0xca, 0x75, 0x79, 0x34, 0xce, 0xe0, 0xee, 0xac, 0x29, 0x79, 0x4a, 0x13, 0x4e, 0x0a, 0xfb,
	0x68, 0x32, 0xba, 0x85, 0x3c, 0x7d, 0x03, 0x57, 0x27, 0xed, 0x1e, 0x41, 0x7d, 0xda, 0x56, 0xb8,
	0x06, 0xff, 0x9d, 0x39, 0x47, 0xce, 0xc9, 0x6b, 0xa7, 0xb9, 0x80, 0x01, 0x96, 0xbd, 0xb7, 0x4e,
	0xb7, 0xdf, 0x6b, 0x22, 0x5c, 0x87, 0xff, 0x9d, 0x93, 0x53, 0xdf, 0xeb, 0x3b, 0xa7, 0xcd, 0x0a,
	0x5e, 0x81, 0x25, 0xef, 0xf4, 0xd9, 0x71, 0xbf, 0x59, 0x2d, 0x8e, 0x7d, 0xd7, 0x3d, 0x71, 0x9b,
	0x8b, 0xf6, 0x8f, 0x0a, 0x6c, 0x4d, 0x37, 0xd9, 0x8b, 0x79, 0x48, 0x73, 0xc2, 0xc6, 0x9e, 0x6a,
	0x06, 0x7f, 0x04, 0xec, 0x09, 0x46, 0x82, 0xd1, 0xb4, 0x0a, 0xef, 0xdf, 0xd2, 0xf2, 0xcc, 0x6f,
	0xb8, 0xf9, 0xf0, 0x2f, 0xd5, 0x6a, 0x3e, 0xc6, 0x42, 0x07, 0x3d, 0x42, 0xf8, 0x1b, 0x82, 0xb5,
	0xe7, 0x44, 0x84, 0x57, 0xff, 0xae, 0xf0, 0xfe, 0xe7, 0x9f, 0xbf, 0xbe, 0x56, 0xb6, 0x8d, 0x7b,
	0xc5, 0xeb, 0x15, 0x95, 0x93, 0x38, 0x08, 0xa5, 0xd6, 0x57, 0xc9, 0x52, 0x52, 0x3d, 0x40, 0xbb,
	0x87, 0x4f, 0xa0, 0x1d, 0x53, 0x55, 0x20, 0x65, 0xf4, 0x7a, 0x7c, 0x53, 0xad, 0xc3, 0x95, 0x2e,
	0x8f, 0xf8, 0xa0, 0x78, 0xa7, 0x06, 0xe8, 0x0b, 0x42, 0xe7, 0xcb, 0xf2, 0xcd, 0x7a, 0xfc, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0xbf, 0x1e, 0xb9, 0x92, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientStatusDiscoveryServiceClient is the client API for ClientStatusDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientStatusDiscoveryServiceClient interface {
	StreamClientStatus(ctx context.Context, opts ...grpc.CallOption) (ClientStatusDiscoveryService_StreamClientStatusClient, error)
	FetchClientStatus(ctx context.Context, in *ClientStatusRequest, opts ...grpc.CallOption) (*ClientStatusResponse, error)
}

type clientStatusDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientStatusDiscoveryServiceClient(cc *grpc.ClientConn) ClientStatusDiscoveryServiceClient {
	return &clientStatusDiscoveryServiceClient{cc}
}

func (c *clientStatusDiscoveryServiceClient) StreamClientStatus(ctx context.Context, opts ...grpc.CallOption) (ClientStatusDiscoveryService_StreamClientStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientStatusDiscoveryService_serviceDesc.Streams[0], "/envoy.service.status.v2.ClientStatusDiscoveryService/StreamClientStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStatusDiscoveryServiceStreamClientStatusClient{stream}
	return x, nil
}

type ClientStatusDiscoveryService_StreamClientStatusClient interface {
	Send(*ClientStatusRequest) error
	Recv() (*ClientStatusResponse, error)
	grpc.ClientStream
}

type clientStatusDiscoveryServiceStreamClientStatusClient struct {
	grpc.ClientStream
}

func (x *clientStatusDiscoveryServiceStreamClientStatusClient) Send(m *ClientStatusRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStatusDiscoveryServiceStreamClientStatusClient) Recv() (*ClientStatusResponse, error) {
	m := new(ClientStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientStatusDiscoveryServiceClient) FetchClientStatus(ctx context.Context, in *ClientStatusRequest, opts ...grpc.CallOption) (*ClientStatusResponse, error) {
	out := new(ClientStatusResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.status.v2.ClientStatusDiscoveryService/FetchClientStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientStatusDiscoveryServiceServer is the server API for ClientStatusDiscoveryService service.
type ClientStatusDiscoveryServiceServer interface {
	StreamClientStatus(ClientStatusDiscoveryService_StreamClientStatusServer) error
	FetchClientStatus(context.Context, *ClientStatusRequest) (*ClientStatusResponse, error)
}

// UnimplementedClientStatusDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientStatusDiscoveryServiceServer struct {
}

func (*UnimplementedClientStatusDiscoveryServiceServer) StreamClientStatus(srv ClientStatusDiscoveryService_StreamClientStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClientStatus not implemented")
}
func (*UnimplementedClientStatusDiscoveryServiceServer) FetchClientStatus(ctx context.Context, req *ClientStatusRequest) (*ClientStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClientStatus not implemented")
}

func RegisterClientStatusDiscoveryServiceServer(s *grpc.Server, srv ClientStatusDiscoveryServiceServer) {
	s.RegisterService(&_ClientStatusDiscoveryService_serviceDesc, srv)
}

func _ClientStatusDiscoveryService_StreamClientStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStatusDiscoveryServiceServer).StreamClientStatus(&clientStatusDiscoveryServiceStreamClientStatusServer{stream})
}

type ClientStatusDiscoveryService_StreamClientStatusServer interface {
	Send(*ClientStatusResponse) error
	Recv() (*ClientStatusRequest, error)
	grpc.ServerStream
}

type clientStatusDiscoveryServiceStreamClientStatusServer struct {
	grpc.ServerStream
}

func (x *clientStatusDiscoveryServiceStreamClientStatusServer) Send(m *ClientStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStatusDiscoveryServiceStreamClientStatusServer) Recv() (*ClientStatusRequest, error) {
	m := new(ClientStatusRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientStatusDiscoveryService_FetchClientStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientStatusDiscoveryServiceServer).FetchClientStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.status.v2.ClientStatusDiscoveryService/FetchClientStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientStatusDiscoveryServiceServer).FetchClientStatus(ctx, req.(*ClientStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientStatusDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.status.v2.ClientStatusDiscoveryService",
	HandlerType: (*ClientStatusDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchClientStatus",
			Handler:    _ClientStatusDiscoveryService_FetchClientStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClientStatus",
			Handler:       _ClientStatusDiscoveryService_StreamClientStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/status/v2/csds.proto",
}
