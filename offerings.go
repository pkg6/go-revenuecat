package revenuecat

import (
	"fmt"
	"net/http"
)

const (
	dverrideACustomerSCurrentOfferingPath       = "/v1/subscribers/%s/offerings/%x/override"
	removeACustomerSCurrentOfferingOverridePath = "/v1/subscribers/%s/offerings/override"
	getOfferingsPath                            = "/v1/subscribers/%s/offerings"
)

// OverrideACustomerSCurrentOffering
//https://www.revenuecat.com/reference/override-offering
func (c *Client) OverrideACustomerSCurrentOffering(appUserId, offeringUuid string) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	err = c.call(
		c.PublicAPIKey,
		http.MethodPost,
		fmt.Sprintf(dverrideACustomerSCurrentOfferingPath, appUserId, offeringUuid),
		nil,
		"",
		&resp,
	)
	return
}

// RemoveACustomerSCurrentOfferingOverride
//https://www.revenuecat.com/reference/delete-offering-override
func (c *Client) RemoveACustomerSCurrentOfferingOverride(appUserId string) (resp *SubscriberResponse, err error) {
	resp = new(SubscriberResponse)
	err = c.call(
		c.PublicAPIKey,
		http.MethodDelete,
		fmt.Sprintf(removeACustomerSCurrentOfferingOverridePath, appUserId),
		nil,
		"",
		&resp,
	)
	return
}

type GetOfferingsResponse struct {
	CurrentOfferingId string              `json:"current_offering_id,omitempty"`
	Offerings         []*OfferingResponse `json:"offerings,omitempty"`
}

type OfferingResponse struct {
	Description string             `json:"description,omitempty"`
	Identifier  string             `json:"identifier,omitempty"`
	Packages    []*OfferingPackage `json:"packages,omitempty"`
}

type OfferingPackage struct {
	Identifier                string `json:"identifier,omitempty"`
	PlatformProductIdentifier string `json:"platform_product_identifier,omitempty"`
}

// GetOfferings
//https://www.revenuecat.com/reference/get-offerings
func (c *Client) GetOfferings(appUserId string, platforms ...string) (resp *GetOfferingsResponse, err error) {
	resp = new(GetOfferingsResponse)
	platform := PlatformIOS
	if len(platforms) > 0 {
		platform = platforms[0]
	}
	err = c.call(
		c.PublicAPIKey,
		http.MethodGet,
		fmt.Sprintf(getOfferingsPath, appUserId),
		nil,
		platform,
		&resp,
	)
	return
}
