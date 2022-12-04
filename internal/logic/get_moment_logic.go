package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMomentLogic {
	return &GetMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMomentLogic) GetMoment(in *pb.GetMomentReq) (*pb.GetMomentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMomentResp{}, nil
}
