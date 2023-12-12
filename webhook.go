package revenuecat

import (
	"encoding/json"
	"net/http"
)

//https://www.revenuecat.com/docs/sample-events

type WebHookCallback func(resp *WebHookResponse) error

func WebHookWithCallback(r *http.Request, callbacks ...WebHookCallback) error {
	response := new(WebHookResponse)
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return err
	}
	for _, callback := range callbacks {
		if err := callback(response); err != nil {
			return err
		}
	}
	return nil
}

type WebHookResponse struct {
	Event      *WebHookEvent `json:"event"`
	ApiVersion string        `json:"api_version"`
}

type WebHookEvent struct {
	//Initial Purchase
	Id               string `json:"id"`
	Type             string `json:"type"`
	Store            string `json:"store"`
	Environment      string `json:"environment"`
	EventTimestampMs int64  `json:"event_timestamp_ms"`

	ProductId      string `json:"product_id,omitempty"`
	PeriodType     string `json:"period_type,omitempty"`
	PurchasedAtMs  int64  `json:"purchased_at_ms,omitempty"`
	ExpirationAtMs int64  `json:"expiration_at_ms,omitempty"`

	EntitlementId            string         `json:"entitlement_id,omitempty"`
	EntitlementIds           []string       `json:"entitlement_ids,omitempty"`
	PresentedOfferingId      string         `json:"presented_offering_id,omitempty"`
	TransactionId            string         `json:"transaction_id,omitempty"`
	OriginalTransactionId    string         `json:"original_transaction_id,omitempty"`
	IsFamilyShare            bool           `json:"is_family_share,omitempty"`
	CountryCode              string         `json:"country_code,omitempty"`
	AppUserId                string         `json:"app_user_id,omitempty"`
	Aliases                  []string       `json:"aliases,omitempty"`
	OriginalAppUserId        string         `json:"original_app_user_id,omitempty"`
	Currency                 string         `json:"currency,omitempty"`
	Price                    float64        `json:"price,omitempty"`
	PriceInPurchasedCurrency float64        `json:"price_in_purchased_currency,omitempty"`
	SubscriberAttributes     map[string]any `json:"subscriber_attributes,omitempty"`

	TakehomePercentage float64 `json:"takehome_percentage,omitempty"`
	OfferCode          string  `json:"offer_code,omitempty"`

	AppId string `json:"app_id,omitempty,omitempty"`

	//Renewal
	IsTrialConversion bool `json:"is_trial_conversion,omitempty,omitempty"`

	//Cancellation && Trial Cancelled
	CancelReason string `json:"cancel_reason,omitempty"`

	//Uncancellation && Subscription extended
	TaxPercentage float64 `json:"tax_percentage,omitempty"`

	//Refund && Uncancellation && Subscription extended
	CommissionPercentage float64 `json:"commission_percentage,omitempty"`

	//Non Renewing Purchase

	//Subscription Paused
	AutoResumeAtMs int64 `json:"auto_resume_at_ms,omitempty"`

	//Billing Issue

	//Expiration
	ExpirationReason string `json:"expiration_reason,omitempty"`

	//Transfer
	TransferredFrom []string `json:"transferred_from,omitempty"`
	TransferredTo   []string `json:"transferred_to,omitempty"`

	//Product Change
	NewProductId string `json:"new_product_id,omitempty"`
	//Subscription extended

	//Trial Started

	//Trial Cancelled
}

func (w *WebHookEvent) SubscriberAttributesMap() map[string]*WebHookSubscriberAttributes {
	ret := map[string]*WebHookSubscriberAttributes{}
	for i, a := range w.SubscriberAttributes {
		var subscriptions WebHookSubscriberAttributes
		_ = JsonUnmarshal(a, &subscriptions)
		ret[i] = &subscriptions
	}
	return ret
}

type WebHookSubscriberAttributes struct {
	UpdatedAtMs int64  `json:"updated_at_ms"`
	Value       string `json:"value"`
}
