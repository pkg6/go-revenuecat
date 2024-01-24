package revenuecat

import (
	"encoding/json"
	"io"
	"net/url"
	"time"
)

func JsonUnmarshal(v, s any) error {
	jsonB, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonB, s)
}

func PathQuery(path string, query map[string]string) string {
	u, _ := url.Parse(path)
	v := u.Query()
	for s, s2 := range query {
		if s2 != "" {
			v.Set(s, s2)
		}
	}
	u.RawQuery = v.Encode()
	return u.String()
}

func IsExpired(timestamp int64) bool {
	expiration := time.Unix(timestamp, 0)
	now := time.Now()
	if now.After(expiration) {
		return true
	}
	return false
}

func jsonUnmarshal(data []byte, v any) error {
	decErr := json.Unmarshal(data, v)
	if decErr == io.EOF {
		return nil
	}
	return decErr
}
