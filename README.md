
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

[https://pkg.go.dev/github.com/pkg6/go-revenuecat](https://pkg.go.dev/github.com/pkg6/go-revenuecat)

