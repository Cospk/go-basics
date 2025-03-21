// Code generated by Kitex v0.9.1. DO NOT EDIT.

package math

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	math_service "my_kitex/kitex_gen/math_service"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Add": kitex.NewMethodInfo(
		addHandler,
		newAddArgs,
		newAddResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Sub": kitex.NewMethodInfo(
		subHandler,
		newSubArgs,
		newSubResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	mathServiceInfo                = NewServiceInfo()
	mathServiceInfoForClient       = NewServiceInfoForClient()
	mathServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return mathServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return mathServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return mathServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Math"
	handlerType := (*math_service.Math)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "idl",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(math_service.AddRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(math_service.Math).Add(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddArgs:
		success, err := handler.(math_service.Math).Add(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddArgs() interface{} {
	return &AddArgs{}
}

func newAddResult() interface{} {
	return &AddResult{}
}

type AddArgs struct {
	Req *math_service.AddRequest
}

func (p *AddArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(math_service.AddRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddArgs) Unmarshal(in []byte) error {
	msg := new(math_service.AddRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddArgs_Req_DEFAULT *math_service.AddRequest

func (p *AddArgs) GetReq() *math_service.AddRequest {
	if !p.IsSetReq() {
		return AddArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddResult struct {
	Success *math_service.AddResponse
}

var AddResult_Success_DEFAULT *math_service.AddResponse

func (p *AddResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(math_service.AddResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddResult) Unmarshal(in []byte) error {
	msg := new(math_service.AddResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddResult) GetSuccess() *math_service.AddResponse {
	if !p.IsSetSuccess() {
		return AddResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddResult) SetSuccess(x interface{}) {
	p.Success = x.(*math_service.AddResponse)
}

func (p *AddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddResult) GetResult() interface{} {
	return p.Success
}

func subHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(math_service.SubRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(math_service.Math).Sub(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SubArgs:
		success, err := handler.(math_service.Math).Sub(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SubResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSubArgs() interface{} {
	return &SubArgs{}
}

func newSubResult() interface{} {
	return &SubResult{}
}

type SubArgs struct {
	Req *math_service.SubRequest
}

func (p *SubArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(math_service.SubRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SubArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SubArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SubArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SubArgs) Unmarshal(in []byte) error {
	msg := new(math_service.SubRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SubArgs_Req_DEFAULT *math_service.SubRequest

func (p *SubArgs) GetReq() *math_service.SubRequest {
	if !p.IsSetReq() {
		return SubArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SubArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SubArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SubResult struct {
	Success *math_service.SubResponse
}

var SubResult_Success_DEFAULT *math_service.SubResponse

func (p *SubResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(math_service.SubResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SubResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SubResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SubResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SubResult) Unmarshal(in []byte) error {
	msg := new(math_service.SubResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SubResult) GetSuccess() *math_service.SubResponse {
	if !p.IsSetSuccess() {
		return SubResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SubResult) SetSuccess(x interface{}) {
	p.Success = x.(*math_service.SubResponse)
}

func (p *SubResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SubResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Add(ctx context.Context, Req *math_service.AddRequest) (r *math_service.AddResponse, err error) {
	var _args AddArgs
	_args.Req = Req
	var _result AddResult
	if err = p.c.Call(ctx, "Add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Sub(ctx context.Context, Req *math_service.SubRequest) (r *math_service.SubResponse, err error) {
	var _args SubArgs
	_args.Req = Req
	var _result SubResult
	if err = p.c.Call(ctx, "Sub", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
