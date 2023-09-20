package httpc

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	UA_KEY = "User-Agent"
	UA_VAL = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36"
)

type RetryableTransport struct {
	Transport  http.RoundTripper
	Retries    int
	WaitMaxSec int
}

func (t *RetryableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; i < t.Retries; i++ {
		resp, err = t.Transport.RoundTrip(req)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				time.Sleep(time.Second * time.Duration(t.WaitMaxSec))
				continue
			}
			return nil, err
		}
		break
	}

	return resp, err
}

type HttpClient struct {
	sync.Mutex
	client       *http.Client
	maxRetry     int
	retryWaitMin int
	retryWaitMax int
	headers      map[string]map[string]string
	forms        map[string]map[string]string
	basicAuth    map[string]map[string]string
	params       *HttpClientParams
}

func NewHttpClient(opts ...HttpClientOptions) *HttpClient {

	params := newHttpClientParams(opts...)

	client := &http.Client{
		Transport: &RetryableTransport{
			Transport:  http.DefaultTransport,
			Retries:    params.GetMaxRetries(),
			WaitMaxSec: params.GetMaxRetryWait(),
		},
	}

	return &HttpClient{
		client:    client,
		headers:   make(map[string]map[string]string),
		forms:     make(map[string]map[string]string),
		basicAuth: make(map[string]map[string]string),
		params:    params,
	}
}

func (c *HttpClient) setHeaders(method string, req *http.Request) {
	c.Lock()
	defer c.Unlock()
	headers, exists := c.headers[strings.ToUpper(method)]
	if exists {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
}

func (c *HttpClient) setBasicAuth(method string, req *http.Request) {
	c.Lock()
	defer c.Unlock()
	auth, exists := c.basicAuth[strings.ToUpper(method)]
	if exists {
		for username, password := range auth {
			req.SetBasicAuth(username, password)
		}
	}
}

func (c *HttpClient) doRequest(method, addrs string, payload []byte) ([]byte, error) {
	formValues := c.GetFormValue(method)
	var req *http.Request
	var err error

	if len(formValues) > 0 {
		form := url.Values{}
		for key, value := range formValues {
			form.Add(key, value)
		}
		req, err = http.NewRequest(method, addrs, strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(method, addrs, bytes.NewBuffer(payload))
		if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
			req.Header.Add("Content-Type", "application/json")
		}
	}

	if err != nil {
		return nil, fmt.Errorf("creating request failed: %w", err)
	}

	c.setHeaders(method, req)
	c.setBasicAuth(method, req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("server error: %v", resp.StatusCode)
	} else if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("client error: %v", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (c *HttpClient) Get(addrs string) ([]byte, error) {
	return c.doRequest(http.MethodGet, addrs, nil)
}

func (c *HttpClient) Post(addrs string, payload []byte) ([]byte, error) {
	return c.doRequest(http.MethodPost, addrs, payload)
}

func (c *HttpClient) Put(addrs string, payload []byte) ([]byte, error) {
	return c.doRequest(http.MethodPut, addrs, payload)
}

func (c *HttpClient) Delete(addrs string, payload []byte) ([]byte, error) {
	return c.doRequest(http.MethodDelete, addrs, payload)
}

func (c *HttpClient) Patch(addrs string, payload []byte) ([]byte, error) {
	return c.doRequest(http.MethodPatch, addrs, payload)
}

func (c *HttpClient) Head(addrs string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodHead, addrs, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %w", err)
	}

	c.setHeaders(http.MethodHead, req)
	c.setBasicAuth(http.MethodHead, req)

	return c.client.Do(req)
}

func (c *HttpClient) SetHeader(method, key, value string) {
	c.Lock()
	defer c.Unlock()
	v, ok := c.headers[strings.ToUpper(method)]
	if ok {
		_, ook := v[key]
		if !ook {
			v[key] = value
		}
	} else {
		c.headers[strings.ToUpper(method)] = make(map[string]string)
		c.headers[strings.ToUpper(method)][key] = value
	}
}

func (c *HttpClient) DeleteHeader(method, key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.headers[strings.ToUpper(method)], key)
}

func (c *HttpClient) SetFormValue(method, key, value string) {
	c.Lock()
	defer c.Unlock()
	v, ok := c.forms[strings.ToUpper(method)]
	if ok {
		_, ook := v[key]
		if !ook {
			v[key] = value
		}
	} else {
		c.forms[strings.ToUpper(method)] = make(map[string]string)
		c.forms[strings.ToUpper(method)][key] = value
	}
}

func (c *HttpClient) DeleteFormValue(method, key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.forms[strings.ToUpper(method)], key)
}

func (c *HttpClient) GetHeaders(method string) map[string]string {
	c.Lock()
	defer c.Unlock()
	return c.headers[strings.ToUpper(method)]
}

func (c *HttpClient) GetFormValue(method string) map[string]string {
	c.Lock()
	defer c.Unlock()
	return c.forms[strings.ToUpper(method)]
}

func (c *HttpClient) SetBasicAuth(method, username, password string) {
	c.Lock()
	defer c.Unlock()
	v, ok := c.basicAuth[strings.ToUpper(method)]
	if ok {
		_, ook := v[username]
		if !ook {
			v[username] = password
		}
	} else {
		c.basicAuth[strings.ToUpper(method)] = make(map[string]string)
		c.basicAuth[strings.ToUpper(method)][username] = password
	}
}

func (c *HttpClient) GetBasicAuth(method string) map[string]string {
	c.Lock()
	defer c.Unlock()
	return c.basicAuth[strings.ToUpper(method)]
}
