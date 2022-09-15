package spotmargin

import (
	"net/http"
)

type RequestForSpotMarginInfo struct{}

type ResponseForSpotMarginInfo []SpotMarginInfo

type SpotMarginInfo struct {
	Coin          string  `json:"coin"`
	Borrowed      float64 `json:"borrowed"`
	Free          float64 `json:"free"`
	EstimatedRate float64 `json:"estimatedRate"`
	PreviousRate  float64 `json:"previousRate"`
}

func (req *RequestForSpotMarginInfo) Path() string {
	return "/spot_margin/market_info?market=BTC"
}

func (req *RequestForSpotMarginInfo) Method() string {
	return http.MethodGet
}

func (req *RequestForSpotMarginInfo) Query() string {
	return ""
}

func (req *RequestForSpotMarginInfo) Payload() []byte {
	return nil
}
