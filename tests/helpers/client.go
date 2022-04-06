package helpers

import (
	"agilec-go-client/client"
	"os"
)

func GetClient() *client.Client {
	return client.GetClient(
		os.Getenv("AC_HOST"),
		os.Getenv("AC_USERNAME"),
		os.Getenv("AC_PASSWORD"),
		client.Insecure(true),
		client.TimeoutInSeconds(10),
		client.RetryCount(0),
		client.RetryWaitTimeInSeconds(10),
		client.RetryMaxWaitTimeSeconds(20),
		client.LogLevel(0))
}

func GetFakeClient(url, username, password string) *client.Client {
	baseUrl := os.Getenv("AC_HOST")
	user := os.Getenv("AC_USERNAME")
	pass := os.Getenv("AC_PASSWORD")

	if url != "" {
		baseUrl = url
	}

	if username != "" {
		user = url
	}

	if password != "" {
		pass = url
	}

	return client.NewClient(baseUrl, user, pass, client.Insecure(true), client.LogLevel(6))
}
