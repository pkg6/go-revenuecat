package revenuecat

import (
	"fmt"
	"time"
)

type RespError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *RespError) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Message)
}

type SubscriberResponse struct {
	RequestDate   time.Time   `json:"request_date"`
	RequestDateMs int64       `json:"request_date_ms"`
	Subscriber    *Subscriber `json:"subscriber"`
}

func (s *SubscriberResponse) IsExpiredEntitlement(entitlement string) bool {
	entitlementsMap := s.Subscriber.EntitlementsMap()
	if en, ok := entitlementsMap[entitlement]; ok {
		return IsExpired(en.ExpiresDate.Unix())
	}
	return false
}

type Subscriber struct {
	FirstSeen                  time.Time         `json:"first_seen"`
	LastSeen                   time.Time         `json:"last_seen"`
	ManagementUrl              string            `json:"management_url"`
	NonSubscriptions           *NonSubscriptions `json:"non_subscriptions"`
	OriginalAppUserId          string            `json:"original_app_user_id"`
	OriginalApplicationVersion string            `json:"original_application_version"`
	OriginalPurchaseDate       time.Time         `json:"original_purchase_date"`
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
	ExpiresDate             time.Time `json:"expires_date"`
	PurchaseDate            time.Time `json:"purchase_date"`
	OriginalPurchaseDate    time.Time `json:"original_purchase_date"`
	OwnershipType           string    `json:"ownership_type"`
	PeriodType              string    `json:"period_type"`
	Store                   string    `json:"store"`
	IsSandbox               string    `json:"is_sandbox"`
	UnsubscribeDetectedAt   time.Time `json:"unsubscribe_detected_at"`
	BillingIssuesDetectedAt time.Time `json:"billing_issues_detected_at"`
	GracePeriodExpiresDate  time.Time `json:"grace_period_expires_date"`
	RefundedAt              time.Time `json:"refunded_at"`
	AutoResumeDate          time.Time `json:"auto_resume_date"`
	StoreTransactionId      time.Time `json:"store_transaction_id"`
}

type Entitlements struct {
	ExpiresDate            time.Time `json:"expires_date"`
	GracePeriodExpiresDate time.Time `json:"grace_period_expires_date"`
	PurchaseDate           time.Time `json:"purchase_date"`
	ProductIdentifier      string    `json:"product_identifier"`
}

type Amazon struct {
	PackageName string `json:"package_name"`
}
type MacAppStore struct {
	BundleID string `json:"bundle_id"`
}
type PlayStore struct {
	PackageName string `json:"package_name"`
}
type AppStore struct {
	BundleId string `json:"bundle_id"`
}
