package wallet

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForWithdrawHistories struct {
	StartTime time.Time `url:"start_time,omitempty"`
	EndTime   time.Time `url:"end_time,omitempty"`
}

type queryInput struct {
	StartTime int64 `url:"start_time,omitempty"`
	EndTime   int64 `url:"end_time,omitempty"`
}
type ResponseForWithdrawHistories []Withdraw

type Withdraw struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Status  string `json:"status"`
	Txid    string `json:"txid"`
	Method  string `json:"method"`

	Fee  float64 `json:"fee"`
	Size float64 `json:"size"`

	Time time.Time `json:"time"`

	ID int `json:"id"`
}

func (req *RequestForWithdrawHistories) Path() string {
	return "/wallet/withdrawals"
}

func (req *RequestForWithdrawHistories) Method() string {
	return http.MethodGet
}

func (req *RequestForWithdrawHistories) Query() string {
	values, _ := query.Values(queryInput{StartTime: req.StartTime.Unix(), EndTime: req.EndTime.Unix()})
	return values.Encode()
}

func (req *RequestForWithdrawHistories) Payload() []byte {
	return nil
}
