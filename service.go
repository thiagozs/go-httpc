package httpc

import "net/http"

type HttpClientDriver struct {
	client HttpClientRepo
}

func NewHttpClientDriver(client HttpClientRepo) *HttpClientDriver {
	return &HttpClientDriver{
		client: client,
	}
}

func (d HttpClientDriver) Get(addrs string) ([]byte, error) {
	return d.client.Get(addrs)
}

func (d HttpClientDriver) GetWithResponse(addrs string) (*http.Response, error) {
	return d.client.GetWithResponse(addrs)
}

func (d HttpClientDriver) Post(addrs string, payload []byte) ([]byte, error) {
	return d.client.Post(addrs, payload)
}

func (d HttpClientDriver) PostWithResponse(addrs string, payload []byte) (*http.Response, error) {
	return d.client.PostWithResponse(addrs, payload)
}

func (d HttpClientDriver) SetHeader(method, key, value string) {
	d.client.SetHeader(method, key, value)
}

func (d HttpClientDriver) SetFormValue(method, key, value string) {
	d.client.SetFormValue(method, key, value)
}

func (d HttpClientDriver) SetBasicAuth(method, username, password string) {
	d.client.SetBasicAuth(method, username, password)
}

func (d HttpClientDriver) DeleteHeader(method, key string) {
	d.client.DeleteHeader(method, key)
}

func (d HttpClientDriver) DeleteFormValue(method, key string) {
	d.client.DeleteFormValue(method, key)
}

func (d HttpClientDriver) GetHeaders(method string) map[string]string {
	return d.client.GetHeaders(method)
}

func (d HttpClientDriver) GetFormValue(method string) map[string]string {
	return d.client.GetFormValue(method)
}

func (d HttpClientDriver) GetBasicAuth(method string) map[string]string {
	return d.client.GetBasicAuth(method)
}

func (d HttpClientDriver) DisableLogLevel() {
	d.client.DisableLogLevel()
}

func (d HttpClientDriver) EnableLogLevel() {
	d.client.EnableLogLevel()
}

func (d HttpClientDriver) GetFreePort() (int, error) {
	return d.client.GetFreePort()
}

func (d HttpClientDriver) Delete(addrs string, payload []byte) ([]byte, error) {
	return d.client.Delete(addrs, payload)
}

func (d HttpClientDriver) DeleteWithResponse(addrs string, payload []byte) (*http.Response, error) {
	return d.client.DeleteWithResponse(addrs, payload)
}

func (d HttpClientDriver) PostFormWithResponse(addrs string) (*http.Response, error) {
	return d.client.PostFormWithResponse(addrs)
}
