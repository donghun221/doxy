package main

import (
	"HelloTencent/api"
	constants "HelloTencent/internal/utils"
	"HelloTencent/internal/config"
	"HelloTencent/internal/service"
	"ReplixLogger4Go/src"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Auto generation start
	// 1: Init Logger
	logger.Info("Start initializing logger...")
	initLogger()
	logger.Info("Initializing logger success...")

	// 2: Init config
	logger.Info("Start initializing config...")
	initConfig()
	logger.Info("Initializing config success...")

	// 3: Init EventDatq
	logger.Info("Start initializing event data...")
	initEventData()
	logger.Info("Initializing event data success...")

	// 4: Register SIG for shutdown hook
	logger.Infof("Start registering shutdown hook...")
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)
	logger.Infof("Registering shutdown hook success...")

	// 5: Constructing gRpc server
	logger.Infof("Start initializing grpc server...")
	listener, server := initServer()
	logger.Infof("Initializing grpc server success...")

	// 5: Prepare shutdown hook
	go func() {
		sig := <-gracefulStop
		logger.Infof("Caught sig: %+v", sig)

		server.GracefulStop()
		logger.Infof("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	// 6: Start Server
	logger.Infof("Start serving at:%s ...", listener.Addr().String())
	if err := server.Serve(listener); err != nil {
		errMsg := fmt.Sprintf("Failed to start ReplixDataService, %v", err)
		fmt.Println(errMsg)
		logger.Fatalf(errMsg)
		os.Exit(1)
	}
	// Auto generation end
}

// Auto generated
func initLogger() {
	fileLogConfig := replixlogger.FileLogConfig{
		Filename:	constants.LOG_PATH,
	}

	logConfig := replixlogger.LogConfig{
		File: fileLogConfig,
	}

	replixlogger.InitLogger(&logConfig)
}

// Auto generated
func initConfig() {
	logger.Infof("Start loading config file from %s for domain %s ...",
		constants.APP_CONF_NAME,
		constants.PROD_DOMAIN)

	// Parse config
	_, err := config.NewConfig(constants.CONF_PATH, constants.PROD_DOMAIN)

	if err != nil {
		// Log it!
		logger.Infof("Loading config file from %s for domain %s failed, shutting down... err:%s",
			constants.APP_CONF_NAME,
			constants.PROD_DOMAIN, err)
		os.Exit(1)
	}

	logger.Infof("Loading config file from %s for domain %s success...",
		constants.APP_CONF_NAME,
		constants.PROD_DOMAIN)
}

// Auto generated
func initServer() (net.Listener, *grpc.Server) {
	logger.Infof("Start constructing gRpc listener...")
	listener, err := net.Listen("tcp", ":" + constants.PORT)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to listen on port:%s on TCP %v, shutting down...", constants.PORT, err)
		logger.Fatalf(errMsg)
		os.Exit(1)
	}
	logger.Infof("Constructing gRpc listener success...")

	// 5: Constructing gRpc server
	logger.Infof("Start constructing gRpc server...")
	// Create a gRPC server with gRPC interceptor
	server := grpc.NewServer()

	helloTencentService := service.HelloTencentService{}
	commonService := service.CommonService{}

	api.RegisterHelloTencentServiceServer(server, &helloTencentService)
	api.RegisterCommonServiceServer(server, &commonService)

	return listener, server
}

// Auto generated
func initEventData() {

}