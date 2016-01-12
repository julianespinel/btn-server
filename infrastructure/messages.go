package infrastructure

func GetStringMessage(key string, value string) map[string]string {
	stringMessage := make(map[string]string)
	stringMessage[key] = value
	return stringMessage
}
