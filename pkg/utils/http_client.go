package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"go.uber.org/zap"
)

var (
	ErrCustomHttpClientNotSet = errors.New("custom HTTP client not set")
	ErrHttpStatusNotOK        = errors.New("http response not OK")
)

type HttpWrapper interface {
	Do(ctx context.Context, method, url string, header http.Header, body io.Reader) (io.ReadCloser, error)
}

type HttpClient struct {
	debug           bool
	useCustomClient bool
	logger          *zap.Logger
	client          *http.Client
	customClient    HttpWrapper
}

func NewHttpClient(logger *zap.Logger, timeout time.Duration) *HttpClient {
	return &HttpClient{
		debug:           true,
		useCustomClient: false,
		logger:          logger,
		client:          &http.Client{Timeout: timeout},
	}
}

func (c HttpClient) JsonRequest(ctx context.Context, method, url string, header http.Header, input, output interface{}) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		c.logger.Error("encode JSON failed", zap.Error(err))
		return err
	}
	var httpHeader http.Header
	if header == nil {
		httpHeader = make(http.Header)
	} else {
		httpHeader = header.Clone()
	}
	httpHeader.Set("Content-Type", "application/json")
	return c.DoRequest(ctx, method, url, httpHeader, buf, output)
}

func (c HttpClient) GetRequest(ctx context.Context, url string, input, output interface{}) error {
	return c.JsonRequest(ctx, "GET", url, nil, input, output)
}

func (c HttpClient) PostRequest(ctx context.Context, url string, input, output interface{}) error {
	return c.JsonRequest(ctx, "POST", url, nil, input, output)
}

func (c HttpClient) DoRequest(ctx context.Context, method, url string, header http.Header, body io.Reader, output interface{}) error {
	logger := c.logger.With(zap.String("url", url))

	var (
		err      error
		respBody io.ReadCloser
	)
	if header == nil {
		header = make(http.Header)
	}

	if c.useCustomClient {
		if c.customClient == nil {
			return ErrCustomHttpClientNotSet
		}
		respBody, err = c.customClient.Do(ctx, method, url, header, body)
		if err != nil {
			logger.Error("http request error", zap.Error(err))
			return err
		}
	} else {
		req, err := http.NewRequestWithContext(ctx, method, url, body)
		if err != nil {
			logger.Error("init request failed", zap.Error(err))
			return err
		}
		req.Header = header
		if c.debug {
			b, _ := httputil.DumpRequest(req, true)
			fmt.Println(string(b))
		}
		resp, err := c.client.Do(req)
		if err != nil {
			logger.Error("http request error", zap.Error(err))
			return err
		}
		//if c.debug {
		//	b, _ := httputil.DumpResponse(resp, true)
		//	fmt.Println(string(b))
		//}
		if resp.StatusCode != http.StatusOK {
			logger.Error("http response not OK", zap.Int("statusCode", resp.StatusCode))
			return ErrHttpStatusNotOK
		}
		respBody = resp.Body
	}
	defer respBody.Close()

	if err = json.NewDecoder(respBody).Decode(&output); err != nil {
		logger.Error("decode body failed", zap.Error(err))
		return err
	}
	logger.Info("[DumpResponse]", zap.Any("output", output))
	return nil
}
