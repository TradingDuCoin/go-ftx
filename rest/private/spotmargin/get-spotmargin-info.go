package spotmargin

import (
	"fmt"
	"net/http"
)

type RequestForSpotMarginInfo struct {
	Market string
}

type ResponseForSpotMarginInfo []SpotMarginInfo

type SpotMarginInfo struct {
	Coin          string  `json:"coin"`
	Borrowed      float64 `json:"borrowed"`
	Free          float64 `json:"free"`
	EstimatedRate float64 `json:"estimatedRate"`
	PreviousRate  float64 `json:"previousRate"`
}

func (req *RequestForSpotMarginInfo) Path() string {
	return "/spot_margin/market_info"
}

func (req *RequestForSpotMarginInfo) Method() string {
	return http.MethodGet
}

func (req *RequestForSpotMarginInfo) Query() string {
	if req.Market != "" {
		return fmt.Sprintf("market=%s", req.Market)
	}
	return "market=BTC/USD"
}

func (req *RequestForSpotMarginInfo) Payload() []byte {
	return nil
}
