package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Moment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CatId       primitive.ObjectID `bson:"catId,omitempty"`
	CommunityId primitive.ObjectID `bson:"communityId"`
	Photos      []string           `bson:"photos"`
	Title       string             `bson:"title"`
	Text        string             `bson:"text"`
	UserId      primitive.ObjectID `bson:"userId"`
	IsDeleted   bool               `bson:"isDeleted"`
	DeleteAt    time.Time          `bson:"deleteAt,omitempty" json:"deleteAt,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
