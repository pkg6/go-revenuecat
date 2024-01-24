package revenuecat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	PlatformIOS         = "ios"
	PlatformAndroid     = "android"
	PlatformAmazon      = "amazon"
	PlatformMacos       = "macos"
	PlatformUikitformac = "uikitformac"

	defaultBaseURL = "https://api.revenuecat.com/"
)

const (
	ApiKeyV1 = "v1"
	ApiKeyV2 = "v2"
)

type Client struct {
	//https://www.revenuecat.com/reference/revenuecat-rest-api
	ApiKeyV1 string
	//https://www.revenuecat.com/reference/revenuecat-rest-api
	ApiKeyV2 string
	//https://app.revenuecat.com/projects/{ProjectID}/apps/{RevenueCatAppID}
	PublicAPIKey string

	BaseURL *url.URL
	Client  *http.Client
}

func (c *Client) CallKeyVersion(keyVersion, method, path string, reqBody any, platform string, respBody any) error {
	var apiKey string
	switch keyVersion {
	case ApiKeyV1:
		apiKey = c.ApiKeyV1
	case ApiKeyV2:
		apiKey = c.ApiKeyV2
	default:
		apiKey = c.PublicAPIKey
	}
	return c.call(apiKey, method, path, reqBody, platform, respBody)
}

func (c *Client) call(apiKey, method, path string, reqBody any, platform string, respBody any) error {
	req, err := c.NewRequest(apiKey, method, path, platform, reqBody)
	if err != nil {
		return err
	}
	_, err = c.Do(req, respBody)
	return err
}

func (c *Client) NewRequest(apiKey, method, urlStr, platform string, reqBody any) (*http.Request, error) {
	if c.BaseURL == nil {
		c.BaseURL, _ = url.Parse(defaultBaseURL)
	}

	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	var reqBodyJSON io.Reader
	if reqBody != nil {
		js, err := json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal body: %v", err)
		}
		reqBodyJSON = bytes.NewBuffer(js)
	}
	req, err := http.NewRequest(method, u.String(), reqBodyJSON)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	if platform != "" {
		req.Header.Add("X-Platform", platform)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, d any) (resp *http.Response, err error) {
	if c.Client == nil {
		c.Client = http.DefaultClient
	}
	resp, err = c.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	errorResponse := &RespError{Response: resp}
	bodyDataByte, err := io.ReadAll(resp.Body)
	if err == nil && bodyDataByte != nil {
		if err := jsonUnmarshal(bodyDataByte, &errorResponse); err != nil {
			return resp, &RespError{Response: resp}
		}
	}
	switch {
	case resp.StatusCode > 199 && resp.StatusCode < 300:
		switch v := d.(type) {
		case nil:
		case io.Writer:
			_, err = io.Copy(v, resp.Body)
		default:
			err = jsonUnmarshal(bodyDataByte, d)
		}
		return
	default:
		err = errorResponse
		return
	}
}
