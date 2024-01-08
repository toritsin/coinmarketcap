[![Go Reference](https://pkg.go.dev/badge/github.com/toritsin/coinmarketcap.svg)](https://pkg.go.dev/github.com/toritsin/coinmarketcap)

Client for CoinMarketCap API [Pro Version](https://pro.coinmarketcap.com/api/v1)

## Example Usage

```go
package main

import (
	"fmt"
	cmc "github.com/toritsin/coinmarketcap"
)

func main() {
	client := cmc.NewClient(cmc.Config{
		APIKey:              "your-secret-key",
		RequestTimeoutInSec: 10,
	})

	options := cmc.MapOptions{
		Start: 1,
		Limit: 10,
		Sort:  cmc.MapSortId,
	}
	m, err := client.GetMap(options)

	if err != nil {
		panic(err)
	}

	for _, v := range m.Data {
		fmt.Printf("%v", v)
	}
}
```