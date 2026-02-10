# go-HTTPC - A Golang HTTP Client (Beta)

`httpc` is a small, thread-safe HTTP client wrapper for Go focused on practical features you typically reimplement in every service. It offers retries, per-method settings, and convenient helpers for headers, form bodies, basic auth, and request hooks while still returning the underlying `*http.Response` for full access to status/headers.

Key behaviors:

- Returns `(*http.Response, []byte, error)` for all methods, so you get headers/status and the body in one call.
- Context-aware variants are available for all HTTP verbs.
- Retries on timeouts with configurable max wait and retries per method.
- Optional retry by status code (e.g., 500/502).
- Form values are encoded as `application/x-www-form-urlencoded` when set for a method.
- JSON `Content-Type` is set automatically for POST/PUT/PATCH when payload is non-empty.
- Basic auth can be configured per method.
- Request hooks let you mutate the request before sending (e.g., add headers, tracing IDs).

```golang
opts := []httpc.HttpClientOptions{
	httpc.WithMaxRetries(3),
	httpc.WithMaxRetryWait(3), // seconds
}

client := httpc.NewHttpClient(opts...)

URL := "https://thiagozs.com"

resp, body, err := client.Get(URL)
if err != nil {
	fmt.Printf("error: %+v\n", err)
	return
}
fmt.Println(resp.StatusCode, string(body))
```

## Usage

### With context and reading headers

```golang
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

resp, body, err := client.GetWithContext(ctx, URL)
if err != nil {
	log.Fatal(err)
}
fmt.Println(resp.StatusCode, resp.Header.Get("Content-Type"))
fmt.Println(string(body))
```

### With request hook

```golang
client := httpc.NewHttpClient(httpc.WithRequestHook(func(req *http.Request) {
	req.Header.Set("X-Request-ID", "abc-123")
}))

resp, body, err := client.Get(URL)
if err != nil {
	log.Fatal(err)
}
fmt.Println(resp.StatusCode, string(body))
```

### Head with context

```golang
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

resp, body, err := client.HeadWithContext(ctx, URL)
if err != nil {
	log.Fatal(err)
}
fmt.Println(resp.StatusCode, len(body))
```

### Retry by status code

```golang
client := httpc.NewHttpClient(
	httpc.WithMethodRetries(http.MethodGet, 3),
	httpc.WithRetryStatusCodes(http.StatusInternalServerError, http.StatusBadGateway),
	httpc.WithMaxRetryWait(1),
)

resp, body, err := client.Get(URL)
if err != nil {
	log.Fatal(err)
}
fmt.Println(resp.StatusCode, string(body))
```

## Versioning and License

Our version numbers adhere to the semantic versioning specification. You can explore the available versions by checking the tags on this repository. For more details about our license model, please refer to the LICENSE file.

© 2023, thiagozs.
