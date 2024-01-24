package revenuecat

import (
	"net/http"
)

const (
	transactionsPath = "/v1/receipts"
)

type CreateAPurchaseRequest struct {
	AppUserId                   string      `json:"app_user_id,omitempty"`
	FetchToken                  string      `json:"fetch_token,omitempty"`
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
