package errorx

import "google.golang.org/grpc/status"

var (
	ErrNoSuchMoment    = status.Error(10001, "no such moment")
	ErrInvalidObjectId = status.Error(10002, "invalid objectId")
)
