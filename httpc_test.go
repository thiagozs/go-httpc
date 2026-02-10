package httpc

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHttpClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello, world!"))
	}))

	defer server.Close()

	client := NewHttpClient()
	_, body, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(body) == 0 {
		t.Fatalf("expected non-empty response body")
	}
}

func TestHttpClient_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello, world!"))
	}))

	defer server.Close()

	client := NewHttpClient()
	_, body, err := client.Post(server.URL, []byte("test payload"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(body) == 0 {
		t.Fatalf("expected non-empty response body")
	}
}

func TestHttpClient_SetHeader(t *testing.T) {
	client := NewHttpClient()
	client.SetHeader(http.MethodGet, "Test-Header", "TestValue")

	headers := client.GetHeaders(http.MethodGet)
	if headers["Test-Header"] != "TestValue" {
		t.Fatalf("expected header value to be 'TestValue', got %v", headers["Test-Header"])
	}
}

func TestHttpClient_SetBasicAuth(t *testing.T) {
	client := NewHttpClient()
	client.SetBasicAuth(http.MethodGet, "TestUser", "TestPass")

	auth := client.GetBasicAuth(http.MethodGet)
	if auth["TestUser"] != "TestPass" {
		t.Fatalf("expected password to be 'TestPass', got %v", auth["TestUser"])
	}
}

func TestHttpClient_Put(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello, world!"))
	}))

	defer server.Close()

	client := NewHttpClient()
	_, body, err := client.Put(server.URL, []byte("test payload"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(body) == 0 {
		t.Fatalf("expected non-empty response body")
	}
}

func TestHttpClient_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello, world!"))
	}))

	defer server.Close()

	client := NewHttpClient()
	_, body, err := client.Delete(server.URL, []byte("test payload"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(body) == 0 {
		t.Fatalf("expected non-empty response body")
	}
}

func TestHttpClient_Patch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello, world!"))
	}))

	defer server.Close()

	client := NewHttpClient()
	_, body, err := client.Patch(server.URL, []byte("test payload"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(body) == 0 {
		t.Fatalf("expected non-empty response body")
	}
}

func TestHttpClient_SetFormValue(t *testing.T) {
	client := NewHttpClient()
	client.SetFormValue(http.MethodPost, "Test-Form-Key", "TestValue")

	forms := client.GetFormValue(http.MethodPost)
	if forms["Test-Form-Key"] != "TestValue" {
		t.Fatalf("expected form value to be 'TestValue', got %v", forms["Test-Form-Key"])
	}
}

func TestHttpClient_DeleteFormValue(t *testing.T) {
	client := NewHttpClient()
	client.SetFormValue(http.MethodPost, "Test-Form-Key", "TestValue")
	client.DeleteFormValue(http.MethodPost, "Test-Form-Key")

	forms := client.GetFormValue(http.MethodPost)
	if _, exists := forms["Test-Form-Key"]; exists {
		t.Fatalf("expected form key 'Test-Form-Key' to be deleted")
	}
}

func TestHttpClient_DeleteHeader(t *testing.T) {
	client := NewHttpClient()
	client.SetHeader(http.MethodGet, "Test-Header", "TestValue")
	client.DeleteHeader(http.MethodGet, "Test-Header")

	headers := client.GetHeaders(http.MethodGet)
	if _, exists := headers["Test-Header"]; exists {
		t.Fatalf("expected header key 'Test-Header' to be deleted")
	}
}

func TestHttpClient_PostWithFormValue(t *testing.T) {
	// Create a test server that echoes the request body
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected method POST, got %s", r.Method)
		}

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/x-www-form-urlencoded" {
			t.Fatalf("expected Content-Type application/x-www-form-urlencoded, got %s", contentType)
		}

		err := r.ParseForm()
		if err != nil {
			t.Fatalf("error parsing form: %v", err)
		}

		if r.FormValue("key1") != "value1" || r.FormValue("key2") != "value2" {
			t.Fatalf("unexpected form values: %v", r.Form)
		}

		w.Write([]byte("Form values received"))
	}))
	defer ts.Close()

	client := NewHttpClient()
	client.SetFormValue(http.MethodPost, "key1", "value1")
	client.SetFormValue(http.MethodPost, "key2", "value2")

	_, body, err := client.Post(ts.URL, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(body) != "Form values received" {
		t.Fatalf("unexpected response: %s", body)
	}
}

func TestHttpClient_Head(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			t.Fatalf("expected HEAD request, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewHttpClient()

	t.Run("Valid HEAD request", func(t *testing.T) {
		resp, body, err := client.Head(server.URL)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status code 200, got %d", resp.StatusCode)
		}

		contentType := resp.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Fatalf("expected content type application/json, got %s", contentType)
		}

		if len(body) != 0 {
			t.Fatalf("expected empty response body, got %v", body)
		}
		resp.Body.Close()
	})
}

func TestHttpClient_SetPatchHeader(t *testing.T) {
	client := NewHttpClient()

	key := "Content-Type"
	value := "application/json"
	client.SetPatchHeader(key, value)

	headers := client.GetHeaders(http.MethodPatch)
	if headers == nil {
		t.Fatalf("expected headers to be initialized, got nil")
	}

	if headers[key] != value {
		t.Fatalf("expected header %s to be %s, got %s", key, value, headers[key])
	}
}

func TestHttpClient_FormPostValueOverride(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		assert.NoError(t, err)

		formValues := url.Values{}
		for key, values := range r.PostForm {
			for _, value := range values {
				formValues.Add(key, value)
			}
		}
		w.Write([]byte(formValues.Encode()))
	}))
	defer ts.Close()

	client := NewHttpClient()
	// Set initial form value
	client.SetFormValue(http.MethodPost, "key", "initialValue")
	// Override form value
	client.SetFormValue(http.MethodPost, "key", "updatedValue")

	_, body, err := client.Post(ts.URL, nil)
	assert.NoError(t, err)

	responseFormValues, err := url.ParseQuery(string(body))
	assert.NoError(t, err)
	assert.Equal(t, "updatedValue", responseFormValues.Get("key"))
}

func TestHttpClient_HeaderValueOverride(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeaderValue := r.Header.Get("Test-Header")
		w.Write([]byte(receivedHeaderValue))
	}))
	defer ts.Close()

	client := NewHttpClient()
	// Set initial header value
	client.SetHeader(http.MethodGet, "Test-Header", "initialValue")
	// Override header value
	client.SetHeader(http.MethodGet, "Test-Header", "updatedValue")

	_, body, err := client.Get(ts.URL)
	assert.NoError(t, err)

	assert.Equal(t, "updatedValue", string(body))
}

func TestHttpClient_BasicAuthOverride(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		w.Write([]byte(user + ":" + pass))
	}))
	defer ts.Close()

	client := NewHttpClient()
	// Set initial basic auth credentials
	client.SetBasicAuth(http.MethodGet, "initialUser", "initialPass")
	// Override basic auth credentials
	client.SetBasicAuth(http.MethodGet, "updatedUser", "updatedPass")

	_, body, err := client.Get(ts.URL)
	assert.NoError(t, err)

	assert.Equal(t, "updatedUser:updatedPass", string(body))
}

func TestHttpClient_PatchHeaderOverride(t *testing.T) {
	// Create a test server that echoes back the received header value
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeaderValue := r.Header.Get("Test-Header")
		w.Write([]byte(receivedHeaderValue))
	}))
	defer ts.Close()

	client := NewHttpClient()
	// Set initial header value
	client.SetPatchHeader("Test-Header", "initialValue")
	// Override header value
	client.SetPatchHeader("Test-Header", "updatedValue")

	// Make a PATCH request
	_, body, err := client.Patch(ts.URL, nil)
	assert.NoError(t, err)

	// Check if the header value in the response is the updated value
	assert.Equal(t, "updatedValue", string(body))
}

func TestHttpClient_PostWithoutPayloadDoesNotSetJSONContentType(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "" {
			t.Fatalf("expected empty Content-Type, got %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewHttpClient()
	_, _, err := client.Post(ts.URL, nil)
	assert.NoError(t, err)
}

func TestHttpClient_RequestHookSetsHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.Header.Get("X-Req-Hook")))
	}))
	defer ts.Close()

	client := NewHttpClient(WithRequestHook(func(req *http.Request) {
		req.Header.Set("X-Req-Hook", "hooked")
	}))

	_, body, err := client.Get(ts.URL)
	assert.NoError(t, err)
	assert.Equal(t, "hooked", string(body))
}

func TestHttpClient_GetWithContext_ReturnsPayload(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ctx-get"))
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := NewHttpClient()
	_, body, err := client.GetWithContext(ctx, ts.URL)
	assert.NoError(t, err)
	assert.Equal(t, "ctx-get", string(body))
}

func TestHttpClient_PostWithContext_ReturnsPayload(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := NewHttpClient()
	_, body, err := client.PostWithContext(ctx, ts.URL, []byte("ctx-post"))
	assert.NoError(t, err)
	assert.Equal(t, "ctx-post", string(body))
}

func TestHttpClient_HeadWithContext_ReturnsResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "ok")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := NewHttpClient()
	resp, body, err := client.HeadWithContext(ctx, ts.URL)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Equal(t, "ok", resp.Header.Get("X-Test"))
	assert.Len(t, body, 0)
	resp.Body.Close()
}

func TestHttpClient_HeadWithContext_Cancelled(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	client := NewHttpClient()
	resp, body, err := client.HeadWithContext(ctx, ts.URL)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, body)
}

func TestHttpClient_GetWithContext_Cancelled(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	client := NewHttpClient()
	_, _, err := client.GetWithContext(ctx, ts.URL)
	assert.Error(t, err)
}

func TestHttpClient_RetryStatusCodePerMethod(t *testing.T) {
	attempts := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	client := NewHttpClient(
		WithMethodRetries(http.MethodGet, 3),
		WithRetryStatusCodes(http.StatusInternalServerError),
		WithMaxRetryWait(0),
	)

	_, body, err := client.Get(ts.URL)
	assert.NoError(t, err)
	assert.Equal(t, "ok", string(body))
	assert.Equal(t, 3, attempts)
}
