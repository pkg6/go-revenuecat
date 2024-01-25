package revenuecat

import (
	"fmt"
	"net/http"
)

const (
	grantAPromotionalEntitlementPath  = "/v1/subscribers/%s/entitlements/%s/promotional"
	revokePromotionalEntitlementsPath = "/v1/subscribers/%s/entitlements/%s/revoke_promotionals"
)

type DurationValue string

const (
	// DurationValueDaily
	//24 hour access
	DurationValueDaily DurationValue = "daily"
	// DurationValueThreeDay
	//72 hour access
	DurationValueThreeDay DurationValue = "three_day"
	// DurationValueWeekly
	//7 day access
	DurationValueWeekly DurationValue = "weekly"
	// DurationValueMonthly
	//1 month (31-day) access
	DurationValueMonthly DurationValue = "monthly"
	// DurationValueTwoMonth
	//2 month (61-day) access
	DurationValueTwoMonth DurationValue = "two_month"
	// DurationValueThreeMonth
	//3 month (92-day) access
	DurationValueThreeMonth DurationValue = "three_month"
	// DurationValueSixMonth
	//6 month (183-day) access
	DurationValueSixMonth DurationValue = "six_month"
	// DurationValueYearly
	//1 year (365-day) access
	DurationValueYearly DurationValue = "yearly"
	// DurationValueLifetime
	//200 year access
	DurationValueLifetime DurationValue = "lifetime"
)

type GrantAPromotionalEntitlementRequest struct {
	Duration    DurationValue `json:"duration,omitempty"`
	StartTimeMs int64         `json:"start_time_ms,omitempty"`
}

// GrantAPromotionalEntitlement
//https://www.revenuecat.com/reference/grant-a-promotional-entitlement
func (c *Client) GrantAPromotionalEntitlement(appUserId, entitlementIdentifier string, req *GrantAPromotionalEntitlementRequest) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	err = c.call(
		c.PublicAPIKey,
		http.MethodPost,
		fmt.Sprintf(grantAPromotionalEntitlementPath, appUserId, entitlementIdentifier),
		req,
		"",
		&resp,
	)
	return
}

// RevokePromotionalEntitlements
//https://www.revenuecat.com/reference/revoke-promotional-entitlements
func (c *Client) RevokePromotionalEntitlements(appUserId, entitlementIdentifier string) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	err = c.call(
		c.PublicAPIKey,
		http.MethodPost,
		fmt.Sprintf(revokePromotionalEntitlementsPath, appUserId, entitlementIdentifier),
		nil,
		"",
		&resp,
	)
	return
}
