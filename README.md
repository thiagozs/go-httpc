# go-HTTPC - A Golang HTTP Client (Beta)

The `httpc` package is a robust HTTP client for Go, equipped with features such as retry logic, custom headers, form values, and basic authentication settings. It utilizes a retryable transport structure to manage request retries effectively, enhancing the reliability of HTTP requests in Go applications.

```golang

    opts := []httpc.HttpClientOptions{
    httpc.WithMaxRetries(3),
    httpc.WithMaxRetryWait(3), //seconds
    }

    client := httpc.NewHttpClient(opts...)

    URL : = "https://thiagozs.com"

    resp, err := client.Get(URL)
    if err != nil {
        fmt.Printf("error: %+v\n",err)
        return
    }

```

## Features

### Structs

- **RetryableTransport**: Manages the retry logic for HTTP requests with configurable retry counts and wait times.
- **HttpClient**: A thread-safe HTTP client with methods to perform various HTTP requests and manage request headers, form values, and basic authentication settings.

### Methods

#### NewHttpClient(opts ...HttpClientOptions) *HttpClient

Initializes a new HttpClient instance with specified options.

#### Get(addrs string) ([]byte, error)

Sends a GET request to the specified address and returns the response body as a byte slice.

#### Post(addrs string, payload []byte) ([]byte, error)

Sends a POST request to the specified address with the provided payload and returns the response body as a byte slice.

#### Put(addrs string, payload []byte) ([]byte, error)

Sends a PUT request to the specified address with the provided payload and returns the response body as a byte slice.

#### Delete(addrs string, payload []byte) ([]byte, error)

Sends a DELETE request to the specified address with the provided payload and returns the response body as a byte slice.

#### Patch(addrs string, payload []byte) ([]byte, error)

Sends a PATCH request to the specified address with the provided payload and returns the response body as a byte slice.

#### Head(addrs string) (*http.Response, error)

Sends a HEAD request to the specified address and returns the full http.Response object.

#### SetHeader(method, key, value string)

Sets or updates a header for a specified HTTP method.

#### DeleteHeader(method, key string)

Deletes a header for a specified HTTP method.

#### SetFormValue(method, key, value string)

Sets or updates a form value for a specified HTTP method.

#### DeleteFormValue(method, key string)

Deletes a form value for a specified HTTP method.

#### GetFormValue(method string) map[string]string

Retrieves all form values set for a specified HTTP method.

#### SetBasicAuth(method, username, password string)

Sets the Basic Authentication for a specified HTTP method.

#### GetBasicAuth(method string) map[string]string

Retrieves the Basic Authentication set for a specified HTTP method.

#### SetPatchHeader(key, value string)

Sets or updates a header specifically for PATCH requests.

#### GetHeaders(method string) map[string]string

Retrieves all headers set for a specified HTTP method.

## Usage

Refer to the source code for detailed usage and examples.

## Versioning and License

Our version numbers adhere to the semantic versioning specification. You can explore the available versions by checking the tags on this repository. For more details about our license model, please refer to the LICENSE file.

© 2023, thiagozs.
