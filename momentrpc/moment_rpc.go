// Code generated by goctl. DO NOT EDIT.
// Source: moment.proto

package momentrpc

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateMomentReq    = pb.CreateMomentReq
	CreateMomentResp   = pb.CreateMomentResp
	DeleteMomentReq    = pb.DeleteMomentReq
	DeleteMomentResp   = pb.DeleteMomentResp
	ListMomentReq      = pb.ListMomentReq
	ListMomentResp     = pb.ListMomentResp
	Moment             = pb.Moment
	RetrieveMomentReq  = pb.RetrieveMomentReq
	RetrieveMomentResp = pb.RetrieveMomentResp
	UpdateMomentReq    = pb.UpdateMomentReq
	UpdateMomentResp   = pb.UpdateMomentResp

	MomentRpc interface {
		ListMoment(ctx context.Context, in *ListMomentReq, opts ...grpc.CallOption) (*ListMomentResp, error)
		RetrieveMoment(ctx context.Context, in *RetrieveMomentReq, opts ...grpc.CallOption) (*RetrieveMomentResp, error)
		CreateMoment(ctx context.Context, in *CreateMomentReq, opts ...grpc.CallOption) (*CreateMomentResp, error)
		UpdateMoment(ctx context.Context, in *UpdateMomentReq, opts ...grpc.CallOption) (*UpdateMomentResp, error)
		DeleteMoment(ctx context.Context, in *DeleteMomentReq, opts ...grpc.CallOption) (*DeleteMomentResp, error)
	}

	defaultMomentRpc struct {
		cli zrpc.Client
	}
)

func NewMomentRpc(cli zrpc.Client) MomentRpc {
	return &defaultMomentRpc{
		cli: cli,
	}
}

func (m *defaultMomentRpc) ListMoment(ctx context.Context, in *ListMomentReq, opts ...grpc.CallOption) (*ListMomentResp, error) {
	client := pb.NewMomentRpcClient(m.cli.Conn())
	return client.ListMoment(ctx, in, opts...)
}

func (m *defaultMomentRpc) RetrieveMoment(ctx context.Context, in *RetrieveMomentReq, opts ...grpc.CallOption) (*RetrieveMomentResp, error) {
	client := pb.NewMomentRpcClient(m.cli.Conn())
	return client.RetrieveMoment(ctx, in, opts...)
}

func (m *defaultMomentRpc) CreateMoment(ctx context.Context, in *CreateMomentReq, opts ...grpc.CallOption) (*CreateMomentResp, error) {
	client := pb.NewMomentRpcClient(m.cli.Conn())
	return client.CreateMoment(ctx, in, opts...)
}

func (m *defaultMomentRpc) UpdateMoment(ctx context.Context, in *UpdateMomentReq, opts ...grpc.CallOption) (*UpdateMomentResp, error) {
	client := pb.NewMomentRpcClient(m.cli.Conn())
	return client.UpdateMoment(ctx, in, opts...)
}

func (m *defaultMomentRpc) DeleteMoment(ctx context.Context, in *DeleteMomentReq, opts ...grpc.CallOption) (*DeleteMomentResp, error) {
	client := pb.NewMomentRpcClient(m.cli.Conn())
	return client.DeleteMoment(ctx, in, opts...)
}
