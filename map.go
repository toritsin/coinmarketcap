package coinmarketcap

type MapSort int

const (
	MapSortId MapSort = iota
	MapSortCmcRank
)

func (s MapSort) String() string {
	return []string{"id", "cmc_rank"}[s]
}

type MapItem struct {
	Id                  uint     `json:"id"`
	Rank                uint     `json:"rank"`
	Name                string   `json:"name"`
	Symbol              string   `json:"symbol"`
	Slug                string   `json:"slug"`
	IsActive            int      `json:"is_active"`
	FirstHistoricalData string   `json:"first_historical_data"`
	LastHistoricalData  string   `json:"last_historical_data"`
	Platform            Platform `json:"platform"`
}

type MapResponse struct {
	Status Status    `json:"status"`
	Data   []MapItem `json:"data"`
}

type MapOptions struct {
	Start  int
	Limit  int
	Symbol string
	Sort   MapSort
}
