package httpc

import (
	"bytes"
	"fmt"
	"io"
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
			time.Sleep(time.Second * time.Duration(t.WaitMaxSec))
			continue
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
	setBasicAuth map[string]map[string]string
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
		client:       client,
		headers:      make(map[string]map[string]string),
		forms:        make(map[string]map[string]string),
		setBasicAuth: make(map[string]map[string]string),
		params:       params,
	}
}

func (c *HttpClient) Get(addrs string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, addrs, nil)
	if err != nil {
		return []byte{}, err
	}
	if len(c.headers[http.MethodGet]) > 0 {
		for k, v := range c.headers[http.MethodGet] {
			req.Header.Set(k, v)
		}
	}
	if len(c.setBasicAuth[http.MethodPost]) > 0 {
		auth := c.setBasicAuth[http.MethodPost]
		var user, pass string
		for k, v := range auth {
			user = k
			pass = v
		}
		req.SetBasicAuth(user, pass)
	}

	rr, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer rr.Body.Close()

	if rr.StatusCode >= 500 {
		return nil, fmt.Errorf("server error: %v", rr.StatusCode)
	}

	return io.ReadAll(rr.Body)
}

func (c *HttpClient) Post(addrs string, payload []byte) ([]byte, error) {

	req, err := http.NewRequest(http.MethodPost, addrs, bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, err
	}
	if len(c.headers[http.MethodPost]) > 0 {
		for k, v := range c.headers[http.MethodPost] {
			req.Header.Set(k, v)
		}
	}
	if len(c.setBasicAuth[http.MethodPost]) > 0 {
		auth := c.setBasicAuth[http.MethodPost]
		var user, pass string
		for k, v := range auth {
			user = k
			pass = v
		}
		req.SetBasicAuth(user, pass)
	}

	rr, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer rr.Body.Close()

	if rr.StatusCode >= 500 {
		return nil, fmt.Errorf("server error: %v", rr.StatusCode)
	}

	return io.ReadAll(rr.Body)
}

func (c *HttpClient) Delete(addrs string, payload []byte) ([]byte, error) {

	req, err := http.NewRequest(http.MethodDelete, addrs, bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, err
	}
	if len(c.headers[http.MethodDelete]) > 0 {
		for k, v := range c.headers[http.MethodDelete] {
			req.Header.Set(k, v)
		}
	}

	rr, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer rr.Body.Close()

	if rr.StatusCode >= 500 {
		return nil, fmt.Errorf("server error: %v", rr.StatusCode)
	}

	return io.ReadAll(rr.Body)
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

func (c *HttpClient) GetWithResponse(addrs string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, addrs, nil)
	if err != nil {
		return &http.Response{}, err
	}
	if len(c.headers[http.MethodGet]) > 0 {
		for k, v := range c.headers[http.MethodGet] {
			req.Header.Set(k, v)
		}
	}
	if len(c.setBasicAuth[http.MethodPost]) > 0 {
		auth := c.setBasicAuth[http.MethodPost]
		var user, pass string
		for k, v := range auth {
			user = k
			pass = v
		}
		req.SetBasicAuth(user, pass)
	}

	rr, err := c.client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	return rr, nil
}

func (c *HttpClient) PostWithResponse(addrs string, payload []byte) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodPost, addrs, bytes.NewBuffer(payload))
	if err != nil {
		return &http.Response{}, err
	}
	if len(c.headers[http.MethodPost]) > 0 {
		for k, v := range c.headers[http.MethodPost] {
			req.Header.Set(k, v)
		}
	}

	if len(c.setBasicAuth[http.MethodPost]) > 0 {
		auth := c.setBasicAuth[http.MethodPost]
		var user, pass string
		for k, v := range auth {
			user = k
			pass = v
		}
		req.SetBasicAuth(user, pass)
	}

	rr, err := c.client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return rr, nil
}

func (c *HttpClient) DeleteWithResponse(addrs string, payload []byte) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodDelete, addrs, bytes.NewBuffer(payload))
	if err != nil {
		return &http.Response{}, err
	}
	if len(c.headers[http.MethodDelete]) > 0 {
		for k, v := range c.headers[http.MethodDelete] {
			req.Header.Set(k, v)
		}
	}
	if len(c.setBasicAuth[http.MethodPost]) > 0 {
		auth := c.setBasicAuth[http.MethodPost]
		var user, pass string
		for k, v := range auth {
			user = k
			pass = v
		}
		req.SetBasicAuth(user, pass)
	}

	rr, err := c.client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return rr, nil
}

func (c *HttpClient) PostFormWithResponse(addrs string) (*http.Response, error) {
	forms := url.Values{}
	if len(c.forms[http.MethodPost]) > 0 {
		for k, v := range c.forms[http.MethodPost] {
			forms.Add(k, v)
		}
	}

	req, err := http.NewRequest(http.MethodPost, addrs, strings.NewReader(forms.Encode()))
	if err != nil {
		return &http.Response{}, err
	}
	if len(c.headers[http.MethodPost]) > 0 {
		for k, v := range c.headers[http.MethodPost] {
			req.Header.Set(k, v)
		}
	}
	if len(c.setBasicAuth[http.MethodPost]) > 0 {
		auth := c.setBasicAuth[http.MethodPost]
		var user, pass string
		for k, v := range auth {
			user = k
			pass = v
		}
		req.SetBasicAuth(user, pass)
	}

	rr, err := c.client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return rr, nil

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
	v, ok := c.setBasicAuth[strings.ToUpper(method)]
	if ok {
		_, ook := v[username]
		if !ook {
			v[username] = password
		}
	} else {
		c.setBasicAuth[strings.ToUpper(method)] = make(map[string]string)
		c.setBasicAuth[strings.ToUpper(method)][username] = password
	}
}

func (c *HttpClient) GetBasicAuth(method string) map[string]string {
	c.Lock()
	defer c.Unlock()
	return c.setBasicAuth[strings.ToUpper(method)]
}
