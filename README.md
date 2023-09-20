# go-HTTPC - A Golang HTTP Client (Beta)

The `httpc` package provides a set of functions to perform HTTP requests with retry logic. It also allows setting custom headers, form values, and basic authentication for requests, making it a versatile tool for various web projects. It's easy to use and integrates seamlessly with existing Go applications.

Easy to use.

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

### NewHttpClient(opts ...HttpClientOptions) *HttpClient

Creates a new HttpClient instance with the provided options. It sets up a RetryableTransport that will retry requests a specified number of times with a specified wait time between retries.

### Get(addrs string) ([]byte, error)

Sends a GET request to the specified address and returns the response body as a byte slice. If the server responds with a 500 status code, it returns an error.

### Post(addrs string, payload []byte) ([]byte, error)

Sends a POST request to the specified address with the provided payload and returns the response body as a byte slice. If the server responds with a 500 status code, it returns an error.

### Delete(addrs string, payload []byte) ([]byte, error)

Sends a DELETE request to the specified address with the provided payload and returns the response body as a byte slice. If the server responds with a 500 status code, it returns an error.

### SetHeader(method, key, value string)

Sets a header for a specified HTTP method. If the header already exists, it updates the value.

### DeleteHeader(method, key string)

Deletes a header for a specified HTTP method.

### SetFormValue(method, key, value string)

Sets a form value for a specified HTTP method. If the form value already exists, it updates the value.

### DeleteFormValue(method, key string)

Deletes a form value for a specified HTTP method.

### GetHeaders(method string) map[string]string

Returns all headers set for a specified HTTP method.

### GetFormValue(method string) map[string]string

Returns all form values set for a specified HTTP method.

### SetBasicAuth(method, username, password string)

Sets the Basic Authentication for a specified HTTP method.

### GetBasicAuth(method string) map[string]string

Returns the Basic Authentication set for a specified HTTP method.

## Versioning and License

Our version numbers adhere to the semantic versioning specification. You can explore the available versions by checking the tags on this repository. For more details about our license model, please refer to the LICENSE file.

© 2023, thiagozs.
