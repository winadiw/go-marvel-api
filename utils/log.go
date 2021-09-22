package utils

// LoggerFormat returns log format for middleware
func LoggerFormat() string {
	return "[${time}] ${pid} ${locals:requestid} ${status} ${path} {\"status\":${status},\"reqBody\":\"${body}\",\"resBody\":\"${resBody}\"}\n"
}
