package revenuecat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	PlatformIOS         = "ios"
	PlatformAndroid     = "android"
	PlatformAmazon      = "amazon"
	PlatformMacos       = "macos"
	PlatformUikitformac = "uikitformac"

	Host = "https://api.revenuecat.com"
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
	HttpClient   *http.Client
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
	if c.HttpClient == nil {
		c.HttpClient = http.DefaultClient
	}
	var reqBodyJSON io.Reader
	if reqBody != nil {
		js, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("json.Marshal body: %v", err)
		}
		reqBodyJSON = bytes.NewBuffer(js)
	}
	req, err := http.NewRequest(method, Host+path, reqBodyJSON)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	if platform != "" {
		req.Header.Add("X-Platform", platform)
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("c.http.Do: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		respErr := new(RespError)
		err = json.NewDecoder(resp.Body).Decode(&respErr)
		if err != nil {
			return err
		}
		return respErr
	}
	err = json.NewDecoder(resp.Body).Decode(respBody)
	if err != nil {
		return fmt.Errorf("json.NewDecoder.Decode: %v", err)
	}
	return nil
}
