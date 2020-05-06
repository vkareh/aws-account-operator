// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/awsclient/client.go

// Package mock is a generated GoMock package.
package mock

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	iam "github.com/aws/aws-sdk-go/service/iam"
	organizations "github.com/aws/aws-sdk-go/service/organizations"
	route53 "github.com/aws/aws-sdk-go/service/route53"
	s3 "github.com/aws/aws-sdk-go/service/s3"
	sts "github.com/aws/aws-sdk-go/service/sts"
	support "github.com/aws/aws-sdk-go/service/support"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// RunInstances mocks base method.
func (m *MockClient) RunInstances(arg0 *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInstances", arg0)
	ret0, _ := ret[0].(*ec2.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunInstances indicates an expected call of RunInstances.
func (mr *MockClientMockRecorder) RunInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInstances", reflect.TypeOf((*MockClient)(nil).RunInstances), arg0)
}

// DescribeInstanceStatus mocks base method.
func (m *MockClient) DescribeInstanceStatus(arg0 *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstanceStatus", arg0)
	ret0, _ := ret[0].(*ec2.DescribeInstanceStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstanceStatus indicates an expected call of DescribeInstanceStatus.
func (mr *MockClientMockRecorder) DescribeInstanceStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceStatus", reflect.TypeOf((*MockClient)(nil).DescribeInstanceStatus), arg0)
}

// TerminateInstances mocks base method.
func (m *MockClient) TerminateInstances(arg0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateInstances", arg0)
	ret0, _ := ret[0].(*ec2.TerminateInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TerminateInstances indicates an expected call of TerminateInstances.
func (mr *MockClientMockRecorder) TerminateInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstances", reflect.TypeOf((*MockClient)(nil).TerminateInstances), arg0)
}

// DescribeVolumes mocks base method.
func (m *MockClient) DescribeVolumes(arg0 *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVolumes", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVolumes indicates an expected call of DescribeVolumes.
func (mr *MockClientMockRecorder) DescribeVolumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVolumes", reflect.TypeOf((*MockClient)(nil).DescribeVolumes), arg0)
}

// DeleteVolume mocks base method.
func (m *MockClient) DeleteVolume(arg0 *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0)
	ret0, _ := ret[0].(*ec2.DeleteVolumeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockClientMockRecorder) DeleteVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockClient)(nil).DeleteVolume), arg0)
}

// DescribeSnapshots mocks base method.
func (m *MockClient) DescribeSnapshots(arg0 *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSnapshots", arg0)
	ret0, _ := ret[0].(*ec2.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSnapshots indicates an expected call of DescribeSnapshots.
func (mr *MockClientMockRecorder) DescribeSnapshots(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshots", reflect.TypeOf((*MockClient)(nil).DescribeSnapshots), arg0)
}

// DeleteSnapshot mocks base method.
func (m *MockClient) DeleteSnapshot(arg0 *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0)
	ret0, _ := ret[0].(*ec2.DeleteSnapshotOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockClientMockRecorder) DeleteSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockClient)(nil).DeleteSnapshot), arg0)
}

// CreateAccessKey mocks base method.
func (m *MockClient) CreateAccessKey(arg0 *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessKey", arg0)
	ret0, _ := ret[0].(*iam.CreateAccessKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessKey indicates an expected call of CreateAccessKey.
func (mr *MockClientMockRecorder) CreateAccessKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessKey", reflect.TypeOf((*MockClient)(nil).CreateAccessKey), arg0)
}

// CreateUser mocks base method.
func (m *MockClient) CreateUser(arg0 *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*iam.CreateUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockClientMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockClient)(nil).CreateUser), arg0)
}

// DeleteAccessKey mocks base method.
func (m *MockClient) DeleteAccessKey(arg0 *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessKey", arg0)
	ret0, _ := ret[0].(*iam.DeleteAccessKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccessKey indicates an expected call of DeleteAccessKey.
func (mr *MockClientMockRecorder) DeleteAccessKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessKey", reflect.TypeOf((*MockClient)(nil).DeleteAccessKey), arg0)
}

// DeleteUser mocks base method.
func (m *MockClient) DeleteUser(arg0 *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(*iam.DeleteUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockClientMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockClient)(nil).DeleteUser), arg0)
}

// DeleteUserPolicy mocks base method.
func (m *MockClient) DeleteUserPolicy(arg0 *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.DeleteUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserPolicy indicates an expected call of DeleteUserPolicy.
func (mr *MockClientMockRecorder) DeleteUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPolicy", reflect.TypeOf((*MockClient)(nil).DeleteUserPolicy), arg0)
}

// GetUser mocks base method.
func (m *MockClient) GetUser(arg0 *iam.GetUserInput) (*iam.GetUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*iam.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockClientMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockClient)(nil).GetUser), arg0)
}

// ListUsers mocks base method.
func (m *MockClient) ListUsers(arg0 *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].(*iam.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockClientMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockClient)(nil).ListUsers), arg0)
}

// ListUsersPages mocks base method.
func (m *MockClient) ListUsersPages(arg0 *iam.ListUsersInput, arg1 func(*iam.ListUsersOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsersPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListUsersPages indicates an expected call of ListUsersPages.
func (mr *MockClientMockRecorder) ListUsersPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersPages", reflect.TypeOf((*MockClient)(nil).ListUsersPages), arg0, arg1)
}

// ListUserTags mocks base method.
func (m *MockClient) ListUserTags(arg0 *iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserTags", arg0)
	ret0, _ := ret[0].(*iam.ListUserTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserTags indicates an expected call of ListUserTags.
func (mr *MockClientMockRecorder) ListUserTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserTags", reflect.TypeOf((*MockClient)(nil).ListUserTags), arg0)
}

// ListAccessKeys mocks base method.
func (m *MockClient) ListAccessKeys(arg0 *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessKeys", arg0)
	ret0, _ := ret[0].(*iam.ListAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessKeys indicates an expected call of ListAccessKeys.
func (mr *MockClientMockRecorder) ListAccessKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessKeys", reflect.TypeOf((*MockClient)(nil).ListAccessKeys), arg0)
}

// ListUserPolicies mocks base method.
func (m *MockClient) ListUserPolicies(arg0 *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserPolicies", arg0)
	ret0, _ := ret[0].(*iam.ListUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserPolicies indicates an expected call of ListUserPolicies.
func (mr *MockClientMockRecorder) ListUserPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserPolicies", reflect.TypeOf((*MockClient)(nil).ListUserPolicies), arg0)
}

// PutUserPolicy mocks base method.
func (m *MockClient) PutUserPolicy(arg0 *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.PutUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutUserPolicy indicates an expected call of PutUserPolicy.
func (mr *MockClientMockRecorder) PutUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUserPolicy", reflect.TypeOf((*MockClient)(nil).PutUserPolicy), arg0)
}

// AttachUserPolicy mocks base method.
func (m *MockClient) AttachUserPolicy(arg0 *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.AttachUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachUserPolicy indicates an expected call of AttachUserPolicy.
func (mr *MockClientMockRecorder) AttachUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachUserPolicy", reflect.TypeOf((*MockClient)(nil).AttachUserPolicy), arg0)
}

// DetachUserPolicy mocks base method.
func (m *MockClient) DetachUserPolicy(arg0 *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.DetachUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachUserPolicy indicates an expected call of DetachUserPolicy.
func (mr *MockClientMockRecorder) DetachUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachUserPolicy", reflect.TypeOf((*MockClient)(nil).DetachUserPolicy), arg0)
}

// ListPolicies mocks base method.
func (m *MockClient) ListPolicies(arg0 *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPolicies", arg0)
	ret0, _ := ret[0].(*iam.ListPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicies indicates an expected call of ListPolicies.
func (mr *MockClientMockRecorder) ListPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicies", reflect.TypeOf((*MockClient)(nil).ListPolicies), arg0)
}

// ListAttachedUserPolicies mocks base method.
func (m *MockClient) ListAttachedUserPolicies(arg0 *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttachedUserPolicies", arg0)
	ret0, _ := ret[0].(*iam.ListAttachedUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttachedUserPolicies indicates an expected call of ListAttachedUserPolicies.
func (mr *MockClientMockRecorder) ListAttachedUserPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedUserPolicies", reflect.TypeOf((*MockClient)(nil).ListAttachedUserPolicies), arg0)
}

// CreatePolicy mocks base method.
func (m *MockClient) CreatePolicy(arg0 *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePolicy", arg0)
	ret0, _ := ret[0].(*iam.CreatePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePolicy indicates an expected call of CreatePolicy.
func (mr *MockClientMockRecorder) CreatePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePolicy", reflect.TypeOf((*MockClient)(nil).CreatePolicy), arg0)
}

// DeletePolicy mocks base method.
func (m *MockClient) DeletePolicy(input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicy", input)
	ret0, _ := ret[0].(*iam.DeletePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePolicy indicates an expected call of DeletePolicy.
func (mr *MockClientMockRecorder) DeletePolicy(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicy", reflect.TypeOf((*MockClient)(nil).DeletePolicy), input)
}

// AttachRolePolicy mocks base method.
func (m *MockClient) AttachRolePolicy(arg0 *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachRolePolicy", arg0)
	ret0, _ := ret[0].(*iam.AttachRolePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachRolePolicy indicates an expected call of AttachRolePolicy.
func (mr *MockClientMockRecorder) AttachRolePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachRolePolicy", reflect.TypeOf((*MockClient)(nil).AttachRolePolicy), arg0)
}

// DetachRolePolicy mocks base method.
func (m *MockClient) DetachRolePolicy(arg0 *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachRolePolicy", arg0)
	ret0, _ := ret[0].(*iam.DetachRolePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachRolePolicy indicates an expected call of DetachRolePolicy.
func (mr *MockClientMockRecorder) DetachRolePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachRolePolicy", reflect.TypeOf((*MockClient)(nil).DetachRolePolicy), arg0)
}

// ListAttachedRolePolicies mocks base method.
func (m *MockClient) ListAttachedRolePolicies(arg0 *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttachedRolePolicies", arg0)
	ret0, _ := ret[0].(*iam.ListAttachedRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttachedRolePolicies indicates an expected call of ListAttachedRolePolicies.
func (mr *MockClientMockRecorder) ListAttachedRolePolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedRolePolicies", reflect.TypeOf((*MockClient)(nil).ListAttachedRolePolicies), arg0)
}

// CreateRole mocks base method.
func (m *MockClient) CreateRole(arg0 *iam.CreateRoleInput) (*iam.CreateRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", arg0)
	ret0, _ := ret[0].(*iam.CreateRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockClientMockRecorder) CreateRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockClient)(nil).CreateRole), arg0)
}

// GetRole mocks base method.
func (m *MockClient) GetRole(arg0 *iam.GetRoleInput) (*iam.GetRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0)
	ret0, _ := ret[0].(*iam.GetRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockClientMockRecorder) GetRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockClient)(nil).GetRole), arg0)
}

// DeleteRole mocks base method.
func (m *MockClient) DeleteRole(arg0 *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0)
	ret0, _ := ret[0].(*iam.DeleteRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockClientMockRecorder) DeleteRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockClient)(nil).DeleteRole), arg0)
}

// ListAccounts mocks base method.
func (m *MockClient) ListAccounts(arg0 *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0)
	ret0, _ := ret[0].(*organizations.ListAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockClientMockRecorder) ListAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockClient)(nil).ListAccounts), arg0)
}

// CreateAccount mocks base method.
func (m *MockClient) CreateAccount(arg0 *organizations.CreateAccountInput) (*organizations.CreateAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(*organizations.CreateAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockClientMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockClient)(nil).CreateAccount), arg0)
}

// DescribeCreateAccountStatus mocks base method.
func (m *MockClient) DescribeCreateAccountStatus(arg0 *organizations.DescribeCreateAccountStatusInput) (*organizations.DescribeCreateAccountStatusOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCreateAccountStatus", arg0)
	ret0, _ := ret[0].(*organizations.DescribeCreateAccountStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCreateAccountStatus indicates an expected call of DescribeCreateAccountStatus.
func (mr *MockClientMockRecorder) DescribeCreateAccountStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCreateAccountStatus", reflect.TypeOf((*MockClient)(nil).DescribeCreateAccountStatus), arg0)
}

// AssumeRole mocks base method.
func (m *MockClient) AssumeRole(arg0 *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRole", arg0)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRole indicates an expected call of AssumeRole.
func (mr *MockClientMockRecorder) AssumeRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRole", reflect.TypeOf((*MockClient)(nil).AssumeRole), arg0)
}

// GetCallerIdentity mocks base method.
func (m *MockClient) GetCallerIdentity(arg0 *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentity", arg0)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity.
func (mr *MockClientMockRecorder) GetCallerIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockClient)(nil).GetCallerIdentity), arg0)
}

// GetFederationToken mocks base method.
func (m *MockClient) GetFederationToken(arg0 *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederationToken", arg0)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederationToken indicates an expected call of GetFederationToken.
func (mr *MockClientMockRecorder) GetFederationToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationToken", reflect.TypeOf((*MockClient)(nil).GetFederationToken), arg0)
}

// CreateCase mocks base method.
func (m *MockClient) CreateCase(arg0 *support.CreateCaseInput) (*support.CreateCaseOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCase", arg0)
	ret0, _ := ret[0].(*support.CreateCaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCase indicates an expected call of CreateCase.
func (mr *MockClientMockRecorder) CreateCase(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCase", reflect.TypeOf((*MockClient)(nil).CreateCase), arg0)
}

// DescribeCases mocks base method.
func (m *MockClient) DescribeCases(arg0 *support.DescribeCasesInput) (*support.DescribeCasesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCases", arg0)
	ret0, _ := ret[0].(*support.DescribeCasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCases indicates an expected call of DescribeCases.
func (mr *MockClientMockRecorder) DescribeCases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCases", reflect.TypeOf((*MockClient)(nil).DescribeCases), arg0)
}

// ListBuckets mocks base method.
func (m *MockClient) ListBuckets(arg0 *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuckets", arg0)
	ret0, _ := ret[0].(*s3.ListBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets.
func (mr *MockClientMockRecorder) ListBuckets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockClient)(nil).ListBuckets), arg0)
}

// DeleteBucket mocks base method.
func (m *MockClient) DeleteBucket(arg0 *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0)
	ret0, _ := ret[0].(*s3.DeleteBucketOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockClientMockRecorder) DeleteBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockClient)(nil).DeleteBucket), arg0)
}

// BatchDeleteBucketObjects mocks base method.
func (m *MockClient) BatchDeleteBucketObjects(bucketName *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDeleteBucketObjects", bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchDeleteBucketObjects indicates an expected call of BatchDeleteBucketObjects.
func (mr *MockClientMockRecorder) BatchDeleteBucketObjects(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDeleteBucketObjects", reflect.TypeOf((*MockClient)(nil).BatchDeleteBucketObjects), bucketName)
}

// ListObjectsV2 mocks base method.
func (m *MockClient) ListObjectsV2(arg0 *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectsV2", arg0)
	ret0, _ := ret[0].(*s3.ListObjectsV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectsV2 indicates an expected call of ListObjectsV2.
func (mr *MockClientMockRecorder) ListObjectsV2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsV2", reflect.TypeOf((*MockClient)(nil).ListObjectsV2), arg0)
}

// ListHostedZones mocks base method.
func (m *MockClient) ListHostedZones(arg0 *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostedZones", arg0)
	ret0, _ := ret[0].(*route53.ListHostedZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZones indicates an expected call of ListHostedZones.
func (mr *MockClientMockRecorder) ListHostedZones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZones", reflect.TypeOf((*MockClient)(nil).ListHostedZones), arg0)
}

// DeleteHostedZone mocks base method.
func (m *MockClient) DeleteHostedZone(arg0 *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHostedZone", arg0)
	ret0, _ := ret[0].(*route53.DeleteHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteHostedZone indicates an expected call of DeleteHostedZone.
func (mr *MockClientMockRecorder) DeleteHostedZone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHostedZone", reflect.TypeOf((*MockClient)(nil).DeleteHostedZone), arg0)
}

// ListResourceRecordSets mocks base method.
func (m *MockClient) ListResourceRecordSets(arg0 *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRecordSets", arg0)
	ret0, _ := ret[0].(*route53.ListResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceRecordSets indicates an expected call of ListResourceRecordSets.
func (mr *MockClientMockRecorder) ListResourceRecordSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRecordSets", reflect.TypeOf((*MockClient)(nil).ListResourceRecordSets), arg0)
}

// ChangeResourceRecordSets mocks base method.
func (m *MockClient) ChangeResourceRecordSets(arg0 *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeResourceRecordSets", arg0)
	ret0, _ := ret[0].(*route53.ChangeResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeResourceRecordSets indicates an expected call of ChangeResourceRecordSets.
func (mr *MockClientMockRecorder) ChangeResourceRecordSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeResourceRecordSets", reflect.TypeOf((*MockClient)(nil).ChangeResourceRecordSets), arg0)
}
