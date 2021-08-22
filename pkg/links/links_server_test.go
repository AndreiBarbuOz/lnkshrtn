package links

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpServer_GetHealth(t *testing.T) {
	ctx := context.Background()
	server := NewServer(ctx)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	server.GetHealth(w, req)
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode, "GetHealth status code")
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "GetHealth error")
	assert.JSONEq(t, `{"status": "ok"}`, string(data), "GetHealth body")

	headers := w.Header()
	contentType := headers.Get("Content-Type")
	assert.NotEmpty(t, contentType, "GetHealth: Content type not returned")
	assert.Contains(t, contentType, "application/json", "GetHealth: content type")
}
