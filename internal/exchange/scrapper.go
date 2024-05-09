package exchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Scrapper struct {
	Endpoint   string
	Passkey    string
	Passphrase string
}

func (s *Scrapper) GetLatestExchangeRates() (map[string]float32, error) {
	m := make(map[string]float32)
	rd := url.Values{}
	rd.Set(s.Passkey, s.Passphrase) // Pretty dumb security ngl
	rr := strings.NewReader(rd.Encode())
	res, err := http.Post(s.Endpoint, "application/x-www-form-urlencoded", rr)
	if err != nil {
		return m, err
	}

	if res.Header.Get("Content-Type") != "application/json" {
		return m, errors.New(fmt.Sprintf("expected application/json response, got %s", res.Header.Get("Content-Type")))
	}

	rs := make([]map[string]string, 2)
	err = json.NewDecoder(res.Body).Decode(&rs)
	if err != nil {
		return m, err
	}

	if len(rs) < 1 {
		return m, errors.New("expected at least 1 element got 0")
	}

	for k, v := range rs[0] {
		if strings.HasSuffix(k, "_buy") || strings.HasSuffix(k, "_sell") {
			f, err := strconv.ParseFloat(v, 32)
			if err == nil {
				m[k] = float32(f)
			}
		}
	}

	return m, nil
}
