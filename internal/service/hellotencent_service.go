package service

import (
	logger "github.com/sirupsen/logrus"

	"HelloTencent/api"
	"context"
)

type HelloTencentService struct {}

// HelloTencent Stub
func (service *HelloTencentService) HelloTencent(ctx context.Context, request *api.HelloTencentRequest) (*api.HelloTencentResponse, error) {
	logger.Infof("Received HelloTencent request...")

	// TODO: Implement your own logic

	return &api.HelloTencentResponse{}, nil
}
