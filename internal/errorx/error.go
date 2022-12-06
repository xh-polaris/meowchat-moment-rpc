package errorx

import "google.golang.org/grpc/status"

var (
	ErrNoSuchMoment = status.New(10001, "no such moment")
)
