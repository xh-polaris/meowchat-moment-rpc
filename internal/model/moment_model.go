package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const MomentCollectionName = "moment"

var _ MomentModel = (*customMomentModel)(nil)

type (
	// MomentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMomentModel.
	MomentModel interface {
		momentModel
		DeleteSoftly(ctx context.Context, id string) error
		FindOneValid(ctx context.Context, id string) (*Moment, error)
		FindManyValid(ctx context.Context, id string, count, skip int64) ([]*Moment, error)
		UpdateValid(ctx context.Context, data *Moment) error
	}

	customMomentModel struct {
		*defaultMomentModel
	}
)

// NewMomentModel returns a model for the mongo.
func NewMomentModel(url, db string, c cache.CacheConf) MomentModel {
	conn := monc.MustNewModel(url, db, MomentCollectionName, c)
	return &customMomentModel{
		defaultMomentModel: newDefaultMomentModel(conn),
	}
}

func (m *defaultMomentModel) DeleteSoftly(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}
	key := prefixMomentCacheKey + id
	_, err = m.conn.UpdateOne(ctx, key, bson.M{"_id": oid, "isDeleted": false}, bson.M{"$set": bson.M{
		"isDeleted": true,
		"deleteAt":  time.Now(),
	}})
	return err
}

func (m *defaultMomentModel) FindOneValid(ctx context.Context, id string) (*Moment, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}
	key := prefixMomentCacheKey + id
	var data Moment
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid, "isDeleted": false})
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (m *defaultMomentModel) FindManyValid(ctx context.Context, communityId string, count, skip int64) ([]*Moment, error) {
	oid, err := primitive.ObjectIDFromHex(communityId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}
	data := make([]*Moment, 0, 20)
	opt := &options.FindOptions{
		Skip:  &skip,
		Limit: &count,
	}
	err = m.conn.Find(ctx, &data, bson.M{"communityId": oid, "isDeleted": false}, opt)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (m *defaultMomentModel) UpdateValid(ctx context.Context, data *Moment) error {
	data.UpdateAt = time.Now()
	key := prefixMomentCacheKey + data.ID.Hex()
	_, err := m.conn.ReplaceOne(ctx, key, bson.M{"_id": data.ID, "isDeleted": false}, data)
	return err
}
