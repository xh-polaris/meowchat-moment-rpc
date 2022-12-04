package errorx

import "google.golang.org/grpc/status"

var (
	ErrNoSuchMoment = status.New(1001, "no such moment")
)
