// Copyright 2024 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: pkg/plugin/api/v1alpha1/platform/api.proto

package platform

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	model "github.com/pipe-cd/pipecd/pkg/model"
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

type DetermineStrategyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input *PlanPluginInput `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *DetermineStrategyRequest) Reset() {
	*x = DetermineStrategyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetermineStrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineStrategyRequest) ProtoMessage() {}

func (x *DetermineStrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineStrategyRequest.ProtoReflect.Descriptor instead.
func (*DetermineStrategyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{0}
}

func (x *DetermineStrategyRequest) GetInput() *PlanPluginInput {
	if x != nil {
		return x.Input
	}
	return nil
}

type DetermineStrategyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The determined sync strategy.
	SyncStrategy model.SyncStrategy `protobuf:"varint,1,opt,name=sync_strategy,json=syncStrategy,proto3,enum=model.SyncStrategy" json:"sync_strategy,omitempty"`
	// Text summary of the determined strategy.
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *DetermineStrategyResponse) Reset() {
	*x = DetermineStrategyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetermineStrategyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetermineStrategyResponse) ProtoMessage() {}

func (x *DetermineStrategyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetermineStrategyResponse.ProtoReflect.Descriptor instead.
func (*DetermineStrategyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{1}
}

func (x *DetermineStrategyResponse) GetSyncStrategy() model.SyncStrategy {
	if x != nil {
		return x.SyncStrategy
	}
	return model.SyncStrategy(0)
}

func (x *DetermineStrategyResponse) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type QuickSyncPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input *PlanPluginInput `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *QuickSyncPlanRequest) Reset() {
	*x = QuickSyncPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickSyncPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickSyncPlanRequest) ProtoMessage() {}

func (x *QuickSyncPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickSyncPlanRequest.ProtoReflect.Descriptor instead.
func (*QuickSyncPlanRequest) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{2}
}

func (x *QuickSyncPlanRequest) GetInput() *PlanPluginInput {
	if x != nil {
		return x.Input
	}
	return nil
}

type QuickSyncPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stages of deployment pipeline under quick sync strategy.
	Stages []*model.PipelineStage `protobuf:"bytes,1,rep,name=stages,proto3" json:"stages,omitempty"`
}

func (x *QuickSyncPlanResponse) Reset() {
	*x = QuickSyncPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickSyncPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickSyncPlanResponse) ProtoMessage() {}

func (x *QuickSyncPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickSyncPlanResponse.ProtoReflect.Descriptor instead.
func (*QuickSyncPlanResponse) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{3}
}

func (x *QuickSyncPlanResponse) GetStages() []*model.PipelineStage {
	if x != nil {
		return x.Stages
	}
	return nil
}

type PipelineSyncPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input *PlanPluginInput `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *PipelineSyncPlanRequest) Reset() {
	*x = PipelineSyncPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineSyncPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineSyncPlanRequest) ProtoMessage() {}

func (x *PipelineSyncPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineSyncPlanRequest.ProtoReflect.Descriptor instead.
func (*PipelineSyncPlanRequest) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{4}
}

func (x *PipelineSyncPlanRequest) GetInput() *PlanPluginInput {
	if x != nil {
		return x.Input
	}
	return nil
}

type PipelineSyncPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stages of deployment pipeline under pipeline sync strategy.
	Stages []*model.PipelineStage `protobuf:"bytes,1,rep,name=stages,proto3" json:"stages,omitempty"`
}

func (x *PipelineSyncPlanResponse) Reset() {
	*x = PipelineSyncPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineSyncPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineSyncPlanResponse) ProtoMessage() {}

func (x *PipelineSyncPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineSyncPlanResponse.ProtoReflect.Descriptor instead.
func (*PipelineSyncPlanResponse) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{5}
}

func (x *PipelineSyncPlanResponse) GetStages() []*model.PipelineStage {
	if x != nil {
		return x.Stages
	}
	return nil
}

type PlanPluginInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The deployment to build a plan for.
	Deployment *model.Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	// The remote URL of the deployment source, where plugin can find the deployments sources (manifests).
	SourceRemoteUrl string `protobuf:"bytes,2,opt,name=source_remote_url,json=sourceRemoteUrl,proto3" json:"source_remote_url,omitempty"`
	// Last successful commit hash and config file name.
	// Use to build deployment source object for last successful deployment.
	LastSuccessfulCommitHash     string `protobuf:"bytes,3,opt,name=last_successful_commit_hash,json=lastSuccessfulCommitHash,proto3" json:"last_successful_commit_hash,omitempty"`
	LastSuccessfulConfigFileName string `protobuf:"bytes,4,opt,name=last_successful_config_file_name,json=lastSuccessfulConfigFileName,proto3" json:"last_successful_config_file_name,omitempty"`
	// The configuration of plugin that handles the deployment.
	PluginConfig []byte `protobuf:"bytes,5,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
}

func (x *PlanPluginInput) Reset() {
	*x = PlanPluginInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanPluginInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanPluginInput) ProtoMessage() {}

func (x *PlanPluginInput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanPluginInput.ProtoReflect.Descriptor instead.
func (*PlanPluginInput) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{6}
}

func (x *PlanPluginInput) GetDeployment() *model.Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *PlanPluginInput) GetSourceRemoteUrl() string {
	if x != nil {
		return x.SourceRemoteUrl
	}
	return ""
}

func (x *PlanPluginInput) GetLastSuccessfulCommitHash() string {
	if x != nil {
		return x.LastSuccessfulCommitHash
	}
	return ""
}

func (x *PlanPluginInput) GetLastSuccessfulConfigFileName() string {
	if x != nil {
		return x.LastSuccessfulConfigFileName
	}
	return ""
}

func (x *PlanPluginInput) GetPluginConfig() []byte {
	if x != nil {
		return x.PluginConfig
	}
	return nil
}

type ExecuteStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage       *model.PipelineStage `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage,omitempty"`
	StageConfig []byte               `protobuf:"bytes,2,opt,name=stage_config,json=stageConfig,proto3" json:"stage_config,omitempty"`
	PipedConfig []byte               `protobuf:"bytes,3,opt,name=piped_config,json=pipedConfig,proto3" json:"piped_config,omitempty"`
	Deployment  *model.Deployment    `protobuf:"bytes,4,opt,name=deployment,proto3" json:"deployment,omitempty"`
}

func (x *ExecuteStageRequest) Reset() {
	*x = ExecuteStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteStageRequest) ProtoMessage() {}

func (x *ExecuteStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteStageRequest.ProtoReflect.Descriptor instead.
func (*ExecuteStageRequest) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{7}
}

func (x *ExecuteStageRequest) GetStage() *model.PipelineStage {
	if x != nil {
		return x.Stage
	}
	return nil
}

func (x *ExecuteStageRequest) GetStageConfig() []byte {
	if x != nil {
		return x.StageConfig
	}
	return nil
}

func (x *ExecuteStageRequest) GetPipedConfig() []byte {
	if x != nil {
		return x.PipedConfig
	}
	return nil
}

func (x *ExecuteStageRequest) GetDeployment() *model.Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

type ExecuteStageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status model.StageStatus `protobuf:"varint,1,opt,name=status,proto3,enum=model.StageStatus" json:"status,omitempty"`
	Log    string            `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *ExecuteStageResponse) Reset() {
	*x = ExecuteStageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteStageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteStageResponse) ProtoMessage() {}

func (x *ExecuteStageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteStageResponse.ProtoReflect.Descriptor instead.
func (*ExecuteStageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP(), []int{8}
}

func (x *ExecuteStageResponse) GetStatus() model.StageStatus {
	if x != nil {
		return x.Status
	}
	return model.StageStatus(0)
}

func (x *ExecuteStageResponse) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

var File_pkg_plugin_api_v1alpha1_platform_api_proto protoreflect.FileDescriptor

var file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x18, 0x44,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x6f, 0x0a, 0x19, 0x44, 0x65,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x14, 0x51,
	0x75, 0x69, 0x63, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x53,
	0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x22, 0x6c, 0x0a,
	0x17, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x48, 0x0a, 0x18, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x73, 0x22, 0xaf, 0x02, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x3d, 0x0a, 0x1b, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x46, 0x0a, 0x20, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x66, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xe0, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x7a, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2a, 0x0a, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x10, 0x01,
	0x52, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a,
	0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x14, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67,
	0x32, 0xb4, 0x03, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x3a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x0d, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x36, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x53,
	0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x10, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x39,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x95, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0c,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x35, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x2d, 0x63, 0x64, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x63, 0x64, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescOnce sync.Once
	file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescData = file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDesc
)

func file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescGZIP() []byte {
	file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescOnce.Do(func() {
		file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescData)
	})
	return file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDescData
}

var file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_plugin_api_v1alpha1_platform_api_proto_goTypes = []interface{}{
	(*DetermineStrategyRequest)(nil),  // 0: grpc.plugin.platformapi.v1alpha1.DetermineStrategyRequest
	(*DetermineStrategyResponse)(nil), // 1: grpc.plugin.platformapi.v1alpha1.DetermineStrategyResponse
	(*QuickSyncPlanRequest)(nil),      // 2: grpc.plugin.platformapi.v1alpha1.QuickSyncPlanRequest
	(*QuickSyncPlanResponse)(nil),     // 3: grpc.plugin.platformapi.v1alpha1.QuickSyncPlanResponse
	(*PipelineSyncPlanRequest)(nil),   // 4: grpc.plugin.platformapi.v1alpha1.PipelineSyncPlanRequest
	(*PipelineSyncPlanResponse)(nil),  // 5: grpc.plugin.platformapi.v1alpha1.PipelineSyncPlanResponse
	(*PlanPluginInput)(nil),           // 6: grpc.plugin.platformapi.v1alpha1.PlanPluginInput
	(*ExecuteStageRequest)(nil),       // 7: grpc.plugin.platformapi.v1alpha1.ExecuteStageRequest
	(*ExecuteStageResponse)(nil),      // 8: grpc.plugin.platformapi.v1alpha1.ExecuteStageResponse
	(model.SyncStrategy)(0),           // 9: model.SyncStrategy
	(*model.PipelineStage)(nil),       // 10: model.PipelineStage
	(*model.Deployment)(nil),          // 11: model.Deployment
	(model.StageStatus)(0),            // 12: model.StageStatus
}
var file_pkg_plugin_api_v1alpha1_platform_api_proto_depIdxs = []int32{
	6,  // 0: grpc.plugin.platformapi.v1alpha1.DetermineStrategyRequest.input:type_name -> grpc.plugin.platformapi.v1alpha1.PlanPluginInput
	9,  // 1: grpc.plugin.platformapi.v1alpha1.DetermineStrategyResponse.sync_strategy:type_name -> model.SyncStrategy
	6,  // 2: grpc.plugin.platformapi.v1alpha1.QuickSyncPlanRequest.input:type_name -> grpc.plugin.platformapi.v1alpha1.PlanPluginInput
	10, // 3: grpc.plugin.platformapi.v1alpha1.QuickSyncPlanResponse.stages:type_name -> model.PipelineStage
	6,  // 4: grpc.plugin.platformapi.v1alpha1.PipelineSyncPlanRequest.input:type_name -> grpc.plugin.platformapi.v1alpha1.PlanPluginInput
	10, // 5: grpc.plugin.platformapi.v1alpha1.PipelineSyncPlanResponse.stages:type_name -> model.PipelineStage
	11, // 6: grpc.plugin.platformapi.v1alpha1.PlanPluginInput.deployment:type_name -> model.Deployment
	10, // 7: grpc.plugin.platformapi.v1alpha1.ExecuteStageRequest.stage:type_name -> model.PipelineStage
	11, // 8: grpc.plugin.platformapi.v1alpha1.ExecuteStageRequest.deployment:type_name -> model.Deployment
	12, // 9: grpc.plugin.platformapi.v1alpha1.ExecuteStageResponse.status:type_name -> model.StageStatus
	0,  // 10: grpc.plugin.platformapi.v1alpha1.PlannerService.DetermineStrategy:input_type -> grpc.plugin.platformapi.v1alpha1.DetermineStrategyRequest
	2,  // 11: grpc.plugin.platformapi.v1alpha1.PlannerService.QuickSyncPlan:input_type -> grpc.plugin.platformapi.v1alpha1.QuickSyncPlanRequest
	4,  // 12: grpc.plugin.platformapi.v1alpha1.PlannerService.PipelineSyncPlan:input_type -> grpc.plugin.platformapi.v1alpha1.PipelineSyncPlanRequest
	7,  // 13: grpc.plugin.platformapi.v1alpha1.ExecutorService.ExecuteStage:input_type -> grpc.plugin.platformapi.v1alpha1.ExecuteStageRequest
	1,  // 14: grpc.plugin.platformapi.v1alpha1.PlannerService.DetermineStrategy:output_type -> grpc.plugin.platformapi.v1alpha1.DetermineStrategyResponse
	3,  // 15: grpc.plugin.platformapi.v1alpha1.PlannerService.QuickSyncPlan:output_type -> grpc.plugin.platformapi.v1alpha1.QuickSyncPlanResponse
	5,  // 16: grpc.plugin.platformapi.v1alpha1.PlannerService.PipelineSyncPlan:output_type -> grpc.plugin.platformapi.v1alpha1.PipelineSyncPlanResponse
	8,  // 17: grpc.plugin.platformapi.v1alpha1.ExecutorService.ExecuteStage:output_type -> grpc.plugin.platformapi.v1alpha1.ExecuteStageResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_plugin_api_v1alpha1_platform_api_proto_init() }
func file_pkg_plugin_api_v1alpha1_platform_api_proto_init() {
	if File_pkg_plugin_api_v1alpha1_platform_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetermineStrategyRequest); i {
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
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetermineStrategyResponse); i {
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
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuickSyncPlanRequest); i {
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
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuickSyncPlanResponse); i {
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
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineSyncPlanRequest); i {
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
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineSyncPlanResponse); i {
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
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanPluginInput); i {
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
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteStageRequest); i {
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
		file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteStageResponse); i {
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
			RawDescriptor: file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_plugin_api_v1alpha1_platform_api_proto_goTypes,
		DependencyIndexes: file_pkg_plugin_api_v1alpha1_platform_api_proto_depIdxs,
		MessageInfos:      file_pkg_plugin_api_v1alpha1_platform_api_proto_msgTypes,
	}.Build()
	File_pkg_plugin_api_v1alpha1_platform_api_proto = out.File
	file_pkg_plugin_api_v1alpha1_platform_api_proto_rawDesc = nil
	file_pkg_plugin_api_v1alpha1_platform_api_proto_goTypes = nil
	file_pkg_plugin_api_v1alpha1_platform_api_proto_depIdxs = nil
}