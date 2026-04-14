// Code generated manually to keep protobuf stubs in-repo. DO NOT EDIT.

package ratesv1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

type CalculationMethod int32

const (
	CalculationMethod_CALCULATION_METHOD_UNSPECIFIED CalculationMethod = 0
	CalculationMethod_CALCULATION_METHOD_TOP_N       CalculationMethod = 1
	CalculationMethod_CALCULATION_METHOD_AVG_N_M     CalculationMethod = 2
)

var (
	CalculationMethod_name = map[int32]string{
		0: "CALCULATION_METHOD_UNSPECIFIED",
		1: "CALCULATION_METHOD_TOP_N",
		2: "CALCULATION_METHOD_AVG_N_M",
	}
	CalculationMethod_value = map[string]int32{
		"CALCULATION_METHOD_UNSPECIFIED": 0,
		"CALCULATION_METHOD_TOP_N":       1,
		"CALCULATION_METHOD_AVG_N_M":     2,
	}
)

func (x CalculationMethod) Enum() *CalculationMethod {
	p := new(CalculationMethod)
	*p = x
	return p
}

func (x CalculationMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CalculationMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_rates_v1_rates_proto_enumTypes[0].Descriptor()
}

func (CalculationMethod) Type() protoreflect.EnumType {
	return &file_rates_v1_rates_proto_enumTypes[0]
}

func (x CalculationMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type GetRatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method CalculationMethod `protobuf:"varint,1,opt,name=method,proto3,enum=rates.v1.CalculationMethod" json:"method,omitempty"`
	N      uint32            `protobuf:"varint,2,opt,name=n,proto3" json:"n,omitempty"`
	M      uint32            `protobuf:"varint,3,opt,name=m,proto3" json:"m,omitempty"`
}

func (x *GetRatesRequest) Reset() {
	*x = GetRatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rates_v1_rates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesRequest) ProtoMessage() {}

func (x *GetRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rates_v1_rates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetRatesRequest) Descriptor() ([]byte, []int) {
	return file_rates_v1_rates_proto_rawDescGZIP(), []int{0}
}

func (x *GetRatesRequest) GetMethod() CalculationMethod {
	if x != nil {
		return x.Method
	}
	return CalculationMethod_CALCULATION_METHOD_UNSPECIFIED
}

func (x *GetRatesRequest) GetN() uint32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *GetRatesRequest) GetM() uint32 {
	if x != nil {
		return x.M
	}
	return 0
}

type GetRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ask         string                 `protobuf:"bytes,1,opt,name=ask,proto3" json:"ask,omitempty"`
	Bid         string                 `protobuf:"bytes,2,opt,name=bid,proto3" json:"bid,omitempty"`
	RetrievedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=retrieved_at,json=retrievedAt,proto3" json:"retrieved_at,omitempty"`
}

func (x *GetRatesResponse) Reset() {
	*x = GetRatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rates_v1_rates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesResponse) ProtoMessage() {}

func (x *GetRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rates_v1_rates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return file_rates_v1_rates_proto_rawDescGZIP(), []int{1}
}

func (x *GetRatesResponse) GetAsk() string {
	if x != nil {
		return x.Ask
	}
	return ""
}

func (x *GetRatesResponse) GetBid() string {
	if x != nil {
		return x.Bid
	}
	return ""
}

func (x *GetRatesResponse) GetRetrievedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RetrievedAt
	}
	return nil
}

var File_rates_v1_rates_proto protoreflect.FileDescriptor

var file_rates_v1_rates_proto_rawDescOnce sync.Once
var file_rates_v1_rates_proto_rawDescData []byte

func file_rates_v1_rates_proto_rawDescGZIP() []byte {
	file_rates_v1_rates_proto_rawDescOnce.Do(func() {
		file_rates_v1_rates_proto_rawDescData = protoimpl.X.CompressGZIP(file_rates_v1_rates_proto_rawDescData)
	})
	return file_rates_v1_rates_proto_rawDescData
}

var file_rates_v1_rates_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rates_v1_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rates_v1_rates_proto_goTypes = []any{
	(CalculationMethod)(0),
	(*GetRatesRequest)(nil),
	(*GetRatesResponse)(nil),
	(*timestamppb.Timestamp)(nil),
}
var file_rates_v1_rates_proto_depIdxs = []int32{
	0, // 0: rates.v1.GetRatesRequest.method:type_name -> rates.v1.CalculationMethod
	3, // 1: rates.v1.GetRatesResponse.retrieved_at:type_name -> google.protobuf.Timestamp
	1, // 2: rates.v1.RatesService.GetRates:input_type -> rates.v1.GetRatesRequest
	2, // 3: rates.v1.RatesService.GetRates:output_type -> rates.v1.GetRatesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rates_v1_rates_proto_init() }

func file_rates_v1_rates_proto_init() {
	if File_rates_v1_rates_proto != nil {
		return
	}

	fileDescriptor := &descriptorpb.FileDescriptorProto{
		Syntax:     proto.String("proto3"),
		Name:       proto.String("rates/v1/rates.proto"),
		Package:    proto.String("rates.v1"),
		Dependency: []string{"google/protobuf/timestamp.proto"},
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("exchange_rate/gen/rates/v1;ratesv1"),
		},
		EnumType: []*descriptorpb.EnumDescriptorProto{
			{
				Name: proto.String("CalculationMethod"),
				Value: []*descriptorpb.EnumValueDescriptorProto{
					{Name: proto.String("CALCULATION_METHOD_UNSPECIFIED"), Number: proto.Int32(0)},
					{Name: proto.String("CALCULATION_METHOD_TOP_N"), Number: proto.Int32(1)},
					{Name: proto.String("CALCULATION_METHOD_AVG_N_M"), Number: proto.Int32(2)},
				},
			},
		},
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("GetRatesRequest"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     proto.String("method"),
						Number:   proto.Int32(1),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
						TypeName: proto.String(".rates.v1.CalculationMethod"),
					},
					{
						Name:   proto.String("n"),
						Number: proto.Int32(2),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_UINT32.Enum(),
					},
					{
						Name:   proto.String("m"),
						Number: proto.Int32(3),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_UINT32.Enum(),
					},
				},
			},
			{
				Name: proto.String("GetRatesResponse"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("ask"),
						Number: proto.Int32(1),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
					{
						Name:   proto.String("bid"),
						Number: proto.Int32(2),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
					{
						Name:     proto.String("retrieved_at"),
						JsonName: proto.String("retrievedAt"),
						Number:   proto.Int32(3),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: proto.String(".google.protobuf.Timestamp"),
					},
				},
			},
		},
		Service: []*descriptorpb.ServiceDescriptorProto{
			{
				Name: proto.String("RatesService"),
				Method: []*descriptorpb.MethodDescriptorProto{
					{
						Name:       proto.String("GetRates"),
						InputType:  proto.String(".rates.v1.GetRatesRequest"),
						OutputType: proto.String(".rates.v1.GetRatesResponse"),
					},
				},
			},
		},
	}

	rawDescriptor, err := proto.Marshal(fileDescriptor)
	if err != nil {
		panic(err)
	}
	file_rates_v1_rates_proto_rawDescData = rawDescriptor

	if !protoimpl.UnsafeEnabled {
		file_rates_v1_rates_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetRatesRequest); i {
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
		file_rates_v1_rates_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetRatesResponse); i {
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
			RawDescriptor: file_rates_v1_rates_proto_rawDescData,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rates_v1_rates_proto_goTypes,
		DependencyIndexes: file_rates_v1_rates_proto_depIdxs,
		EnumInfos:         file_rates_v1_rates_proto_enumTypes,
		MessageInfos:      file_rates_v1_rates_proto_msgTypes,
	}.Build()

	File_rates_v1_rates_proto = out.File
	file_rates_v1_rates_proto_rawDescData = nil
	file_rates_v1_rates_proto_goTypes = nil
	file_rates_v1_rates_proto_depIdxs = nil
}
