package revenuecat

import (
	"fmt"
	"net/http"
)

const (
	transactionsPath                          = "/v1/receipts"
	googlePlayRefundAndRevokeSubscriptionPath = "/v1/subscribers/%s/subscriptions/%s/revoke"
	googlePlayDeferASubscriptionPath          = "/v1/subscribers/%s/subscriptions/%s/defer"
	googlePlayRefundAndRevokePurchasePath     = "/v1/subscribers/%s/transactions/%s/refund"
	googlePlayCancelASubscriptionPath         = " /v1/subscribers/%s/subscriptions/%s/cancel"
)

type CreateAPurchaseRequest struct {
	AppUserId                   string      `json:"app_user_id"`
	FetchToken                  string      `json:"fetch_token"`
	ProductId                   string      `json:"product_id,omitempty"`
	Price                       float64     `json:"price,omitempty"`
	Currency                    string      `json:"currency,omitempty"`
	IsRestore                   bool        `json:"is_restore,omitempty"`
	PaymentMode                 int         `json:"payment_mode,omitempty"`
	CreateEvents                bool        `json:"create_events,omitempty"`
	ShouldUpdateLastSeenFields  bool        `json:"should_update_last_seen_fields,omitempty"`
	SubscriptionGroupId         string      `json:"subscription_group_id,omitempty"`
	StoreUserId                 string      `json:"store_user_id,omitempty"`
	PresentedOfferingIdentifier string      `json:"presented_offering_identifier,omitempty"`
	ObserverMode                bool        `json:"observer_mode,omitempty"`
	IntroductoryPrice           float64     `json:"introductory_price,omitempty"`
	NormalDuration              string      `json:"normal_duration,omitempty"`
	IntroDuration               string      `json:"intro_duration,omitempty"`
	TrialDuration               string      `json:"trial_duration,omitempty"`
	StoreCountry                string      `json:"store_country,omitempty"`
	IpAddress                   string      `json:"ip_address,omitempty"`
	Attributes                  *Attributes `json:"attributes,omitempty"`
}
type Attributes struct {
	ProductId  string             `json:"product_id,omitempty"`
	Price      float64            `json:"price,omitempty"`
	Currency   string             `json:"currency,omitempty"`
	IsRestore  string             `json:"is_restore,omitempty"`
	Attributes *AttributesKeyName `json:"attributes,omitempty"`
	AppUserId  string             `json:"app_user_id,omitempty"`
	FetchToken string             `json:"fetch_token,omitempty"`
}
type AttributesKeyName struct {
	KeyName *KeyName `json:"key_name,omitempty"`
}
type KeyName struct {
	Value       string `json:"value,omitempty"`
	UpdatedAtMs string `json:"updated_at_ms,omitempty"`
}

// CreateAPurchase
//https://www.revenuecat.com/reference/receipts
func (c *Client) CreateAPurchase(req *CreateAPurchaseRequest, platforms ...string) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	platform := PlatformIOS
	if len(platforms) > 0 {
		platform = platforms[0]
	}
	err = c.call(
		c.PublicAPIKey,
		http.MethodPost,
		transactionsPath,
		req,
		platform,
		&resp,
	)
	return
}

// GooglePlayRefundAndRevokeSubscription
//https://www.revenuecat.com/reference/revoke-a-google-subscription
func (c *Client) GooglePlayRefundAndRevokeSubscription(appUserId, productIdentifier string) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	err = c.call(
		c.PublicAPIKey,
		http.MethodPost,
		fmt.Sprintf(googlePlayRefundAndRevokeSubscriptionPath, appUserId, productIdentifier),
		nil,
		"",
		&resp,
	)
	return
}

type GooglePlayDeferASubscriptionRequest struct {
	ExpiryTimeMs int64 `json:"expiry_time_msm"`
}

// GooglePlayDeferASubscription
//https://www.revenuecat.com/reference/defer-a-google-subscription
func (c *Client) GooglePlayDeferASubscription(appUserId, productIdentifier string, req *GooglePlayDeferASubscriptionRequest) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	err = c.call(
		c.PublicAPIKey,
		http.MethodPost,
		fmt.Sprintf(googlePlayDeferASubscriptionPath, appUserId, productIdentifier),
		req,
		"",
		&resp,
	)
	return
}

// GooglePlayRefundAndRevokePurchase
//https://www.revenuecat.com/reference/refund-a-google-purchase
func (c *Client) GooglePlayRefundAndRevokePurchase(appUserId, storeTransactionIdentifier string) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	err = c.call(
		c.PublicAPIKey,
		http.MethodPost,
		fmt.Sprintf(googlePlayRefundAndRevokePurchasePath, appUserId, storeTransactionIdentifier),
		nil,
		"",
		&resp,
	)
	return
}

// GooglePlayCancelASubscription
//https://www.revenuecat.com/reference/cancel-a-google-subscription
func (c *Client) GooglePlayCancelASubscription(appUserId, storeTransactionIdentifier string) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	err = c.call(
		c.PublicAPIKey,
		http.MethodPost,
		fmt.Sprintf(googlePlayCancelASubscriptionPath, appUserId, storeTransactionIdentifier),
		nil,
		"",
		&resp,
	)
	return
}
