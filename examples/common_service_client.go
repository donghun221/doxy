package main

import (
	"HelloTencent/api"
	constants "HelloTencent/internal/utils"
	"HelloTencent/internal/utils"
	"context"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"log"
	"os"

	//"os"
	"time"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(constants.LOCAL_HOST + constants.COLON + constants.PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := api.NewCommonServiceClient(conn)

	// Contact the server and print out its response.
 	index := ""
	if len(os.Args) > 1 {
		index = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uuid, _ := uuid.NewV4()

	header := api.RequestHeader{}
	header.RequestId = uuid.String()
	header.StartTime = utils.GetCurrentTimeMillis()

	switch index {
	case "GC":
		callGC(&ctx, client, &header)
	case "GetConfig":
		callGetConfig(&ctx, client, &header)
	case "IsHealthy":
		callIsHealthy(&ctx, client, &header)
	case "Help":
		callHelp(&ctx, client, &header)
	case "Info":
		callInfo(&ctx, client, &header)
	case "Log":
		callLog(&ctx, client, &header)
	case "Ping":
		callPing(&ctx, client, &header)
	case "Shutdown":
		callShutdown(&ctx, client, &header)
	default:
		callGC(&ctx, client, &header)
	}

}

func callGC(ctx *context.Context, client api.CommonServiceClient, header *api.RequestHeader) {
	request := api.GCRequest{}
	request.Header = header

	response, err := client.GC(*ctx, &request)

	if err != nil {
		log.Println("Failed to call GC %e...", err.Error())
		return
	}

	log.Println("Received GC Response...")

	log.Printf("\n header:[%s, %s, %d, %s] \n message:[%s]",
		response.GetHeader().RetCode.String(),
		response.GetHeader().RequestId,
		response.GetHeader().RespTime,
		response.GetHeader().ErrMsg,
		response.GetMessage())
}

func callGetConfig(ctx *context.Context, client api.CommonServiceClient, header *api.RequestHeader) {
	request := api.GetConfigRequest{}
	request.Header = header

	response, err := client.GetConfig(*ctx, &request)

	if err != nil {
		log.Println("Failed to call GetConfig %e...", err.Error())
		return
	}

	log.Println("Received GetConfig Response...")

	log.Printf("\n header:[%s, %s, %d, %s] \n message:[%s]",
		response.GetHeader().RetCode.String(),
		response.GetHeader().RequestId,
		response.GetHeader().RespTime,
		response.GetHeader().ErrMsg,
		response.GetMessage())
}

func callIsHealthy(ctx *context.Context, client api.CommonServiceClient, header *api.RequestHeader) {
	request := api.IsHealthyRequest{}
	request.Header = header

	response, err := client.IsHealthy(*ctx, &request)

	if err != nil {
		log.Println("Failed to call IsHealthy %e...", err.Error())
		return
	}

	log.Println("Received IsHealthy Response...")

	log.Printf("\n header:[%s, %s, %d, %s] \n message:[%s]",
		response.GetHeader().RetCode.String(),
		response.GetHeader().RequestId,
		response.GetHeader().RespTime,
		response.GetHeader().ErrMsg,
		response.GetMessage())
}

func callHelp(ctx *context.Context, client api.CommonServiceClient, header *api.RequestHeader) {
	request := api.HelpRequest{}
	request.Header = header

	response, err := client.Help(*ctx, &request)

	if err != nil {
		log.Println("Failed to call Help %e...", err.Error())
		return
	}

	log.Println("Received Help Response...")

	log.Printf("\n header:[%s, %s, %d, %s] \n message:[%s]",
		response.GetHeader().RetCode.String(),
		response.GetHeader().RequestId,
		response.GetHeader().RespTime,
		response.GetHeader().ErrMsg,
		response.GetMessage())
}

func callInfo(ctx *context.Context, client api.CommonServiceClient, header *api.RequestHeader) {
	request := api.InfoRequest{}
	request.Header = header

	response, err := client.Info(*ctx, &request)

	if err != nil {
		log.Println("Failed to call Info %e...", err.Error())
		return
	}

	log.Println("Received Info Response...")

	log.Printf("\n header:[%s, %s, %d, %s] \n message:[%s]",
		response.GetHeader().RetCode.String(),
		response.GetHeader().RequestId,
		response.GetHeader().RespTime,
		response.GetHeader().ErrMsg,
		response.GetMessage())
}

func callLog(ctx *context.Context, client api.CommonServiceClient, header *api.RequestHeader) {
	request := api.LogRequest{}
	request.Header = header

	response, err := client.Log(*ctx, &request)

	if err != nil {
		log.Println("Failed to call Log %e...", err.Error())
		return
	}

	log.Println("Received Log Response...")

	log.Printf("\n header:[%s, %s, %d, %s] \n message:[%s]",
		response.GetHeader().RetCode.String(),
		response.GetHeader().RequestId,
		response.GetHeader().RespTime,
		response.GetHeader().ErrMsg,
		response.GetMessage())
}

func callPing(ctx *context.Context, client api.CommonServiceClient, header *api.RequestHeader) {
	request := api.PingRequest{}
	request.Header = header

	response, err := client.Ping(*ctx, &request)

	if err != nil {
		log.Println("Failed to call Ping %e...", err.Error())
		return
	}

	log.Println("Received Ping Response...")

	log.Printf("\n header:[%s, %s, %d, %s] \n message:[%s]",
		response.GetHeader().RetCode.String(),
		response.GetHeader().RequestId,
		response.GetHeader().RespTime,
		response.GetHeader().ErrMsg,
		response.GetMessage())
}

func callShutdown(ctx *context.Context, client api.CommonServiceClient, header *api.RequestHeader) {
	request := api.ShutdownRequest{}
	request.Header = header

	response, err := client.Shutdown(*ctx, &request)

	if err != nil {
		log.Println("Failed to call Shutdown %e...", err.Error())
		return
	}

	log.Println("Received Shutdown Response...")

	log.Printf("\n header:[%s, %s, %d, %s] \n message:[%s]",
		response.GetHeader().RetCode.String(),
		response.GetHeader().RequestId,
		response.GetHeader().RespTime,
		response.GetHeader().ErrMsg,
		response.GetMessage())
}