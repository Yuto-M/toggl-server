package integration_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	. "github.com/Eun/go-hit"
)

const (
	host = "app:8080"

	// HTTP REST
	basePath = "http://" + host + "/v1"
)

func TestMain(m *testing.M) {
	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

// HTTP POST: /translation/do-translate.
func TestHTTPDoTranslate(t *testing.T) {
	body := `{
		"destination": "en",
		"original": "текст для перевода",
		"source": "auto"
	}`
	Test(t,
		Description("DoTranslate Success"),
		Post(basePath+"/translation/do-translate"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".translation").Equal("text to be translated"),
	)

	body = `{
		"destination": "en",
		"original": "текст для перевода"
	}`
	Test(t,
		Description("DoTranslate Fail"),
		Post(basePath+"/translation/do-translate"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").Equal("invalid request body"),
	)
}

// HTTP GET: /translation/history.
func TestHTTPHistory(t *testing.T) {
	Test(t,
		Description("History Success"),
		Get(basePath+"/translation/history"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Contains(`{"history":[{`),
	)
}
