package model

import (
	"github.com/xh-polaris/meowchat-moment-rpc/errorx"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errorx.ErrInvalidObjectId
)
