package httpc

import (
	"net/http"
	"testing"
)

func TestHttpClientParams(t *testing.T) {
	params := newHttpClientParams()

	// Test default values
	if params.GetMaxRetryWait() != 10 {
		t.Fatalf("expected MaxRetryWait to be 10, got %d", params.GetMaxRetryWait())
	}

	if params.GetMaxRetries() != 3 {
		t.Fatalf("expected MaxRetries to be 3, got %d", params.GetMaxRetries())
	}

	// Test WithMaxRetryWait option
	params = newHttpClientParams(WithMaxRetryWait(20))
	if params.GetMaxRetryWait() != 20 {
		t.Fatalf("expected MaxRetryWait to be 20, got %d", params.GetMaxRetryWait())
	}

	// Test WithMaxRetries option
	params = newHttpClientParams(WithMaxRetries(5))
	if params.GetMaxRetries() != 5 {
		t.Fatalf("expected MaxRetries to be 5, got %d", params.GetMaxRetries())
	}

	// Test SetMaxRetryWait method
	params.SetMaxRetryWait(30)
	if params.GetMaxRetryWait() != 30 {
		t.Fatalf("expected MaxRetryWait to be 30, got %d", params.GetMaxRetryWait())
	}

	// Test SetMaxRetries method
	params.SetMaxRetries(7)
	if params.GetMaxRetries() != 7 {
		t.Fatalf("expected MaxRetries to be 7, got %d", params.GetMaxRetries())
	}

	// Test WithRetryStatusCodes option
	params = newHttpClientParams(WithRetryStatusCodes(500, 502))
	if _, ok := params.GetRetryStatusCodes()[500]; !ok {
		t.Fatalf("expected retry status code 500 to be set")
	}
	if _, ok := params.GetRetryStatusCodes()[502]; !ok {
		t.Fatalf("expected retry status code 502 to be set")
	}

	// Test WithMethodRetries option
	params = newHttpClientParams(WithMethodRetries("GET", 5))
	if params.GetMethodRetries()["GET"] != 5 {
		t.Fatalf("expected method retries for GET to be 5, got %d", params.GetMethodRetries()["GET"])
	}

	// Test WithRequestHook option
	hookCalled := false
	hook := func(_ *http.Request) { hookCalled = true }
	params = newHttpClientParams(WithRequestHook(hook))
	hooks := params.GetRequestHooks()
	if len(hooks) != 1 {
		t.Fatalf("expected 1 request hook, got %d", len(hooks))
	}
	hooks[0](nil)
	if !hookCalled {
		t.Fatalf("expected request hook to be called")
	}
}
