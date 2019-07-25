package service

import (
	"HelloTencent/api"
	"HelloTencent/internal/common"
	"HelloTencent/internal/utils"
	constants "HelloTencent/internal/utils"
	"context"
	logger "github.com/sirupsen/logrus"
	"runtime"
	"strconv"
	"time"
)

type CommonService struct {}

// Invoking GC for GoLang.
func (service *CommonService) GC(ctx context.Context, request *api.GCRequest) (*api.GCResponse, error) {
	event := utils.GetThreadSafeEventData()

	event.SetOperation("GC")

	utils.StartEventData(event)
	defer utils.EndEventData(event)

	if request == nil {
		logger.Errorf("Request is missing, skip GC")
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.GCResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, skip GC")
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.GCResponse{}
		response.Header = header

		return &response, nil
	}

	remoteAddr, err := common.GetRemoteAddr(ctx)
	if err != nil {
		logger.Errorf("Unable to parse remote address, skip GC")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.GCResponse{}
		response.Header = header

		return &response, nil
	}
	event.SetRemoteAddr(remoteAddr)

	event.AddNameValuePair("RequestId", request.GetHeader().GetRequestId())
	event.AddNameValuePair("clientRequestTime", strconv.FormatUint(request.GetHeader().GetStartTime(), 10))

	basicLoggerFields := MakeBasicLoggerFields("GC", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received GC request...")

	// Invoke GC
	event.StartTimer("GC")
	runtime.GC()
	event.EndTimer("GC")

	response := api.GCResponse{}
	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(utils.GetCurrentTimeMillis())
	header.RetCode = api.RetCode_OK

	event.AppendNameValuePair("ResponseCode", api.RetCode_OK.String())

	response.Header = &header

	logger.WithFields(basicLoggerFields).Infof("Finished GC request...")

	return &response, nil
}

// Dump Config in memory.
func (service *CommonService) GetConfig(ctx context.Context, request *api.GetConfigRequest) (*api.GetConfigResponse, error) {
	event := utils.GetThreadSafeEventData()

	event.SetOperation("GC")

	utils.StartEventData(event)
	defer utils.EndEventData(event)

	remoteAddr, err := common.GetRemoteAddr(ctx)

	if err != nil {
		logger.Errorf("Unable to parse remote address, skip GetConfig")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.GetConfigResponse{}
		response.Header = header

		return &response, nil
	}


	if request == nil {
		logger.Errorf("Request is missing, remoteAddr:%s, skip GetConfig", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.GetConfigResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, remoteAddr:%s, skip GetConfig", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.GetConfigResponse{}
		response.Header = header

		return &response, nil
	}

	event.SetRemoteAddr(remoteAddr)

	event.AddNameValuePair("RequestId", request.GetHeader().GetRequestId())
	event.AddNameValuePair("clientRequestTime", strconv.FormatUint(request.GetHeader().GetStartTime(), 10))

	basicLoggerFields := MakeBasicLoggerFields("GetConfig", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received GC request...")

	configStr := "This is the config string"

	response := api.GetConfigResponse{}
	response.Message = configStr

	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(time.Now().UnixNano())
	header.RetCode = api.RetCode_OK

	event.AppendNameValuePair("ResponseCode", api.RetCode_OK.String())

	response.Header = &header

	logger.WithFields(basicLoggerFields).Infof("Finished GetConfig request...")

	return &response, nil
}

// By default, it will return "Healthy".
func (service *CommonService) IsHealthy(ctx context.Context, request *api.IsHealthyRequest) (*api.IsHealthyResponse, error) {
	event := utils.GetThreadSafeEventData()

	event.SetOperation("GC")

	utils.StartEventData(event)
	defer utils.EndEventData(event)

	remoteAddr, err := common.GetRemoteAddr(ctx)

	if err != nil {
		logger.Errorf("Unable to parse remote address, skip IsHealthy")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.IsHealthyResponse{}
		response.Header = header

		return &response, nil
	}


	if request == nil {
		logger.Errorf("Request is missing, remoteAddr:%s, skip IsHealthy", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.IsHealthyResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, remoteAddr:%s, skip IsHealthy", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		event.AddNameValuePair("ResponseCode", header.GetRetCode().String())

		response := api.IsHealthyResponse{}
		response.Header = header

		return &response, nil
	}

	event.SetRemoteAddr(remoteAddr)

	event.AddNameValuePair("RequestId", request.GetHeader().GetRequestId())
	event.AddNameValuePair("clientRequestTime", strconv.FormatUint(request.GetHeader().GetStartTime(), 10))

	basicLoggerFields := MakeBasicLoggerFields("IsHealthy", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received GC request...")

	isHealthyStr := "Healthy"

	response := api.IsHealthyResponse{}
	response.Message = isHealthyStr

	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(time.Now().UnixNano())
	header.RetCode = api.RetCode_OK

	event.AppendNameValuePair("ResponseCode", api.RetCode_OK.String())

	response.Header = &header

	logger.WithFields(basicLoggerFields).Infof("Finished IsHealthy request...")

	return &response, nil
}
// List Stubs.
func (service *CommonService) Help(ctx context.Context, request *api.HelpRequest) (*api.HelpResponse, error) {
	remoteAddr, err := common.GetRemoteAddr(ctx)

	if err != nil {
		logger.Errorf("Unable to parse remote address, skip Help")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		response := api.HelpResponse{}
		response.Header = header

		return &response, nil
	}


	if request == nil {
		logger.Errorf("Request is missing, remoteAddr:%s, skip Help", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		response := api.HelpResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, remoteAddr:%s, skip Help", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		response := api.HelpResponse{}
		response.Header = header

		return &response, nil
	}

	basicLoggerFields := MakeBasicLoggerFields("Help", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received GC request...")

	helpStr := "This is help response string"

	response := api.HelpResponse{}
	response.Message = helpStr

	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(time.Now().UnixNano())
	header.RetCode = api.RetCode_OK

	response.Header = &header

	logger.WithFields(basicLoggerFields).Infof("Finished Help request...")

	return &response, nil
}

// Return CPU, Memory, RPC related information.
func (service *CommonService) Info(ctx context.Context, request *api.InfoRequest) (*api.InfoResponse, error) {
	remoteAddr, err := common.GetRemoteAddr(ctx)

	if err != nil {
		logger.Errorf("Unable to parse remote address, skip Info")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		response := api.InfoResponse{}
		response.Header = header

		return &response, nil
	}


	if request == nil {
		logger.Errorf("Request is missing, remoteAddr:%s, skip Info", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		response := api.InfoResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, remoteAddr:%s, skip Info", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		response := api.InfoResponse{}
		response.Header = header

		return &response, nil
	}

	basicLoggerFields := MakeBasicLoggerFields("Info", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received Info request...")

	InfoStr := "This is info response string"

	response := api.InfoResponse{}
	response.Message = InfoStr

	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(time.Now().UnixNano())
	header.RetCode = api.RetCode_OK

	response.Header = &header

	logger.WithFields(basicLoggerFields).Infof("Finished Info request...")

	return &response, nil
}

// Set logging level.
func (service *CommonService) Log(ctx context.Context, request *api.LogRequest) (*api.LogResponse, error) {
	remoteAddr, err := common.GetRemoteAddr(ctx)

	if err != nil {
		logger.Errorf("Unable to parse remote address, skip Log")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		response := api.LogResponse{}
		response.Header = header

		return &response, nil
	}


	if request == nil {
		logger.Errorf("Request is missing, remoteAddr:%s, skip Log", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		response := api.LogResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, remoteAddr:%s, skip Log", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		response := api.LogResponse{}
		response.Header = header

		return &response, nil
	}

	basicLoggerFields := MakeBasicLoggerFields("Log", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received Log request...")

	logStr := "This is log response string"

	response := api.LogResponse{}
	response.Message = logStr

	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(time.Now().UnixNano())
	header.RetCode = api.RetCode_OK

	response.Header = &header

	logger.WithFields(basicLoggerFields).Infof("Finished Log request...")

	return &response, nil
}

// Set logging level.
func (service *CommonService) Ping(ctx context.Context, request *api.PingRequest) (*api.PingResponse, error) {
	remoteAddr, err := common.GetRemoteAddr(ctx)

	if err != nil {
		logger.Errorf("Unable to parse remote address, skip Ping")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		response := api.PingResponse{}
		response.Header = header

		return &response, nil
	}


	if request == nil {
		logger.Errorf("Request is missing, remoteAddr:%s, skip Ping", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		response := api.PingResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, remoteAddr:%s, skip Ping", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		response := api.PingResponse{}
		response.Header = header

		return &response, nil
	}

	basicLoggerFields := MakeBasicLoggerFields("Ping", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received Log request...")

	pingStr := "This is ping response string"

	response := api.PingResponse{}
	response.Message = pingStr

	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(time.Now().UnixNano())
	header.RetCode = api.RetCode_OK

	response.Header = &header

	logger.WithFields(basicLoggerFields).Infof("Finished Ping request...")

	return &response, nil
}

// Stop main process.
func (service *CommonService) Shutdown(ctx context.Context, request *api.ShutdownRequest) (*api.ShutdownResponse, error) {
	remoteAddr, err := common.GetRemoteAddr(ctx)

	if err != nil {
		logger.Errorf("Unable to parse remote address, skip Shutdown")
		header := MakeFailedResponseHeader(request.GetHeader().GetRequestId(), "Failed to parse remote address")

		response := api.ShutdownResponse{}
		response.Header = header

		return &response, nil
	}


	if request == nil {
		logger.Errorf("Request is missing, remoteAddr:%s, skip Shutdown", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request is missing")

		response := api.ShutdownResponse{}
		response.Header = header

		return &response, nil
	}

	// Validate basic required fields
	if request.GetHeader() == nil {
		logger.Errorf("Request Header is missing, remoteAddr:%s, skip Shutdown", remoteAddr)
		header := MakeFailedResponseHeader("Not Provided", "Request Header is missing")

		response := api.ShutdownResponse{}
		response.Header = header

		return &response, nil
	}

	basicLoggerFields := MakeBasicLoggerFields("Shutdown", request.GetHeader().RequestId, remoteAddr)
	logger.WithFields(basicLoggerFields).Infof("Received Shutdown request...")

	pingStr := "This is ping response string"

	response := api.ShutdownResponse{}
	response.Message = pingStr

	header := api.ResponseHeader{}

	header.RequestId = request.GetHeader().GetRequestId()
	header.RespTime = uint64(time.Now().UnixNano())
	header.RetCode = api.RetCode_OK

	response.Header = &header

	logger.WithFields(basicLoggerFields).Infof("Finished Shutdown request...")

	return &response, nil

}

func MakeBasicLoggerFields(operation, requestId, remoteAddr string) logger.Fields {
	res := logger.Fields{
		constants.OPERATION: operation,
		constants.REQUEST_ID: requestId,
		constants.REMOTE_ADDR: remoteAddr,
	}

	return res
}

func MakeFailedResponseHeader(requestId, errMsg string) *api.ResponseHeader {
	header := api.ResponseHeader{}
	header.RetCode = api.RetCode_FAIL
	header.RespTime = uint64(time.Now().UnixNano())
	header.RequestId = requestId
	header.ErrMsg = errMsg
	return &header
}