package querylog

const (
	// logger related
	EVENT_DATA_LOG_PATH string = "./logs/service_log.log"
	EVENT_DATA_MAX_LOG_SIZE int = 2 * 1024 // 1 GB
	EVENT_DATA_MAX_BACKUPS int = 10 // 10 log files
	EVENT_DATA_MAX_AGE int = 90 // 3 months
	EVENT_DATA_LOCAL_TIME bool = true

	EMPTY_STRING string = ""
)
