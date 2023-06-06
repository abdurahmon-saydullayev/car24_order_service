// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: mechanic.proto

package order_service

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

type Mechanic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fullname     string `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	PhoneNumber  string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Photo        string `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
	PricePerHour string `protobuf:"bytes,5,opt,name=price_per_hour,json=pricePerHour,proto3" json:"price_per_hour,omitempty"`
	Status       bool   `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Mechanic) Reset() {
	*x = Mechanic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mechanic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mechanic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mechanic) ProtoMessage() {}

func (x *Mechanic) ProtoReflect() protoreflect.Message {
	mi := &file_mechanic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mechanic.ProtoReflect.Descriptor instead.
func (*Mechanic) Descriptor() ([]byte, []int) {
	return file_mechanic_proto_rawDescGZIP(), []int{0}
}

func (x *Mechanic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Mechanic) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *Mechanic) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Mechanic) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Mechanic) GetPricePerHour() string {
	if x != nil {
		return x.PricePerHour
	}
	return ""
}

func (x *Mechanic) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type CreateMechanic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname     string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	PhoneNumber  string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Photo        string `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo,omitempty"`
	PricePerHour bool   `protobuf:"varint,4,opt,name=price_per_hour,json=pricePerHour,proto3" json:"price_per_hour,omitempty"`
}

func (x *CreateMechanic) Reset() {
	*x = CreateMechanic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mechanic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMechanic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMechanic) ProtoMessage() {}

func (x *CreateMechanic) ProtoReflect() protoreflect.Message {
	mi := &file_mechanic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMechanic.ProtoReflect.Descriptor instead.
func (*CreateMechanic) Descriptor() ([]byte, []int) {
	return file_mechanic_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMechanic) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreateMechanic) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateMechanic) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CreateMechanic) GetPricePerHour() bool {
	if x != nil {
		return x.PricePerHour
	}
	return false
}

type UpdateMechanic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateMechanic) Reset() {
	*x = UpdateMechanic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mechanic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMechanic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMechanic) ProtoMessage() {}

func (x *UpdateMechanic) ProtoReflect() protoreflect.Message {
	mi := &file_mechanic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMechanic.ProtoReflect.Descriptor instead.
func (*UpdateMechanic) Descriptor() ([]byte, []int) {
	return file_mechanic_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMechanic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMechanic) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type MechanicPK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MechanicPK) Reset() {
	*x = MechanicPK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mechanic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MechanicPK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicPK) ProtoMessage() {}

func (x *MechanicPK) ProtoReflect() protoreflect.Message {
	mi := &file_mechanic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicPK.ProtoReflect.Descriptor instead.
func (*MechanicPK) Descriptor() ([]byte, []int) {
	return file_mechanic_proto_rawDescGZIP(), []int{3}
}

func (x *MechanicPK) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_mechanic_proto protoreflect.FileDescriptor

var file_mechanic_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0xad, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x68,
	0x6f, 0x75, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x8b, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x72, 0x22, 0x38, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1c, 0x0a, 0x0a, 0x4d, 0x65, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x63, 0x50, 0x4b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mechanic_proto_rawDescOnce sync.Once
	file_mechanic_proto_rawDescData = file_mechanic_proto_rawDesc
)

func file_mechanic_proto_rawDescGZIP() []byte {
	file_mechanic_proto_rawDescOnce.Do(func() {
		file_mechanic_proto_rawDescData = protoimpl.X.CompressGZIP(file_mechanic_proto_rawDescData)
	})
	return file_mechanic_proto_rawDescData
}

var file_mechanic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mechanic_proto_goTypes = []interface{}{
	(*Mechanic)(nil),       // 0: order_service.Mechanic
	(*CreateMechanic)(nil), // 1: order_service.CreateMechanic
	(*UpdateMechanic)(nil), // 2: order_service.UpdateMechanic
	(*MechanicPK)(nil),     // 3: order_service.MechanicPK
}
var file_mechanic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mechanic_proto_init() }
func file_mechanic_proto_init() {
	if File_mechanic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mechanic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mechanic); i {
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
		file_mechanic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMechanic); i {
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
		file_mechanic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMechanic); i {
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
		file_mechanic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MechanicPK); i {
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
			RawDescriptor: file_mechanic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mechanic_proto_goTypes,
		DependencyIndexes: file_mechanic_proto_depIdxs,
		MessageInfos:      file_mechanic_proto_msgTypes,
	}.Build()
	File_mechanic_proto = out.File
	file_mechanic_proto_rawDesc = nil
	file_mechanic_proto_goTypes = nil
	file_mechanic_proto_depIdxs = nil
}
