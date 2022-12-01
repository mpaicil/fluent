package fluent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCreation(t *testing.T) {
	exampleUrl := "http://example.org"
	client := GET(exampleUrl)

	assert.Equal(t, client.url, exampleUrl, "The URL is not present in the client")
	assert.Equal(t, client.method, "GET", "The method is not correct")
}

func TestGetWithHeader(t *testing.T) {
	client := GET("http://example.org").withHeader("key", "value")

	assert.Equal(t, client.headers["key"], "value", "The value is not present in the headers")
}

func TestBasicCall(t *testing.T) {
	response := GET("http://example.com:3002/test").call()

	assert.Equal(t, 204, response.StatusCode, "The status code is not 204")
}

func TestAuthCall(t *testing.T) {
	response := GET("http://example.com:3002/authtest").withHeader("Authorization", "Bearer 123").call()

	assert.Equal(t, 204, response.StatusCode, "The status code is not 204")
}

func TestBasicPostCall(t *testing.T) {
	response := POST("http://localhost:3002/basic").call()

	assert.Equal(t, 204, response.StatusCode, "The status code is not 204")
}

func TestPostCallWithBody(t *testing.T) {
	type body struct {
		Field string `json:"field"`
	}

	response := POST("http://localhost:3002/body").withBody(body{Field: "value"}).withHeader("Content-Type", "application/json").call()

	assert.Equal(t, 201, response.StatusCode, "The status code is not 204")
}
