package coinmarketcap

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMap(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected method 'GET', but got %s", r.Method)
		}

		if r.URL.Path != "/v1/cryptocurrency/map" {
			t.Errorf("Expected path '/v1/cryptocurrency/map', but got %s", r.URL.Path)
		}

		if r.URL.RawQuery != "sort=id&start=1&limit=1&symbol=BNB" {
			t.Errorf("Expected params 'sort=id&start=1&limit=1&symbol=BNB', but got %s", r.URL.RawQuery)
		}

		if apiKey := r.Header.Get("X-CMC_PRO_API_KEY"); apiKey != "test_api_key" {
			t.Errorf("Expected API key 'test_api_key', but got %s", apiKey)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": [
				{
					"id": 1839,
					"rank": 3,
					"name": "Binance Coin",
					"symbol": "BNB",
					"slug": "binance-coin",
					"is_active": 1,
					"first_historical_data": "2017-07-25T04:30:05.000Z",
					"last_historical_data": "2020-05-05T20:44:02.000Z",
					"platform": {
						"id": 1027,
						"name": "Ethereum",
						"symbol": "ETH",
						"slug": "ethereum",
						"token_address": "0xB8c77482e45F1F44dE1745F52C74426C631bDD52"
					}
				}
			],
			"status": {
				"timestamp": "2018-06-02T22:51:28.209Z",
				"error_code": 0,
				"error_message": "",
				"elapsed": 10,
				"credit_count": 1
			}
		}`))
	}))

	defer server.Close()

	cfg := Config{
		APIKey:              "test_api_key",
		RequestTimeoutInSec: 5,
	}

	c := NewClient(cfg)

	c.httpClient = server.Client()
	c.apiBaseUrl = server.URL

	options := MapOptions{
		Sort:   MapSortId,
		Limit:  1,
		Symbol: "BNB",
		Start:  1,
	}

	resp, err := c.GetMap(options)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(resp.Data) != 1 {
		t.Errorf("Expected 1 result, but got %d", len(resp.Data))
	}

	assertBNBCoin(resp.Data[0], t)
}

func assertBNBCoin(data MapItem, t *testing.T) {
	if 1839 != data.Id {
		t.Errorf("Expected coin id equals 1839, but got %d", data.Id)
	}

	if 3 != data.Rank {
		t.Errorf("Expected coin rank equals 3, but got %d", data.Rank)
	}

	if "Binance Coin" != data.Name {
		t.Errorf("Expected coin name equals 'Binance Coin', but got %s", data.Name)
	}

	if "BNB" != data.Symbol {
		t.Errorf("Expected coin symbol equals 'BNB', but got %s", data.Symbol)
	}

	if "binance-coin" != data.Slug {
		t.Errorf("Expected coin slug equals 'binance-coin', but got %s", data.Slug)
	}

	if 1 != data.IsActive {
		t.Errorf("Expected coin is_active equals 1, but got %d", data.IsActive)
	}

	if "2017-07-25T04:30:05.000Z" != data.FirstHistoricalData {
		t.Errorf("Expected coin first_historical_data equals '2017-07-25T04:30:05.000Z', but got %s", data.FirstHistoricalData)
	}

	if "2020-05-05T20:44:02.000Z" != data.LastHistoricalData {
		t.Errorf("Expected coin last_historical_data equals '2020-05-05T20:44:02.000Z', but got %s", data.LastHistoricalData)
	}

	if 1027 != data.Platform.Id {
		t.Errorf("Expected coin paltform id equals 1027, but got %d", data.Platform.Id)
	}

	if "Ethereum" != data.Platform.Name {
		t.Errorf("Expected coin paltform name equals 'Ethereum', but got %s", data.Platform.Name)
	}

	if "ETH" != data.Platform.Symbol {
		t.Errorf("Expected coin paltform symbol equals 'ETH', but got %s", data.Platform.Symbol)
	}

	if "ethereum" != data.Platform.Slug {
		t.Errorf("Expected coin paltform slug equals 'ethereum', but got %s", data.Platform.Slug)
	}

	if "0xB8c77482e45F1F44dE1745F52C74426C631bDD52" != data.Platform.TokenAddress {
		t.Errorf("Expected coin paltform token_address equals '0xB8c77482e45F1F44dE1745F52C74426C631bDD52', but got %s", data.Platform.TokenAddress)
	}
}
