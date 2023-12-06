package httpc

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello, world!"))
	}))

	defer server.Close()

	client := NewHttpClient()
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp) == 0 {
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
	resp, err := client.Post(server.URL, []byte("test payload"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp) == 0 {
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
	resp, err := client.Put(server.URL, []byte("test payload"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp) == 0 {
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
	resp, err := client.Delete(server.URL, []byte("test payload"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp) == 0 {
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
	resp, err := client.Patch(server.URL, []byte("test payload"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp) == 0 {
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

func TestRetryableTransport_RoundTrip(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	transport := &RetryableTransport{
		Transport:  http.DefaultTransport,
		Retries:    3,
		WaitMaxSec: 1,
	}

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatalf("creating request failed: %v", err)
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code 200, got %v", resp.StatusCode)
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

	resp, err := client.Post(ts.URL, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(resp) != "Form values received" {
		t.Fatalf("unexpected response: %s", resp)
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
		resp, err := client.Head(server.URL)
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

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}

		if len(body) != 0 {
			t.Fatalf("expected empty response body, got %v", body)
		}
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

	resp, err := client.Post(ts.URL, nil)
	assert.NoError(t, err)

	responseFormValues, err := url.ParseQuery(string(resp))
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

	resp, err := client.Get(ts.URL)
	assert.NoError(t, err)

	assert.Equal(t, "updatedValue", string(resp))
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

	resp, err := client.Get(ts.URL)
	assert.NoError(t, err)

	assert.Equal(t, "updatedUser:updatedPass", string(resp))
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
	resp, err := client.Patch(ts.URL, nil)
	assert.NoError(t, err)

	// Check if the header value in the response is the updated value
	assert.Equal(t, "updatedValue", string(resp))
}
