package httpc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type HttpClient struct {
	sync.RWMutex
	client       *http.Client
	headers      map[string]map[string]string
	forms        map[string]map[string]string
	basicAuth    map[string]map[string]string
	params       *HttpClientParams
}

func NewHttpClient(opts ...HttpClientOptions) *HttpClient {

	params := newHttpClientParams(opts...)

	client := &http.Client{}

	return &HttpClient{
		client:    client,
		headers:   make(map[string]map[string]string),
		forms:     make(map[string]map[string]string),
		basicAuth: make(map[string]map[string]string),
		params:    params,
	}
}

func (c *HttpClient) setHeaders(method string, req *http.Request) {
	c.RLock()
	defer c.RUnlock()
	headers, exists := c.headers[methodKey(method)]
	if exists {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
}

func (c *HttpClient) setBasicAuth(method string, req *http.Request) {
	c.RLock()
	defer c.RUnlock()
	auth, exists := c.basicAuth[methodKey(method)]
	if exists {
		for username, password := range auth {
			req.SetBasicAuth(username, password)
		}
	}
}

func (c *HttpClient) doRequest(method, addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequestWithContext(context.Background(), method, addrs, payload)
}

func (c *HttpClient) doRequestWithContext(ctx context.Context, method, addrs string, payload []byte) (*http.Response, []byte, error) {
	resp, body, err := c.doRequestWithContextRaw(ctx, method, addrs, payload, true)
	return resp, body, err
}

func (c *HttpClient) buildRequest(ctx context.Context, method, addrs string, payload []byte) (*http.Request, error) {
	formValues := c.formValuesFor(method)
	if len(formValues) > 0 {
		form := url.Values{}
		for key, value := range formValues {
			form.Add(key, value)
		}
		req, err := http.NewRequestWithContext(ctx, method, addrs, strings.NewReader(form.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		return req, nil
	}

	req, err := http.NewRequestWithContext(ctx, method, addrs, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	if len(payload) > 0 && (method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch) {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}

func (c *HttpClient) Get(addrs string) (*http.Response, []byte, error) {
	return c.doRequest(http.MethodGet, addrs, nil)
}

func (c *HttpClient) Post(addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequest(http.MethodPost, addrs, payload)
}

func (c *HttpClient) Put(addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequest(http.MethodPut, addrs, payload)
}

func (c *HttpClient) Delete(addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequest(http.MethodDelete, addrs, payload)
}

func (c *HttpClient) Patch(addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequest(http.MethodPatch, addrs, payload)
}

func (c *HttpClient) Head(addrs string) (*http.Response, []byte, error) {
	return c.HeadWithContext(context.Background(), addrs)
}

func (c *HttpClient) GetWithContext(ctx context.Context, addrs string) (*http.Response, []byte, error) {
	return c.doRequestWithContext(ctx, http.MethodGet, addrs, nil)
}

func (c *HttpClient) PostWithContext(ctx context.Context, addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequestWithContext(ctx, http.MethodPost, addrs, payload)
}

func (c *HttpClient) PutWithContext(ctx context.Context, addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequestWithContext(ctx, http.MethodPut, addrs, payload)
}

func (c *HttpClient) DeleteWithContext(ctx context.Context, addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequestWithContext(ctx, http.MethodDelete, addrs, payload)
}

func (c *HttpClient) PatchWithContext(ctx context.Context, addrs string, payload []byte) (*http.Response, []byte, error) {
	return c.doRequestWithContext(ctx, http.MethodPatch, addrs, payload)
}

func (c *HttpClient) HeadWithContext(ctx context.Context, addrs string) (*http.Response, []byte, error) {
	resp, body, err := c.doRequestWithContextRaw(ctx, http.MethodHead, addrs, nil, true)
	return resp, body, err
}

func (c *HttpClient) SetHeader(method, key, value string) {
	c.Lock()
	defer c.Unlock()
	method = methodKey(method)

	if _, exists := c.headers[method]; !exists {
		c.headers[method] = make(map[string]string)
	}
	c.headers[method][key] = value
}

func (c *HttpClient) DeleteHeader(method, key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.headers[methodKey(method)], key)
}

func (c *HttpClient) SetFormValue(method, key, value string) {
	c.Lock()
	defer c.Unlock()
	method = methodKey(method)

	if _, exists := c.forms[method]; !exists {
		c.forms[method] = make(map[string]string)
	}
	c.forms[method][key] = value
}

func (c *HttpClient) DeleteFormValue(method, key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.forms[methodKey(method)], key)
}

func (c *HttpClient) GetFormValue(method string) map[string]string {
	c.RLock()
	defer c.RUnlock()
	return cloneStringMap(c.forms[methodKey(method)])
}

func (c *HttpClient) SetBasicAuth(method, username, password string) {
	c.Lock()
	defer c.Unlock()
	method = methodKey(method)

	c.basicAuth[method] = map[string]string{
		username: password,
	}
}

func (c *HttpClient) GetBasicAuth(method string) map[string]string {
	c.RLock()
	defer c.RUnlock()
	return cloneStringMap(c.basicAuth[methodKey(method)])
}

func (c *HttpClient) SetPatchHeader(key, value string) {
	c.Lock()
	defer c.Unlock()

	if _, exists := c.headers[http.MethodPatch]; !exists {
		c.headers[http.MethodPatch] = make(map[string]string)
	}

	c.headers[http.MethodPatch][key] = value
}

func (c *HttpClient) GetHeaders(method string) map[string]string {
	c.RLock()
	defer c.RUnlock()
	return cloneStringMap(c.headers[methodKey(method)])
}

func (c *HttpClient) formValuesFor(method string) map[string]string {
	c.RLock()
	defer c.RUnlock()
	return cloneStringMap(c.forms[methodKey(method)])
}

func cloneStringMap(values map[string]string) map[string]string {
	if len(values) == 0 {
		return nil
	}
	clone := make(map[string]string, len(values))
	for k, v := range values {
		clone[k] = v
	}
	return clone
}

func methodKey(method string) string {
	return strings.ToUpper(method)
}

func (c *HttpClient) retriesForMethod(method string) int {
	if c.params == nil {
		return 1
	}
	key := methodKey(method)
	if c.params.MethodRetries != nil {
		if retries, ok := c.params.MethodRetries[key]; ok {
			return retries
		}
		if retries, ok := c.params.MethodRetries[method]; ok {
			return retries
		}
	}
	return c.params.MaxRetries
}

func (c *HttpClient) shouldRetryStatus(statusCode int) bool {
	if c.params == nil || c.params.RetryStatusCodes == nil {
		return false
	}
	_, ok := c.params.RetryStatusCodes[statusCode]
	return ok
}

func (c *HttpClient) shouldRetryError(err error) bool {
	if err == nil {
		return false
	}
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	return false
}

func (c *HttpClient) sleepRetry(ctx context.Context) error {
	if c.params == nil || c.params.MaxRetryWait <= 0 {
		return nil
	}
	wait := time.Second * time.Duration(c.params.MaxRetryWait)
	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (c *HttpClient) doRequestWithContextRaw(ctx context.Context, method, addrs string, payload []byte, readBody bool) (*http.Response, []byte, error) {
	attempts := c.retriesForMethod(method)
	if attempts < 1 {
		attempts = 1
	}

	for attempt := 0; attempt < attempts; attempt++ {
		req, err := c.buildRequest(ctx, method, addrs, payload)
		if err != nil {
			return nil, nil, fmt.Errorf("creating request failed: %w", err)
		}

		c.setHeaders(method, req)
		c.setBasicAuth(method, req)
		c.applyRequestHooks(req)

		resp, err := c.client.Do(req)
		if err != nil {
			if c.shouldRetryError(err) && attempt+1 < attempts {
				if sleepErr := c.sleepRetry(ctx); sleepErr != nil {
					return nil, nil, sleepErr
				}
				continue
			}
			return nil, nil, fmt.Errorf("request failed: %w", err)
		}

		if resp.StatusCode >= 400 {
			if c.shouldRetryStatus(resp.StatusCode) && attempt+1 < attempts {
				io.Copy(io.Discard, resp.Body)
				resp.Body.Close()
				if sleepErr := c.sleepRetry(ctx); sleepErr != nil {
					return nil, nil, sleepErr
				}
				continue
			}
			bts, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			return nil, nil, fmt.Errorf("http error: status(%d) %s", resp.StatusCode, string(bts))
		}

		if !readBody {
			return resp, nil, nil
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, nil, fmt.Errorf("reading response failed: %w", err)
		}
		return resp, body, nil
	}

	return nil, nil, fmt.Errorf("request failed: exhausted retries")
}

func (c *HttpClient) applyRequestHooks(req *http.Request) {
	if c.params == nil || len(c.params.RequestHooks) == 0 {
		return
	}
	for _, hook := range c.params.RequestHooks {
		if hook == nil {
			continue
		}
		hook(req)
	}
}
