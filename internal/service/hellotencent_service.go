package service

import (
	logger "github.com/sirupsen/logrus"

	"HelloTencent/api"
	"HelloTencent/internal/utils"
	"context"
)

type HelloTencentService struct {}

// HelloTencent Stub
func (service *HelloTencentService) HelloTencent(ctx context.Context, request *api.HelloTencentRequest) (*api.HelloTencentResponse, error) {
	event := utils.GetThreadSafeEventData()

	event.SetOperation("HelloTencent")

	utils.StartEventData(event)
	defer utils.EndEventData(event)


	logger.Infof("Received HelloTencent request...")

	return &api.HelloTencentResponse{}, nil
}
