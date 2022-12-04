package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMomentLogic {
	return &AddMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMomentLogic) AddMoment(in *pb.AddMomentReq) (*pb.AddMomentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddMomentResp{}, nil
}
