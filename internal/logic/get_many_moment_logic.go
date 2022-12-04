package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetManyMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetManyMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetManyMomentLogic {
	return &GetManyMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetManyMomentLogic) GetManyMoment(in *pb.GetManyMomentReq) (*pb.GetManyMomentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetManyMomentResp{}, nil
}
