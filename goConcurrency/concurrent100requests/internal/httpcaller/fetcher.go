package httpcaller

import (
	"io"
	"net/http"
	"time"
)

type HttpCallerStruct struct {
	client *http.Client
}

func NewHttpCaller() *HttpCallerStruct {
	return &HttpCallerStruct{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *HttpCallerStruct) HTTPSendRequest(requestURL, method, bodyType string, body any, headers map[string]string) (any, error) {
	resp, err := h.client.Get(requestURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
