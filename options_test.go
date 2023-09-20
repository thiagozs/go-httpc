package httpc

import (
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
}
