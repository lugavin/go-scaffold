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

	"github.com/lugavin/go-scaffold/pkg/log"
)

var (
	ErrCustomHttpClientNotSet = errors.New("custom HTTP client not set")
	ErrHttpStatusNotOK        = errors.New("http response not OK")
)

type HttpWrapper interface {
	Do(ctx context.Context, method, url string, header http.Header, body io.Reader) (io.ReadCloser, error)
}

type Client struct {
	debug           bool
	useCustomClient bool
	logger          log.Logger
	client          *http.Client
	customClient    HttpWrapper
}

func NewHttpClient(logger log.Logger, timeout time.Duration) *Client {
	return &Client{
		debug:           true,
		useCustomClient: false,
		logger:          logger,
		client:          &http.Client{Timeout: timeout},
	}
}

func (c Client) JSONRequest(ctx context.Context, method, url string, header http.Header, input, output interface{}) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		c.logger.Error(fmt.Errorf("encode JSON failed: %w", err))
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

func (c Client) GetRequest(ctx context.Context, url string, input, output interface{}) error {
	return c.JSONRequest(ctx, "GET", url, nil, input, output)
}

func (c Client) PostRequest(ctx context.Context, url string, input, output interface{}) error {
	return c.JSONRequest(ctx, "POST", url, nil, input, output)
}

func (c Client) DoRequest(ctx context.Context, method, url string, header http.Header, body io.Reader, output interface{}) error {
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
			c.logger.Error(fmt.Errorf("request %s error: %w", url, err))
			return err
		}
	} else {
		req, err := http.NewRequestWithContext(ctx, method, url, body)
		if err != nil {
			c.logger.Error(fmt.Errorf("init request %s failed: %w", url, err))
			return err
		}
		req.Header = header
		if c.debug {
			b, _ := httputil.DumpRequest(req, true)
			fmt.Println(string(b))
		}
		resp, err := c.client.Do(req)
		if err != nil {
			c.logger.Error(fmt.Errorf("request %s error: %w", url, err))
			return err
		}
		//if c.debug {
		//	b, _ := httputil.DumpResponse(resp, true)
		//	fmt.Println(string(b))
		//}
		if resp.StatusCode != http.StatusOK {
			c.logger.Error(fmt.Errorf("request %s response not OK: %d", url, resp.StatusCode))
			return ErrHttpStatusNotOK
		}
		respBody = resp.Body
	}
	defer respBody.Close()

	if err = json.NewDecoder(respBody).Decode(&output); err != nil {
		c.logger.Error(fmt.Errorf("decode body failed: %w", err))
		return err
	}
	return nil
}
