package utils

func LoggerFormat() string {
	return "[${time}] ${pid} ${locals:requestid} ${status} ${path} {\"status\":${status},\"reqBody\":\"${body}\",\"resBody\":\"${resBody}\"}\n"
}
