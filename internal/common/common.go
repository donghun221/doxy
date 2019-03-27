package common

import (
	logger "github.com/sirupsen/logrus"
	constants "HelloTencent/internal/utils"
	"net"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	"time"
)

// Parse Remote address based on context
// If any error occurs, Empty string will be returned with error
func GetRemoteAddr(ctx context.Context) (string, error) {
	peer, ok := peer.FromContext(ctx)

	if !ok {
		errMsg := "Failed to parse remote address from context"
		logger.Error(errMsg)
		return constants.EMPTY_STRING, errors.New(errMsg)
	}

	if peer.Addr == net.Addr(nil) {
		errMsg := "Failed to get remote address from context, peer address equals to nil"
		logger.Error(errMsg)
		return constants.EMPTY_STRING, errors.New(errMsg)
	}

	return peer.Addr.String(), nil
}

func GetDurationMS(startTime time.Time) float64 {
	durationNs := time.Since(startTime)

	ms := durationNs / time.Millisecond
	return float64(ms)
}