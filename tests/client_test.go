package tests

import (
	helper "github.com/outscope-solutions/acdn-go-client/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientAuthenticateOK(t *testing.T) {
	client := helper.GetClient()
	err := client.Authenticate()
	assert.Nil(t, err)
}

func TestClientAuthenticateNOK(t *testing.T) {
	client := helper.GetFakeClient("", "test", "test")
	err := client.Authenticate()
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "authentication Failed", err)
	}
}

func TestClientBadHostScheme(t *testing.T) {
	client := helper.GetFakeClient("dummyserver", "", "")
	err := client.Authenticate()
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "Post \"dummyserver/controller/v2/tokens\": unsupported protocol scheme \"\"", err)
	}
}

func TestClientHostConnectionRefused(t *testing.T) {
	client := helper.GetFakeClient("https://127.0.0.1:80", "", "")
	err := client.Authenticate()
	if assert.NotNil(t, err) {
		assert.EqualError(t, err, "Post \"https://127.0.0.1:80/controller/v2/tokens\": dial tcp 127.0.0.1:80: connect: connection refused", err)
	}
}
