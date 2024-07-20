// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: todo.proto

package __contract

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TodoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Todo        string `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
	Details     string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	CreatedDate string `protobuf:"bytes,4,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
}

func (x *TodoItem) Reset() {
	*x = TodoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItem) ProtoMessage() {}

func (x *TodoItem) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItem.ProtoReflect.Descriptor instead.
func (*TodoItem) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoItem) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

func (x *TodoItem) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *TodoItem) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

type TodoItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todolist []*TodoItem `protobuf:"bytes,1,rep,name=todolist,proto3" json:"todolist,omitempty"`
}

func (x *TodoItems) Reset() {
	*x = TodoItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItems) ProtoMessage() {}

func (x *TodoItems) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItems.ProtoReflect.Descriptor instead.
func (*TodoItems) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoItems) GetTodolist() []*TodoItem {
	if x != nil {
		return x.Todolist
	}
	return nil
}

type AddTodoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo    string `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *AddTodoReq) Reset() {
	*x = AddTodoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodoReq) ProtoMessage() {}

func (x *AddTodoReq) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodoReq.ProtoReflect.Descriptor instead.
func (*AddTodoReq) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *AddTodoReq) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

func (x *AddTodoReq) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x22, 0x3b, 0x0a, 0x09, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2e, 0x0a,
	0x08, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x08, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0x7e, 0x0a, 0x0a, 0x54, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x39, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_todo_proto_goTypes = []interface{}{
	(*TodoItem)(nil),      // 0: contract.TodoItem
	(*TodoItems)(nil),     // 1: contract.TodoItems
	(*AddTodoReq)(nil),    // 2: contract.AddTodoReq
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_todo_proto_depIdxs = []int32{
	0, // 0: contract.TodoItems.todolist:type_name -> contract.TodoItem
	3, // 1: contract.ThingsTodo.TaskList:input_type -> google.protobuf.Empty
	2, // 2: contract.ThingsTodo.AddTask:input_type -> contract.AddTodoReq
	1, // 3: contract.ThingsTodo.TaskList:output_type -> contract.TodoItems
	0, // 4: contract.ThingsTodo.AddTask:output_type -> contract.TodoItem
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItem); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItems); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodoReq); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}