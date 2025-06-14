// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: common/proto/task.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// task requests
type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Token         string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"` //JWT token
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_common_proto_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTaskRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ListTasksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` //JWT token
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksRequest) Reset() {
	*x = ListTasksRequest{}
	mi := &file_common_proto_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksRequest) ProtoMessage() {}

func (x *ListTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksRequest.ProtoReflect.Descriptor instead.
func (*ListTasksRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_task_proto_rawDescGZIP(), []int{1}
}

func (x *ListTasksRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CompleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteTaskRequest) Reset() {
	*x = CompleteTaskRequest{}
	mi := &file_common_proto_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskRequest) ProtoMessage() {}

func (x *CompleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskRequest.ProtoReflect.Descriptor instead.
func (*CompleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_task_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CompleteTaskRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_common_proto_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_task_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteTaskRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// task responses
type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Completed     bool                   `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_common_proto_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_common_proto_task_proto_rawDescGZIP(), []int{4}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type TaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	mi := &file_common_proto_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type TaskListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	mi := &file_common_proto_task_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_task_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_task_proto_rawDescGZIP(), []int{6}
}

func (x *TaskListResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	mi := &file_common_proto_task_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_task_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_task_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_common_proto_task_proto protoreflect.FileDescriptor

const file_common_proto_task_proto_rawDesc = "" +
	"\n" +
	"\x17common/proto/task.proto\x12\x04task\"a\n" +
	"\x11CreateTaskRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x14\n" +
	"\x05token\x18\x03 \x01(\tR\x05token\"(\n" +
	"\x10ListTasksRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\";\n" +
	"\x13CompleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\"9\n" +
	"\x11DeleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\"l\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1c\n" +
	"\tcompleted\x18\x04 \x01(\bR\tcompleted\".\n" +
	"\fTaskResponse\x12\x1e\n" +
	"\x04task\x18\x01 \x01(\v2\n" +
	".task.TaskR\x04task\"4\n" +
	"\x10TaskListResponse\x12 \n" +
	"\x05tasks\x18\x01 \x03(\v2\n" +
	".task.TaskR\x05tasks\".\n" +
	"\x12DeleteTaskResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2\x85\x02\n" +
	"\vTaskService\x129\n" +
	"\n" +
	"CreateTask\x12\x17.task.CreateTaskRequest\x1a\x12.task.TaskResponse\x12;\n" +
	"\tListTasks\x12\x16.task.ListTasksRequest\x1a\x16.task.TaskListResponse\x12=\n" +
	"\fCompleteTask\x12\x19.task.CompleteTaskRequest\x1a\x12.task.TaskResponse\x12?\n" +
	"\n" +
	"DeleteTask\x12\x17.task.DeleteTaskRequest\x1a\x18.task.DeleteTaskResponseB\x14Z\x12task-service/protob\x06proto3"

var (
	file_common_proto_task_proto_rawDescOnce sync.Once
	file_common_proto_task_proto_rawDescData []byte
)

func file_common_proto_task_proto_rawDescGZIP() []byte {
	file_common_proto_task_proto_rawDescOnce.Do(func() {
		file_common_proto_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_proto_task_proto_rawDesc), len(file_common_proto_task_proto_rawDesc)))
	})
	return file_common_proto_task_proto_rawDescData
}

var file_common_proto_task_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_proto_task_proto_goTypes = []any{
	(*CreateTaskRequest)(nil),   // 0: task.CreateTaskRequest
	(*ListTasksRequest)(nil),    // 1: task.ListTasksRequest
	(*CompleteTaskRequest)(nil), // 2: task.CompleteTaskRequest
	(*DeleteTaskRequest)(nil),   // 3: task.DeleteTaskRequest
	(*Task)(nil),                // 4: task.Task
	(*TaskResponse)(nil),        // 5: task.TaskResponse
	(*TaskListResponse)(nil),    // 6: task.TaskListResponse
	(*DeleteTaskResponse)(nil),  // 7: task.DeleteTaskResponse
}
var file_common_proto_task_proto_depIdxs = []int32{
	4, // 0: task.TaskResponse.task:type_name -> task.Task
	4, // 1: task.TaskListResponse.tasks:type_name -> task.Task
	0, // 2: task.TaskService.CreateTask:input_type -> task.CreateTaskRequest
	1, // 3: task.TaskService.ListTasks:input_type -> task.ListTasksRequest
	2, // 4: task.TaskService.CompleteTask:input_type -> task.CompleteTaskRequest
	3, // 5: task.TaskService.DeleteTask:input_type -> task.DeleteTaskRequest
	5, // 6: task.TaskService.CreateTask:output_type -> task.TaskResponse
	6, // 7: task.TaskService.ListTasks:output_type -> task.TaskListResponse
	5, // 8: task.TaskService.CompleteTask:output_type -> task.TaskResponse
	7, // 9: task.TaskService.DeleteTask:output_type -> task.DeleteTaskResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_proto_task_proto_init() }
func file_common_proto_task_proto_init() {
	if File_common_proto_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_proto_task_proto_rawDesc), len(file_common_proto_task_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_proto_task_proto_goTypes,
		DependencyIndexes: file_common_proto_task_proto_depIdxs,
		MessageInfos:      file_common_proto_task_proto_msgTypes,
	}.Build()
	File_common_proto_task_proto = out.File
	file_common_proto_task_proto_goTypes = nil
	file_common_proto_task_proto_depIdxs = nil
}
