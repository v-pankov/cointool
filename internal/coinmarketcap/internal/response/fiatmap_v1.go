package response

type FiatMapV1 struct {
	Status StatusPayload `json:"status"`
	Data   FiatMapV1Data `json:"data"`
}

type FiatMapV1Data []FiatMapV1DataItem

type FiatMapV1DataItem struct {
	Symbol string `json:"symbol"`
}
