package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/model"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMomentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMomentLogic {
	return &CreateMomentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMomentLogic) CreateMoment(in *pb.CreateMomentReq) (*pb.CreateMomentResp, error) {
	m := in.Moment
	data := &model.Moment{
		Photos: m.Photos,
		Title:  m.Title,
		Text:   m.Text,
	}
	userId, err := primitive.ObjectIDFromHex(m.UserId)
	if err != nil {
		return nil, err
	}
	data.UserId = userId
	catId, err := primitive.ObjectIDFromHex(m.CatId)
	if err != nil {
		data.CatId = primitive.NilObjectID
	}
	data.CatId = catId
	err = l.svcCtx.MomentModel.Insert(context.Background(), data)
	if err != nil {
		return nil, err
	}

	return &pb.CreateMomentResp{MomentId: data.ID.Hex()}, nil
}
