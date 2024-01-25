
## install

```
go get github.com/pkg6/go-revenuecat
```

## Basic Use

```
client := &revenuecat.Client{
  ApiKeyV1:     "sk_*******************",
  ApiKeyV2:     "sk_*******************",
  PublicAPIKey: "appl_*******************",
}
client.GetOrCreateSubscriber("dev-0")
```

## Implementation List

| Official interface address                                   | Function                      |
| ------------------------------------------------------------ | ----------------------------- |
| [Get or Create Subscriber](https://www.revenuecat.com/reference/subscribers) | GetOrCreateSubscriber         |
| [Delete Subscriber](https://www.revenuecat.com/reference/delete-subscriber) | DeleteSubscriber              |
| [Get a list of projects](https://www.revenuecat.com/reference/list-projects) | GetAListOfProjects            |
| [Create a Purchase](https://www.revenuecat.com/reference/receipts) | CreateAPurchase               |
| [Get a list of apps](https://www.revenuecat.com/reference/list-apps) | GetAListOfApps                |
| [Get an app](https://www.revenuecat.com/reference/get-app)   | GetAnApp                      |
| [Grant a Promotional Entitlement](https://www.revenuecat.com/reference/grant-a-promotional-entitlement) | GrantAPromotionalEntitlement  |
| [Revoke Promotional Entitlements](https://www.revenuecat.com/reference/revoke-promotional-entitlements) | RevokePromotionalEntitlements |

