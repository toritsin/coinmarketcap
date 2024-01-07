package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	apiMapUrl = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map"
)

type Config struct {
	APIKey              string
	RequestTimeoutInSec time.Duration
}

type Client interface {
	GetMap(options MapOptions) (*MapResponse, error)
}

type client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(cfg Config) *client {
	var requestTimeout time.Duration

	if 0 == cfg.RequestTimeoutInSec {
		requestTimeout = time.Duration(60 * time.Second)
	} else {
		requestTimeout = time.Duration(cfg.RequestTimeoutInSec * time.Second)
	}

	return &client{
		httpClient: &http.Client{
			Timeout: requestTimeout,
		},
		apiKey: cfg.APIKey,
	}
}

func (c *client) GetMap(options MapOptions) (*MapResponse, error) {
	var params []string

	params = append(params, fmt.Sprintf("sort=%s", options.Sort.String()))

	if options.Start != 0 {
		params = append(params, fmt.Sprintf("start=%v", options.Start))
	}

	if options.Limit != 0 {
		params = append(params, fmt.Sprintf("limit=%v", options.Limit))
	}

	if options.Symbol != "" {
		params = append(params, fmt.Sprintf("symbol=%s", options.Symbol))
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", apiMapUrl, strings.Join(params, "&")), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := new(MapResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("JSON Error: [%s]. Response body: [%s]", err.Error(), string(body))
	}

	return resp, nil
}

func (c *client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("X-CMC_PRO_API_KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
