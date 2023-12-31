// Code generated by ifacemaker; DO NOT EDIT.

package httpc

import (
	"net/http"
)

// HttpClientRepo ...
type HttpClientRepo interface {
	Get(addrs string) ([]byte, error)
	Post(addrs string, payload []byte) ([]byte, error)
	Put(addrs string, payload []byte) ([]byte, error)
	Delete(addrs string, payload []byte) ([]byte, error)
	Patch(addrs string, payload []byte) ([]byte, error)
	Head(addrs string) (*http.Response, error)
	SetHeader(method, key, value string)
	DeleteHeader(method, key string)
	SetFormValue(method, key, value string)
	DeleteFormValue(method, key string)
	GetFormValue(method string) map[string]string
	SetBasicAuth(method, username, password string)
	GetBasicAuth(method string) map[string]string
	SetPatchHeader(key, value string)
	GetHeaders(method string) map[string]string
}
