package links

import (
	"bytes"
	"context"
	"encoding/json"
	links "github.com/AndreiBarbuOz/lnkshrtn/pkg/links/ports"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func TestHttpServer_CreateLink(t *testing.T) {
	ctx := context.Background()
	server := NewServer(ctx)

	var body links.LinkObjectSpec
	body.Url = "www.test.com"
	var tmp string = "abc123"
	body.Shortned = &tmp
	requestBody, err := json.Marshal(body)

	if err != nil {
		panic("Cannot marshal body")
	}

	req := httptest.NewRequest(http.MethodPost, "/links", bytes.NewReader(requestBody))
	req.Header.Add("content-type", "application/json")
	w := httptest.NewRecorder()

	server.CreateLink(w, req)
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode, "CreateLink status code")
	responseBody, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "CreateLink error")
	assertPostedLinkIsEqual(t, requestBody, responseBody)

	headers := w.Header()
	contentType := headers.Get("Content-Type")
	assert.NotEmpty(t, contentType, "GetHealth: Content type not returned")
	assert.Contains(t, contentType, "application/json", "GetHealth: content type")
}

func TestHttpServer_GetLinkById(t *testing.T) {
	ctx := context.Background()
	server := NewServer(ctx)

	var body links.LinkObjectSpec
	body.Url = "www.test.com"
	var tmp string = "abc123"
	body.Shortned = &tmp
	requestBody, err := json.Marshal(body)

	if err != nil {
		panic("Cannot marshal body")
	}

	req := httptest.NewRequest(http.MethodPost, "/links", bytes.NewReader(requestBody))
	req.Header.Add("content-type", "application/json")
	w := httptest.NewRecorder()

	server.CreateLink(w, req)
	res := w.Result()
	defer res.Body.Close()

	req = httptest.NewRequest(http.MethodGet, "/links/abc123", nil)
	req.Header.Add("content-type", "application/json")
	w = httptest.NewRecorder()

	server.GetLinkById(w, req, "abc123")
	res = w.Result()
	defer res.Body.Close()

	assert.Equal(t, 200, res.StatusCode, "CreateLink status code")
	responseBody, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "CreateLink error")
	assertPostedLinkIsEqual(t, requestBody, responseBody)

	headers := w.Header()
	contentType := headers.Get("Content-Type")
	assert.NotEmpty(t, contentType, "GetHealth: Content type not returned")
	assert.Contains(t, contentType, "application/json", "GetHealth: content type")
}

func assertPostedLinkIsEqual(t *testing.T, posted []byte, returned []byte) {
	t.Helper()
	var p links.LinkObjectSpec
	var r links.LinkObject
	err := json.Unmarshal(posted, &p)
	assert.Nil(t, err, "Could not unmarshal posted body")
	err = json.Unmarshal(returned, &r)
	assert.Nil(t, err, "Could not unmarshal returned body")

	assertLinkEquals(t, &p, &r.Spec)
}

var cmpRoundTimeOpt = cmp.Comparer(func(x, y time.Time) bool {
	return x.Truncate(time.Millisecond).Equal(y.Truncate(time.Millisecond))
})

func assertLinkEquals(t *testing.T, lnk1 *links.LinkObjectSpec, lnk2 *links.LinkObjectSpec) {
	t.Helper()
	cmpOpts := []cmp.Option{
		cmpRoundTimeOpt,
		cmp.AllowUnexported(time.Time{}),
	}
	assert.True(t, cmp.Equal(lnk1, lnk2, cmpOpts...))
}
