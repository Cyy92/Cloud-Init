// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v4.22.2
// source: fxwatcher.proto

package pb

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input []byte `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Info  *Info  `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fxwatcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_fxwatcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_fxwatcher_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Request) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=Output,proto3" json:"Output,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fxwatcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_fxwatcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_fxwatcher_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FunctionName string     `protobuf:"bytes,1,opt,name=FunctionName,proto3" json:"FunctionName,omitempty"`
	Timeout      string     `protobuf:"bytes,2,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Runtime      string     `protobuf:"bytes,3,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Limits       *Resources `protobuf:"bytes,4,opt,name=Limits,proto3" json:"Limits,omitempty"`
	Trigger      *Trigger   `protobuf:"bytes,5,opt,name=Trigger,proto3" json:"Trigger,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fxwatcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_fxwatcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_fxwatcher_proto_rawDescGZIP(), []int{2}
}

func (x *Info) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *Info) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

func (x *Info) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Info) GetLimits() *Resources {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *Info) GetTrigger() *Trigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Topic string `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Time  string `protobuf:"bytes,3,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fxwatcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_fxwatcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_fxwatcher_proto_rawDescGZIP(), []int{3}
}

func (x *Trigger) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Trigger) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Trigger) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memory string `protobuf:"bytes,1,opt,name=Memory,proto3" json:"Memory,omitempty"`
	CPU    string `protobuf:"bytes,2,opt,name=CPU,proto3" json:"CPU,omitempty"`
	GPU    string `protobuf:"bytes,3,opt,name=GPU,proto3" json:"GPU,omitempty"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fxwatcher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_fxwatcher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_fxwatcher_proto_rawDescGZIP(), []int{4}
}

func (x *Resources) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *Resources) GetCPU() string {
	if x != nil {
		return x.CPU
	}
	return ""
}

func (x *Resources) GetGPU() string {
	if x != nil {
		return x.GPU
	}
	return ""
}

var File_fxwatcher_proto protoreflect.FileDescriptor

var file_fxwatcher_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x78, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x3d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x1f, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22,
	0x0a, 0x0c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x06, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x25, 0x0a,
	0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x07, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a,
	0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x43, 0x50, 0x55, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x50, 0x55, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x47, 0x50, 0x55, 0x32, 0x2d, 0x0a, 0x09, 0x46, 0x78, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fxwatcher_proto_rawDescOnce sync.Once
	file_fxwatcher_proto_rawDescData = file_fxwatcher_proto_rawDesc
)

func file_fxwatcher_proto_rawDescGZIP() []byte {
	file_fxwatcher_proto_rawDescOnce.Do(func() {
		file_fxwatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_fxwatcher_proto_rawDescData)
	})
	return file_fxwatcher_proto_rawDescData
}

var file_fxwatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_fxwatcher_proto_goTypes = []interface{}{
	(*Request)(nil),   // 0: pb.Request
	(*Reply)(nil),     // 1: pb.Reply
	(*Info)(nil),      // 2: pb.Info
	(*Trigger)(nil),   // 3: pb.Trigger
	(*Resources)(nil), // 4: pb.Resources
}
var file_fxwatcher_proto_depIdxs = []int32{
	2, // 0: pb.Request.info:type_name -> pb.Info
	4, // 1: pb.Info.Limits:type_name -> pb.Resources
	3, // 2: pb.Info.Trigger:type_name -> pb.Trigger
	0, // 3: pb.FxWatcher.Call:input_type -> pb.Request
	1, // 4: pb.FxWatcher.Call:output_type -> pb.Reply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fxwatcher_proto_init() }
func file_fxwatcher_proto_init() {
	if File_fxwatcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fxwatcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_fxwatcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_fxwatcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_fxwatcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
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
		file_fxwatcher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
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
			RawDescriptor: file_fxwatcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fxwatcher_proto_goTypes,
		DependencyIndexes: file_fxwatcher_proto_depIdxs,
		MessageInfos:      file_fxwatcher_proto_msgTypes,
	}.Build()
	File_fxwatcher_proto = out.File
	file_fxwatcher_proto_rawDesc = nil
	file_fxwatcher_proto_goTypes = nil
	file_fxwatcher_proto_depIdxs = nil
}