// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.1
// source: projects.proto

package grpcapi

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

type GetAllParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllParams) Reset() {
	*x = GetAllParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllParams) ProtoMessage() {}

func (x *GetAllParams) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllParams.ProtoReflect.Descriptor instead.
func (*GetAllParams) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{0}
}

type SaveReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message *string `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *SaveReturn) Reset() {
	*x = SaveReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveReturn) ProtoMessage() {}

func (x *SaveReturn) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveReturn.ProtoReflect.Descriptor instead.
func (*SaveReturn) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{1}
}

func (x *SaveReturn) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SaveReturn) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// A project is a structure used to track time spent and status.
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label       string  `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	HoursSpent  int64   `protobuf:"varint,5,opt,name=hoursSpent,proto3" json:"hoursSpent,omitempty"`
	Subprojects []int64 `protobuf:"varint,6,rep,packed,name=subprojects,proto3" json:"subprojects,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projects_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_projects_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_projects_proto_rawDescGZIP(), []int{3}
}

func (x *Project) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Project) GetHoursSpent() int64 {
	if x != nil {
		return x.HoursSpent
	}
	return 0
}

func (x *Project) GetSubprojects() []int64 {
	if x != nil {
		return x.Subprojects
	}
	return nil
}

var File_projects_proto protoreflect.FileDescriptor

var file_projects_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4f, 0x0a, 0x0a, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xab, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x6f, 0x75, 0x72, 0x73, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x75, 0x62, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32, 0xa4,
	0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0b, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x10, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x13,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_projects_proto_rawDescOnce sync.Once
	file_projects_proto_rawDescData = file_projects_proto_rawDesc
)

func file_projects_proto_rawDescGZIP() []byte {
	file_projects_proto_rawDescOnce.Do(func() {
		file_projects_proto_rawDescData = protoimpl.X.CompressGZIP(file_projects_proto_rawDescData)
	})
	return file_projects_proto_rawDescData
}

var file_projects_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_projects_proto_goTypes = []interface{}{
	(*GetAllParams)(nil), // 0: grpcapi.GetAllParams
	(*SaveReturn)(nil),   // 1: grpcapi.SaveReturn
	(*ID)(nil),           // 2: grpcapi.ID
	(*Project)(nil),      // 3: grpcapi.Project
}
var file_projects_proto_depIdxs = []int32{
	2, // 0: grpcapi.ProjectTracker.GetByID:input_type -> grpcapi.ID
	0, // 1: grpcapi.ProjectTracker.GetAll:input_type -> grpcapi.GetAllParams
	3, // 2: grpcapi.ProjectTracker.Save:input_type -> grpcapi.Project
	3, // 3: grpcapi.ProjectTracker.GetByID:output_type -> grpcapi.Project
	3, // 4: grpcapi.ProjectTracker.GetAll:output_type -> grpcapi.Project
	1, // 5: grpcapi.ProjectTracker.Save:output_type -> grpcapi.SaveReturn
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_projects_proto_init() }
func file_projects_proto_init() {
	if File_projects_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_projects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllParams); i {
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
		file_projects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveReturn); i {
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
		file_projects_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_projects_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
	file_projects_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_projects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_projects_proto_goTypes,
		DependencyIndexes: file_projects_proto_depIdxs,
		MessageInfos:      file_projects_proto_msgTypes,
	}.Build()
	File_projects_proto = out.File
	file_projects_proto_rawDesc = nil
	file_projects_proto_goTypes = nil
	file_projects_proto_depIdxs = nil
}