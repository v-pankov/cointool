package response

type QuotesLatestV2 struct {
	Status StatusPayload      `json:"status"`
	Data   QuotesLatestV2Data `json:"data"`
}

type QuotesLatestV2Data map[string][]QuotesLatestV2DataItem

type QuotesLatestV2DataItem struct {
	Quote QuotesLatestV2DataQuote `json:"quote"`
}

type QuotesLatestV2DataQuote map[string]QuotesLatestV2DataQuoteItem

type QuotesLatestV2DataQuoteItem struct {
	Price float64 `json:"price"`
}
