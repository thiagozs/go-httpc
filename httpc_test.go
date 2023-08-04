package httpc_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thiagozs/go-httpc"
)

func TestHttpClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer ts.Close()

	opts := []httpc.HttpClientOptions{
		httpc.WithMaxRetries(3),
		httpc.WithMaxRetryWait(3),
	}
	client := httpc.NewHttpClient(opts...)

	t.Run("GET", func(t *testing.T) {
		resp, err := client.Get(ts.URL)
		if err != nil {
			t.Fatal(err)
		}

		if string(resp) != "OK" {
			t.Fatalf("expected OK, got %s", resp)
		}
	})

	t.Run("POST", func(t *testing.T) {
		resp, err := client.Post(ts.URL, []byte("payload"))
		if err != nil {
			t.Fatal(err)
		}

		if string(resp) != "OK" {
			t.Fatalf("expected OK, got %s", resp)
		}
	})

	t.Run("DELETE", func(t *testing.T) {
		resp, err := client.Delete(ts.URL, []byte("payload"))
		if err != nil {
			t.Fatal(err)
		}

		if string(resp) != "OK" {
			t.Fatalf("expected OK, got %s", resp)
		}
	})

	t.Run("Retry", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR"))
		}))
		defer ts.Close()

		_, err := client.Get(ts.URL)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
	})

	t.Run("Headers", func(t *testing.T) {
		client.SetHeader("GET", "Content-Type", "application/json")
		headers := client.GetHeaders("GET")
		if headers["Content-Type"] != "application/json" {
			t.Fatalf("expected application/json, got %s", headers["Content-Type"])
		}
	})

	t.Run("FormValue", func(t *testing.T) {
		client.SetFormValue("POST", "key", "value")
		forms := client.GetFormValue("POST")
		if forms["key"] != "value" {
			t.Fatalf("expected value, got %s", forms["key"])
		}
	})

	t.Run("BasicAuth", func(t *testing.T) {
		client.SetBasicAuth("POST", "username", "password")
		auth := client.GetBasicAuth("POST")
		if auth["username"] != "password" {
			t.Fatalf("expected password, got %s", auth["username"])
		}
	})
}
