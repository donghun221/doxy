package service

import (
	logger "github.com/sirupsen/logrus"

	"HelloTencent/api"
	"HelloTencent/internal/common"
	"HelloTencent/internal/utils"
	"context"
	"strconv"
)

type HelloTencentService struct {}

// HelloTencent Stub
func (service *HelloTencentService) HelloTencent(ctx context.Context, request *api.HelloTencentRequest) (*api.HelloTencentResponse, error) {
	event := utils.GetThreadSafeEventData()

	event.SetOperation("HelloTencent")

	utils.StartEventData(event)
	defer utils.EndEventData(event)

	if request == nil {
		logger.Errorf("Request is missing, skip HelloTencent")
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.HelloTencentResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, skip HelloTencent")
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.HelloTencentResponse{}
		response.Header = header

		return &response, nil
	}

	remoteAddr, err := common.GetRemoteAddr(ctx)
	if err != nil {
		logger.Errorf("Unable to parse remote address, skip HelloTencent")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.HelloTencentResponse{}
		response.Header = header

		return &response, nil
	}
	event.SetRemoteAddr(remoteAddr)

	event.AddNameValuePair("RequestId", request.GetHeader().GetRequestId())
	event.AddNameValuePair("clientRequestTime", strconv.FormatUint(request.GetHeader().GetStartTime(), 10))

	basicLoggerFields := MakeBasicLoggerFields("HelloTencent", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received HelloTencent request...")

	// TODO: Implement your own logic
	event.StartTimer("HelloTencent")
	message := "Hello, " + request.GetMessage()
	event.EndTimer("HelloTencent")

	response := api.HelloTencentResponse{}
	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(utils.GetCurrentTimeMillis())
	header.RetCode = api.RetCode_OK

	event.AppendNameValuePair("ResponseCode", api.RetCode_OK.String())

	response.Header = &header
	response.Message = message

	logger.WithFields(basicLoggerFields).Infof("Finished HelloTencent request...")

	return &response, nil
}