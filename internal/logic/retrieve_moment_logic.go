package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveMomentLogic {
	return &RetrieveMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetrieveMomentLogic) RetrieveMoment(in *pb.RetrieveMomentReq) (*pb.RetrieveMomentResp, error) {
	data, err := l.svcCtx.MomentModel.FindOneValid(l.ctx, in.MomentId)
	if err != nil {
		return nil, err
	}
	m := &pb.Moment{
		Id:          data.ID.Hex(),
		CreateAt:    data.CreateAt.Unix(),
		Photos:      data.Photos,
		Title:       data.Title,
		Text:        data.Text,
		UserId:      data.UserId.Hex(),
		CommunityId: data.CommunityId.Hex(),
	}
	if !data.CatId.IsZero() {
		m.CatId = data.CatId.Hex()
	}
	return &pb.RetrieveMomentResp{Moment: m}, nil
}
