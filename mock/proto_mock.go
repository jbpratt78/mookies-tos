// Code generated by MockGen. DO NOT EDIT.
// Source: protofiles/mookies.pb.go

// Package mock_mookiespb is a generated GoMock package.
package mock_mookiespb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	protofiles "github.com/jbpratt78/tos/protofiles"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockMenuServiceClient is a mock of MenuServiceClient interface
type MockMenuServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockMenuServiceClientMockRecorder
}

// MockMenuServiceClientMockRecorder is the mock recorder for MockMenuServiceClient
type MockMenuServiceClientMockRecorder struct {
	mock *MockMenuServiceClient
}

// NewMockMenuServiceClient creates a new mock instance
func NewMockMenuServiceClient(ctrl *gomock.Controller) *MockMenuServiceClient {
	mock := &MockMenuServiceClient{ctrl: ctrl}
	mock.recorder = &MockMenuServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMenuServiceClient) EXPECT() *MockMenuServiceClientMockRecorder {
	return m.recorder
}

// GetMenu mocks base method
func (m *MockMenuServiceClient) GetMenu(ctx context.Context, in *protofiles.Empty, opts ...grpc.CallOption) (*protofiles.Menu, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMenu", varargs...)
	ret0, _ := ret[0].(*protofiles.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu
func (mr *MockMenuServiceClientMockRecorder) GetMenu(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockMenuServiceClient)(nil).GetMenu), varargs...)
}

// CreateMenuItem mocks base method
func (m *MockMenuServiceClient) CreateMenuItem(ctx context.Context, in *protofiles.Item, opts ...grpc.CallOption) (*protofiles.CreateMenuItemResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMenuItem", varargs...)
	ret0, _ := ret[0].(*protofiles.CreateMenuItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMenuItem indicates an expected call of CreateMenuItem
func (mr *MockMenuServiceClientMockRecorder) CreateMenuItem(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenuItem", reflect.TypeOf((*MockMenuServiceClient)(nil).CreateMenuItem), varargs...)
}

// UpdateMenuItem mocks base method
func (m *MockMenuServiceClient) UpdateMenuItem(ctx context.Context, in *protofiles.Item, opts ...grpc.CallOption) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMenuItem", varargs...)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMenuItem indicates an expected call of UpdateMenuItem
func (mr *MockMenuServiceClientMockRecorder) UpdateMenuItem(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMenuItem", reflect.TypeOf((*MockMenuServiceClient)(nil).UpdateMenuItem), varargs...)
}

// DeleteMenuItem mocks base method
func (m *MockMenuServiceClient) DeleteMenuItem(ctx context.Context, in *protofiles.DeleteMenuItemRequest, opts ...grpc.CallOption) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMenuItem", varargs...)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMenuItem indicates an expected call of DeleteMenuItem
func (mr *MockMenuServiceClientMockRecorder) DeleteMenuItem(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMenuItem", reflect.TypeOf((*MockMenuServiceClient)(nil).DeleteMenuItem), varargs...)
}

// CreateMenuItemOption mocks base method
func (m *MockMenuServiceClient) CreateMenuItemOption(ctx context.Context, in *protofiles.Option, opts ...grpc.CallOption) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMenuItemOption", varargs...)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMenuItemOption indicates an expected call of CreateMenuItemOption
func (mr *MockMenuServiceClientMockRecorder) CreateMenuItemOption(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenuItemOption", reflect.TypeOf((*MockMenuServiceClient)(nil).CreateMenuItemOption), varargs...)
}

// MockMenuServiceServer is a mock of MenuServiceServer interface
type MockMenuServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockMenuServiceServerMockRecorder
}

// MockMenuServiceServerMockRecorder is the mock recorder for MockMenuServiceServer
type MockMenuServiceServerMockRecorder struct {
	mock *MockMenuServiceServer
}

// NewMockMenuServiceServer creates a new mock instance
func NewMockMenuServiceServer(ctrl *gomock.Controller) *MockMenuServiceServer {
	mock := &MockMenuServiceServer{ctrl: ctrl}
	mock.recorder = &MockMenuServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMenuServiceServer) EXPECT() *MockMenuServiceServerMockRecorder {
	return m.recorder
}

// GetMenu mocks base method
func (m *MockMenuServiceServer) GetMenu(arg0 context.Context, arg1 *protofiles.Empty) (*protofiles.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenu", arg0, arg1)
	ret0, _ := ret[0].(*protofiles.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu
func (mr *MockMenuServiceServerMockRecorder) GetMenu(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockMenuServiceServer)(nil).GetMenu), arg0, arg1)
}

// CreateMenuItem mocks base method
func (m *MockMenuServiceServer) CreateMenuItem(arg0 context.Context, arg1 *protofiles.Item) (*protofiles.CreateMenuItemResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMenuItem", arg0, arg1)
	ret0, _ := ret[0].(*protofiles.CreateMenuItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMenuItem indicates an expected call of CreateMenuItem
func (mr *MockMenuServiceServerMockRecorder) CreateMenuItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenuItem", reflect.TypeOf((*MockMenuServiceServer)(nil).CreateMenuItem), arg0, arg1)
}

// UpdateMenuItem mocks base method
func (m *MockMenuServiceServer) UpdateMenuItem(arg0 context.Context, arg1 *protofiles.Item) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMenuItem", arg0, arg1)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMenuItem indicates an expected call of UpdateMenuItem
func (mr *MockMenuServiceServerMockRecorder) UpdateMenuItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMenuItem", reflect.TypeOf((*MockMenuServiceServer)(nil).UpdateMenuItem), arg0, arg1)
}

// DeleteMenuItem mocks base method
func (m *MockMenuServiceServer) DeleteMenuItem(arg0 context.Context, arg1 *protofiles.DeleteMenuItemRequest) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMenuItem", arg0, arg1)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMenuItem indicates an expected call of DeleteMenuItem
func (mr *MockMenuServiceServerMockRecorder) DeleteMenuItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMenuItem", reflect.TypeOf((*MockMenuServiceServer)(nil).DeleteMenuItem), arg0, arg1)
}

// CreateMenuItemOption mocks base method
func (m *MockMenuServiceServer) CreateMenuItemOption(arg0 context.Context, arg1 *protofiles.Option) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMenuItemOption", arg0, arg1)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMenuItemOption indicates an expected call of CreateMenuItemOption
func (mr *MockMenuServiceServerMockRecorder) CreateMenuItemOption(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenuItemOption", reflect.TypeOf((*MockMenuServiceServer)(nil).CreateMenuItemOption), arg0, arg1)
}

// MockOrderServiceClient is a mock of OrderServiceClient interface
type MockOrderServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceClientMockRecorder
}

// MockOrderServiceClientMockRecorder is the mock recorder for MockOrderServiceClient
type MockOrderServiceClientMockRecorder struct {
	mock *MockOrderServiceClient
}

// NewMockOrderServiceClient creates a new mock instance
func NewMockOrderServiceClient(ctrl *gomock.Controller) *MockOrderServiceClient {
	mock := &MockOrderServiceClient{ctrl: ctrl}
	mock.recorder = &MockOrderServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderServiceClient) EXPECT() *MockOrderServiceClientMockRecorder {
	return m.recorder
}

// SubmitOrder mocks base method
func (m *MockOrderServiceClient) SubmitOrder(ctx context.Context, in *protofiles.Order, opts ...grpc.CallOption) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitOrder", varargs...)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrder indicates an expected call of SubmitOrder
func (mr *MockOrderServiceClientMockRecorder) SubmitOrder(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockOrderServiceClient)(nil).SubmitOrder), varargs...)
}

// ActiveOrders mocks base method
func (m *MockOrderServiceClient) ActiveOrders(ctx context.Context, in *protofiles.Empty, opts ...grpc.CallOption) (*protofiles.OrdersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ActiveOrders", varargs...)
	ret0, _ := ret[0].(*protofiles.OrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveOrders indicates an expected call of ActiveOrders
func (mr *MockOrderServiceClientMockRecorder) ActiveOrders(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveOrders", reflect.TypeOf((*MockOrderServiceClient)(nil).ActiveOrders), varargs...)
}

// CompleteOrder mocks base method
func (m *MockOrderServiceClient) CompleteOrder(ctx context.Context, in *protofiles.CompleteOrderRequest, opts ...grpc.CallOption) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteOrder", varargs...)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteOrder indicates an expected call of CompleteOrder
func (mr *MockOrderServiceClientMockRecorder) CompleteOrder(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteOrder", reflect.TypeOf((*MockOrderServiceClient)(nil).CompleteOrder), varargs...)
}

// SubscribeToOrders mocks base method
func (m *MockOrderServiceClient) SubscribeToOrders(ctx context.Context, in *protofiles.Empty, opts ...grpc.CallOption) (protofiles.OrderService_SubscribeToOrdersClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeToOrders", varargs...)
	ret0, _ := ret[0].(protofiles.OrderService_SubscribeToOrdersClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToOrders indicates an expected call of SubscribeToOrders
func (mr *MockOrderServiceClientMockRecorder) SubscribeToOrders(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToOrders", reflect.TypeOf((*MockOrderServiceClient)(nil).SubscribeToOrders), varargs...)
}

// MockOrderService_SubscribeToOrdersClient is a mock of OrderService_SubscribeToOrdersClient interface
type MockOrderService_SubscribeToOrdersClient struct {
	ctrl     *gomock.Controller
	recorder *MockOrderService_SubscribeToOrdersClientMockRecorder
}

// MockOrderService_SubscribeToOrdersClientMockRecorder is the mock recorder for MockOrderService_SubscribeToOrdersClient
type MockOrderService_SubscribeToOrdersClientMockRecorder struct {
	mock *MockOrderService_SubscribeToOrdersClient
}

// NewMockOrderService_SubscribeToOrdersClient creates a new mock instance
func NewMockOrderService_SubscribeToOrdersClient(ctrl *gomock.Controller) *MockOrderService_SubscribeToOrdersClient {
	mock := &MockOrderService_SubscribeToOrdersClient{ctrl: ctrl}
	mock.recorder = &MockOrderService_SubscribeToOrdersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderService_SubscribeToOrdersClient) EXPECT() *MockOrderService_SubscribeToOrdersClientMockRecorder {
	return m.recorder
}

// Recv mocks base method
func (m *MockOrderService_SubscribeToOrdersClient) Recv() (*protofiles.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*protofiles.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockOrderService_SubscribeToOrdersClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockOrderService_SubscribeToOrdersClient)(nil).Recv))
}

// Header mocks base method
func (m *MockOrderService_SubscribeToOrdersClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockOrderService_SubscribeToOrdersClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockOrderService_SubscribeToOrdersClient)(nil).Header))
}

// Trailer mocks base method
func (m *MockOrderService_SubscribeToOrdersClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockOrderService_SubscribeToOrdersClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockOrderService_SubscribeToOrdersClient)(nil).Trailer))
}

// CloseSend mocks base method
func (m *MockOrderService_SubscribeToOrdersClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockOrderService_SubscribeToOrdersClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockOrderService_SubscribeToOrdersClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockOrderService_SubscribeToOrdersClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockOrderService_SubscribeToOrdersClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOrderService_SubscribeToOrdersClient)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockOrderService_SubscribeToOrdersClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockOrderService_SubscribeToOrdersClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOrderService_SubscribeToOrdersClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockOrderService_SubscribeToOrdersClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockOrderService_SubscribeToOrdersClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOrderService_SubscribeToOrdersClient)(nil).RecvMsg), m)
}

// MockOrderServiceServer is a mock of OrderServiceServer interface
type MockOrderServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceServerMockRecorder
}

// MockOrderServiceServerMockRecorder is the mock recorder for MockOrderServiceServer
type MockOrderServiceServerMockRecorder struct {
	mock *MockOrderServiceServer
}

// NewMockOrderServiceServer creates a new mock instance
func NewMockOrderServiceServer(ctrl *gomock.Controller) *MockOrderServiceServer {
	mock := &MockOrderServiceServer{ctrl: ctrl}
	mock.recorder = &MockOrderServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderServiceServer) EXPECT() *MockOrderServiceServerMockRecorder {
	return m.recorder
}

// SubmitOrder mocks base method
func (m *MockOrderServiceServer) SubmitOrder(arg0 context.Context, arg1 *protofiles.Order) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitOrder", arg0, arg1)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrder indicates an expected call of SubmitOrder
func (mr *MockOrderServiceServerMockRecorder) SubmitOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockOrderServiceServer)(nil).SubmitOrder), arg0, arg1)
}

// ActiveOrders mocks base method
func (m *MockOrderServiceServer) ActiveOrders(arg0 context.Context, arg1 *protofiles.Empty) (*protofiles.OrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveOrders", arg0, arg1)
	ret0, _ := ret[0].(*protofiles.OrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveOrders indicates an expected call of ActiveOrders
func (mr *MockOrderServiceServerMockRecorder) ActiveOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveOrders", reflect.TypeOf((*MockOrderServiceServer)(nil).ActiveOrders), arg0, arg1)
}

// CompleteOrder mocks base method
func (m *MockOrderServiceServer) CompleteOrder(arg0 context.Context, arg1 *protofiles.CompleteOrderRequest) (*protofiles.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteOrder", arg0, arg1)
	ret0, _ := ret[0].(*protofiles.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteOrder indicates an expected call of CompleteOrder
func (mr *MockOrderServiceServerMockRecorder) CompleteOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteOrder", reflect.TypeOf((*MockOrderServiceServer)(nil).CompleteOrder), arg0, arg1)
}

// SubscribeToOrders mocks base method
func (m *MockOrderServiceServer) SubscribeToOrders(arg0 *protofiles.Empty, arg1 protofiles.OrderService_SubscribeToOrdersServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToOrders", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeToOrders indicates an expected call of SubscribeToOrders
func (mr *MockOrderServiceServerMockRecorder) SubscribeToOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToOrders", reflect.TypeOf((*MockOrderServiceServer)(nil).SubscribeToOrders), arg0, arg1)
}

// MockOrderService_SubscribeToOrdersServer is a mock of OrderService_SubscribeToOrdersServer interface
type MockOrderService_SubscribeToOrdersServer struct {
	ctrl     *gomock.Controller
	recorder *MockOrderService_SubscribeToOrdersServerMockRecorder
}

// MockOrderService_SubscribeToOrdersServerMockRecorder is the mock recorder for MockOrderService_SubscribeToOrdersServer
type MockOrderService_SubscribeToOrdersServerMockRecorder struct {
	mock *MockOrderService_SubscribeToOrdersServer
}

// NewMockOrderService_SubscribeToOrdersServer creates a new mock instance
func NewMockOrderService_SubscribeToOrdersServer(ctrl *gomock.Controller) *MockOrderService_SubscribeToOrdersServer {
	mock := &MockOrderService_SubscribeToOrdersServer{ctrl: ctrl}
	mock.recorder = &MockOrderService_SubscribeToOrdersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderService_SubscribeToOrdersServer) EXPECT() *MockOrderService_SubscribeToOrdersServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockOrderService_SubscribeToOrdersServer) Send(arg0 *protofiles.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockOrderService_SubscribeToOrdersServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOrderService_SubscribeToOrdersServer)(nil).Send), arg0)
}

// SetHeader mocks base method
func (m *MockOrderService_SubscribeToOrdersServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockOrderService_SubscribeToOrdersServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockOrderService_SubscribeToOrdersServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method
func (m *MockOrderService_SubscribeToOrdersServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockOrderService_SubscribeToOrdersServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockOrderService_SubscribeToOrdersServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockOrderService_SubscribeToOrdersServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockOrderService_SubscribeToOrdersServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockOrderService_SubscribeToOrdersServer)(nil).SetTrailer), arg0)
}

// Context mocks base method
func (m *MockOrderService_SubscribeToOrdersServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockOrderService_SubscribeToOrdersServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOrderService_SubscribeToOrdersServer)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockOrderService_SubscribeToOrdersServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockOrderService_SubscribeToOrdersServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOrderService_SubscribeToOrdersServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockOrderService_SubscribeToOrdersServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockOrderService_SubscribeToOrdersServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOrderService_SubscribeToOrdersServer)(nil).RecvMsg), m)
}
