package httpc

import "net/http"

type HttpClientRepo interface {
	Get(addrs string) ([]byte, error)
	Post(addrs string, body []byte) ([]byte, error)
	Delete(addrs string, payload []byte) ([]byte, error)

	GetWithResponse(addrs string) (*http.Response, error)
	PostWithResponse(addrs string, payload []byte) (*http.Response, error)
	DeleteWithResponse(addrs string, payload []byte) (*http.Response, error)
	PostFormWithResponse(addrs string) (*http.Response, error)

	SetHeader(method, key, value string)
	SetFormValue(method, key, value string)
	SetBasicAuth(method, username, password string)

	DeleteHeader(method, key string)
	DeleteFormValue(method, key string)

	GetHeaders(method string) map[string]string
	GetFormValue(method string) map[string]string
	GetBasicAuth(method string) map[string]string

	SetMaxRetry(val int)
	SetMaxRetryWaitMin(val int)
	SetRetryWaitMax(val int)

	DisableLogLevel()
	EnableLogLevel()
}
