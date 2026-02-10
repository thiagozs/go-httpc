package httpc

import "net/http"

type HttpClientParams struct {
	MaxRetryWait     int
	MaxRetries       int
	RetryStatusCodes map[int]struct{}
	MethodRetries    map[string]int
	RequestHooks     []RequestHook
}

type HttpClientOptions func(*HttpClientParams)
type RequestHook func(*http.Request)

func newHttpClientParams(opts ...HttpClientOptions) *HttpClientParams {
	s := &HttpClientParams{
		MaxRetryWait: 10,
		MaxRetries:   3,
	} //default values

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithMaxRetryWait(maxRetryWait int) HttpClientOptions {
	return func(s *HttpClientParams) {
		s.MaxRetryWait = maxRetryWait
	}
}

func WithMaxRetries(maxRetries int) HttpClientOptions {
	return func(s *HttpClientParams) {
		s.MaxRetries = maxRetries
	}
}

func WithRetryStatusCodes(codes ...int) HttpClientOptions {
	return func(s *HttpClientParams) {
		if len(codes) == 0 {
			s.RetryStatusCodes = nil
			return
		}
		s.RetryStatusCodes = make(map[int]struct{}, len(codes))
		for _, code := range codes {
			s.RetryStatusCodes[code] = struct{}{}
		}
	}
}

func WithMethodRetries(method string, retries int) HttpClientOptions {
	return func(s *HttpClientParams) {
		if s.MethodRetries == nil {
			s.MethodRetries = make(map[string]int)
		}
		s.MethodRetries[method] = retries
	}
}

func WithRequestHook(hook RequestHook) HttpClientOptions {
	return func(s *HttpClientParams) {
		if hook == nil {
			return
		}
		s.RequestHooks = append(s.RequestHooks, hook)
	}
}

// getters and setters -----

func (s *HttpClientParams) GetMaxRetryWait() int {
	return s.MaxRetryWait
}

func (s *HttpClientParams) SetMaxRetryWait(maxRetryWait int) {
	s.MaxRetryWait = maxRetryWait
}

func (s *HttpClientParams) GetMaxRetries() int {
	return s.MaxRetries
}

func (s *HttpClientParams) SetMaxRetries(maxRetries int) {
	s.MaxRetries = maxRetries
}

func (s *HttpClientParams) GetRetryStatusCodes() map[int]struct{} {
	return cloneIntSet(s.RetryStatusCodes)
}

func (s *HttpClientParams) GetMethodRetries() map[string]int {
	return cloneStringIntMap(s.MethodRetries)
}

func (s *HttpClientParams) GetRequestHooks() []RequestHook {
	if len(s.RequestHooks) == 0 {
		return nil
	}
	clone := make([]RequestHook, len(s.RequestHooks))
	copy(clone, s.RequestHooks)
	return clone
}

func cloneIntSet(values map[int]struct{}) map[int]struct{} {
	if len(values) == 0 {
		return nil
	}
	clone := make(map[int]struct{}, len(values))
	for k := range values {
		clone[k] = struct{}{}
	}
	return clone
}

func cloneStringIntMap(values map[string]int) map[string]int {
	if len(values) == 0 {
		return nil
	}
	clone := make(map[string]int, len(values))
	for k, v := range values {
		clone[k] = v
	}
	return clone
}
