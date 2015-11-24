// Code generated by protoc-gen-grpc-gateway
// source: protos/slackinvite.proto
// DO NOT EDIT!

/*
Package slackinvite is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package slackinvite

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal
var _ = utilities.PascalFromSnake

var (
	filter_Slack_Stats_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_Slack_Stats_0(ctx context.Context, client SlackClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq Request

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_Slack_Stats_0); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Stats(ctx, &protoReq)
}

var (
	filter_Slack_Invite_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_Slack_Invite_0(ctx context.Context, client SlackClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq Request

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_Slack_Invite_0); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Invite(ctx, &protoReq)
}

// RegisterSlackHandlerFromEndpoint is same as RegisterSlackHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSlackHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string) (err error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSlackHandler(ctx, mux, conn)
}

// RegisterSlackHandler registers the http handlers for service Slack to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSlackHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewSlackClient(conn)

	mux.Handle("GET", pattern_Slack_Stats_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_Slack_Stats_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_Slack_Stats_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Slack_Invite_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_Slack_Invite_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_Slack_Invite_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Slack_Stats_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "slack", "stats"}, ""))

	pattern_Slack_Invite_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "slack", "invite"}, ""))
)

var (
	forward_Slack_Stats_0 = runtime.ForwardResponseMessage

	forward_Slack_Invite_0 = runtime.ForwardResponseMessage
)
