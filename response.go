package revenuecat

import (
	"fmt"
	"net/http"
	"time"
)

type RespError struct {
	Response  *http.Response
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	DocUrl    string `json:"doc_url,omitempty"`
	Object    string `json:"object,omitempty"`
	Retryable bool   `json:"retryable,omitempty"`
	Type      string `json:"type,omitempty"`
}

func (err *RespError) Error() string {
	if err.Response != nil && err.Response.Request != nil {
		return fmt.Sprintf("%v %v: statusCode=%d code=%d message=%v",
			err.Response.Request.Method,
			err.Response.Request.URL,
			err.Response.StatusCode,
			err.Code,
			err.Message,
		)
	}
	if err.Response != nil {
		return fmt.Sprintf("statusCode=%d code=%d messgae=%v", err.Response.StatusCode, err.Code, err.Message)
	}
	return fmt.Sprintf("code=%d message=%v", err.Code, err.Message)
}

type SubscriberResponse struct {
	RequestDate   *time.Time  `json:"request_date,omitempty"`
	RequestDateMs int64       `json:"request_date_ms,omitempty"`
	Subscriber    *Subscriber `json:"subscriber,omitempty"`
}

func (s *SubscriberResponse) IsExpiredEntitlement(entitlement string) bool {
	entitlementsMap := s.Subscriber.EntitlementsMap()
	if en, ok := entitlementsMap[entitlement]; ok {
		return IsExpired(en.ExpiresDate.Unix())
	}
	return false
}

type Subscriber struct {
	FirstSeen                  *time.Time        `json:"first_seen"`
	LastSeen                   *time.Time        `json:"last_seen"`
	ManagementUrl              string            `json:"management_url"`
	NonSubscriptions           *NonSubscriptions `json:"non_subscriptions"`
	OriginalAppUserId          string            `json:"original_app_user_id"`
	OriginalApplicationVersion string            `json:"original_application_version"`
	OriginalPurchaseDate       *time.Time        `json:"original_purchase_date"`
	OtherPurchases             *NonSubscriptions `json:"other_purchases"`
	Subscriptions              map[string]any    `json:"subscriptions"`
	Entitlements               map[string]any    `json:"entitlements"`
}

type NonSubscriptions struct {
	ID                 string     `json:"id,omitempty"`
	PurchaseDate       *time.Time `json:"purchase_date,omitempty"`
	Store              string     `json:"store,omitempty"`
	IsSandbox          string     `json:"is_sandbox,omitempty"`
	StoreTransactionID string     `json:"store_transaction_id,omitempty"`
}

func (s *Subscriber) SubscriptionsMap() map[string]*Subscriptions {
	ret := map[string]*Subscriptions{}
	for i, a := range s.Subscriptions {
		var j Subscriptions
		_ = JsonUnmarshal(a, &j)
		ret[i] = &j
	}
	return ret
}

func (s *Subscriber) EntitlementsMap() map[string]*Entitlements {
	ret := map[string]*Entitlements{}
	for i, a := range s.Entitlements {
		var j Entitlements
		_ = JsonUnmarshal(a, &j)
		ret[i] = &j
	}
	return ret
}

type Subscriptions struct {
	ExpiresDate             *time.Time `json:"expires_date,omitempty"`
	PurchaseDate            *time.Time `json:"purchase_date,omitempty"`
	OriginalPurchaseDate    *time.Time `json:"original_purchase_date,omitempty"`
	OwnershipType           string     `json:"ownership_type,omitempty"`
	PeriodType              string     `json:"period_type,omitempty"`
	Store                   string     `json:"store,omitempty"`
	IsSandbox               string     `json:"is_sandbox,omitempty"`
	UnsubscribeDetectedAt   *time.Time `json:"unsubscribe_detected_at,omitempty"`
	BillingIssuesDetectedAt *time.Time `json:"billing_issues_detected_at,omitempty"`
	GracePeriodExpiresDate  *time.Time `json:"grace_period_expires_date,omitempty"`
	RefundedAt              *time.Time `json:"refunded_at,omitempty"`
	AutoResumeDate          *time.Time `json:"auto_resume_date,omitempty"`
	StoreTransactionId      *time.Time `json:"store_transaction_id,omitempty"`
}

type Entitlements struct {
	ExpiresDate            *time.Time `json:"expires_date,omitempty"`
	GracePeriodExpiresDate *time.Time `json:"grace_period_expires_date,omitempty"`
	PurchaseDate           *time.Time `json:"purchase_date,omitempty"`
	ProductIdentifier      string     `json:"product_identifier,omitempty"`
}

type Amazon struct {
	PackageName string `json:"package_name,omitempty"`
}
type MacAppStore struct {
	BundleID string `json:"bundle_id,omitempty"`
}
type PlayStore struct {
	PackageName string `json:"package_name,omitempty"`
}
type AppStore struct {
	BundleId string `json:"bundle_id,omitempty"`
}
