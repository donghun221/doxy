package utils

const (
	PORT string = "19270"
	NAME string = "HelloTencent"
	CONF_SUFFIX string = "conf"
	APP_CONF_NAME string = NAME + DOT + CONF_SUFFIX

	// common utility constants
	DOT string = "."

	// domain related
	TEST_DOMAIN string = "test"
	BETA_DOMAIN string = "beta"
	GAMMA_DOMAIN string = "gamma"
	ONEBOX_DOMAIN string = "onebox"
	PROD_DOMAIN string = "prod"

	// logger related
	LOG_PATH string = "logs/HelloTencent.log"

	// config related
	CONF_PATH string = "./configs/" + APP_CONF_NAME

	EMPTY_STRING string = ""

	REQUEST_ID string = "RequestId"
	EXTERNAL_ID string = "ExternalId"

	ELAPSED_TIME string = "ElapsedTime"
	NIL string = "nil"
	ERROR string = "Error"
	RET_CODE string = "RetCode"

	PONG string = "PONG"

	// Cli related
	//LOCAL_HOST string = "localhost"
	LOCAL_HOST string = "0.0.0.0"

	// symbol related
	COMMA string = ","
	COLON string = ":"
	DOUBLE_COLON string = "::"

	// event data related
	APPLICATION_NAME string = "Application"
	HOST_NAME string = "Host"
	OPERATION string = "Operation"
	STATUS string = "Status"
	QUERY_LOG_STATUS string = "QueryLogStatus"
	START_TIME string = "StartTime"
	END_TIME string = "EndTime"
	TIMING string = "Timing"
	REMOTE_ADDR string = "Remote_Addr"
	INFO string = "INFO"
	PROGRAM string = "Program"
	TIME string = "Time"
)
