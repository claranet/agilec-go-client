package client

import (
	"os"
)

func GetClientTest() *Client {
	return GetClient(
		os.Getenv("AC_HOST"),
		os.Getenv("AC_USERNAME"),
		os.Getenv("AC_PASSWORD"),
		Insecure(true),
		TimeoutInSeconds(10),
		RetryCount(0),
		RetryWaitTimeInSeconds(10),
		RetryMaxWaitTimeSeconds(20),
		LogLevel(0))
}
