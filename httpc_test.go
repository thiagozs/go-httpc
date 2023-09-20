package httpc

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
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
