package helper

func FailedResponseHelper(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
		"message": msg,
	}
}

func SuccessResponseHelper(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func SuccessDataResponseHelper(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}
