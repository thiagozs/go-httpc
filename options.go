package httpc

type HttpClientParams struct {
	MaxRetryWait int
	MaxRetries   int
}

type HttpClientOptions func(*HttpClientParams)

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
