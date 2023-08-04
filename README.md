# go-HTTPC - A golang client HTTP (beta)

The `httpc` package provides a set of functions to perform HTTP requests with retry logic. It also allows setting custom headers, form values, and basic authentication for requests.

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

This function creates a new HttpClient with the provided options. It sets up a RetryableTransport that will retry requests a specified number of times with a specified wait time between retries.

### Get(addrs string) ([]byte, error)

This function sends a GET request to the specified address and returns the response body as a byte slice. If the server responds with a 500 status code, it returns an error.

### Post(addrs string, payload []byte) ([]byte, error)

This function sends a POST request to the specified address with the provided payload and returns the response body as a byte slice. If the server responds with a 500 status code, it returns an error.

### Delete(addrs string, payload []byte) ([]byte, error)

This function sends a DELETE request to the specified address with the provided payload and returns the response body as a byte slice. If the server responds with a 500 status code, it returns an error.

### SetHeader(method, key, value string)

This function sets a header for a specified HTTP method. If the header already exists, it updates the value.

### DeleteHeader(method, key string)

This function deletes a header for a specified HTTP method.

### SetFormValue(method, key, value string)

This function sets a form value for a specified HTTP method. If the form value already exists, it updates the value.

### DeleteFormValue(method, key string)

This function deletes a form value for a specified HTTP method.

### GetWithResponse(addrs string) (*http.Response, error)

This function sends a GET request to the specified address and returns the full http.Response object.

### PostWithResponse(addrs string, payload []byte) (*http.Response, error)

This function sends a POST request to the specified address with the provided payload and returns the full http.Response object.

### DeleteWithResponse(addrs string, payload []byte) (*http.Response, error)

This function sends a DELETE request to the specified address with the provided payload and returns the full http.Response object.

### PostFormWithResponse(addrs string) (*http.Response, error)

This function sends a POST request with form values to the specified address and returns the full http.Response object.

### GetHeaders(method string) map[string]string

This function returns all headers set for a specified HTTP method.

### GetFormValue(method string) map[string]string

This function returns all form values set for a specified HTTP method.

### SetBasicAuth(method, username, password string)

This function sets the Basic Authentication for a specified HTTP method.

### GetBasicAuth(method string) map[string]string

This function returns the Basic Authentication set for a specified HTTP method.

-----

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-httpc/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE) file.

**2023**, thiagozs.
