// Code generated by goctl. DO NOT EDIT.
// Source: moment.proto

package server

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/logic"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"
)

type MomentRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMomentRpcServer
}

func NewMomentRpcServer(svcCtx *svc.ServiceContext) *MomentRpcServer {
	return &MomentRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *MomentRpcServer) SearchMoment(ctx context.Context, in *pb.SearchMomentReq) (*pb.SearchMomentResp, error) {
	l := logic.NewSearchMomentLogic(ctx, s.svcCtx)
	return l.SearchMoment(in)
}

func (s *MomentRpcServer) ListMoment(ctx context.Context, in *pb.ListMomentReq) (*pb.ListMomentResp, error) {
	l := logic.NewListMomentLogic(ctx, s.svcCtx)
	return l.ListMoment(in)
}

func (s *MomentRpcServer) RetrieveMoment(ctx context.Context, in *pb.RetrieveMomentReq) (*pb.RetrieveMomentResp, error) {
	l := logic.NewRetrieveMomentLogic(ctx, s.svcCtx)
	return l.RetrieveMoment(in)
}

func (s *MomentRpcServer) CreateMoment(ctx context.Context, in *pb.CreateMomentReq) (*pb.CreateMomentResp, error) {
	l := logic.NewCreateMomentLogic(ctx, s.svcCtx)
	return l.CreateMoment(in)
}

func (s *MomentRpcServer) UpdateMoment(ctx context.Context, in *pb.UpdateMomentReq) (*pb.UpdateMomentResp, error) {
	l := logic.NewUpdateMomentLogic(ctx, s.svcCtx)
	return l.UpdateMoment(in)
}

func (s *MomentRpcServer) DeleteMoment(ctx context.Context, in *pb.DeleteMomentReq) (*pb.DeleteMomentResp, error) {
	l := logic.NewDeleteMomentLogic(ctx, s.svcCtx)
	return l.DeleteMoment(in)
}
