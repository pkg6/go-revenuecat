package revenuecat

import (
	"fmt"
	"net/http"
)

//https://www.revenuecat.com/reference/subscribers

const (
	customersPath = "/v1/subscribers/%s"
)

// GetOrCreateSubscriber
//https://www.revenuecat.com/reference/subscribers
func (c *Client) GetOrCreateSubscriber(appUserId string, platforms ...string) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	platform := PlatformIOS
	if len(platforms) > 0 {
		platform = platforms[0]
	}
	err = c.call(c.PublicAPIKey, http.MethodGet, fmt.Sprintf(customersPath, appUserId), nil, platform, &resp)
	return
}

type SubscriberDeleteResponse struct {
	AppUserId string `json:"app_user_id,omitempty"`
	Deleted   bool   `json:"deleted,omitempty"`
}

// DeleteSubscriber
//https://www.revenuecat.com/reference/delete-subscriber
func (c *Client) DeleteSubscriber(appUserId string) (resp *SubscriberDeleteResponse, err error) {
	resp = new(SubscriberDeleteResponse)
	err = c.call(c.ApiKeyV1,
		http.MethodDelete,
		fmt.Sprintf(customersPath, appUserId),
		nil,
		"",
		&resp)
	return
}
