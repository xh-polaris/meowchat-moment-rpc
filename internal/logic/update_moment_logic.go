package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-moment-rpc/errorx"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/model"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	m := in.Moment
	momentId, err := primitive.ObjectIDFromHex(m.Id)
	if err != nil {
		return nil, errorx.ErrInvalidObjectId
	}
	userId, err := primitive.ObjectIDFromHex(m.UserId)
	if err != nil {
		return nil, errorx.ErrInvalidObjectId
	}
	communityId, err := primitive.ObjectIDFromHex(m.CommunityId)
	if err != nil {
		return nil, errorx.ErrInvalidObjectId
	}
	catId, err := primitive.ObjectIDFromHex(m.CatId)
	if err != nil {
		catId = primitive.NilObjectID
	}

	err = l.svcCtx.MomentModel.UpdateValid(l.ctx, &model.Moment{
		ID:          momentId,
		CatId:       catId,
		CommunityId: communityId,
		Photos:      m.Photos,
		Title:       m.Title,
		Text:        m.Text,
		UserId:      userId,
		IsDeleted:   false,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateMomentResp{}, nil
}
