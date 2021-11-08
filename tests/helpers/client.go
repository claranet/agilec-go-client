package helpers

import (
	"github.com/outscope-solutions/acdn-go-client/client"
	"os"
)

func GetClient() *client.Client {
	return client.GetClient(
		os.Getenv("AC_HOST"),
		os.Getenv("AC_USERNAME"),
		client.Password(os.Getenv("AC_PASSWORD")),
		client.Insecure(true))
}

func GetFakeClient(url, username, password string) *client.Client  {
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

	return client.NewClient(baseUrl, user, client.Password(pass), client.Insecure(true))
}