package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMomentLogic {
	return &UpdateMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMomentLogic) UpdateMoment(in *pb.UpdateMomentReq) (*pb.UpdateMomentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMomentResp{}, nil
}
