package metahttp_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-kit/log"
	metahttp "github.com/krishnateja262/meta-http/pkg/meta_http"
)

func TestMetaHTTPClient(t *testing.T) {
	responseBody := "{\"Goodbye\":\"World\"}"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(responseBody))
	}))
	defer server.Close()

	logger := log.NewJSONLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)

	metaHttpClient := metahttp.NewClient(server.URL, logger, 10*time.Second)
	req := struct {
		Hello string
	}{
		Hello: "world",
	}
	var res struct {
		Goodbye string
	}
	err := metaHttpClient.Post(context.Background(), "/test", map[string]string{}, req, &res)
	if err != nil {
		t.Error(err.Error())
	}
	if res.Goodbye != "World" {
		t.Error("Response body is not as expected")
	}
}
