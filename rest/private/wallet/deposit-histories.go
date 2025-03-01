package wallet

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForDepositHistories struct {
	StartTime time.Time `url:"start_time,omitempty"`
	EndTime   time.Time `url:"end_time,omitempty"`
}

type ResponseForDepositHistories []History

type History struct {
	Coin   string `json:"coin"`
	Status string `json:"status"`
	Txid   string `json:"txid"`

	Size float64 `json:"size"`
	Fee  float64 `json:"fee"`

	Confirmations int `json:"confirmations"`
	ID            int `json:"id"`

	ConfirmedTime time.Time `json:"confirmedTime"`
	SentTime      time.Time `json:"sentTime"`
	Time          time.Time `json:"time"`
}

func (req *RequestForDepositHistories) Path() string {
	return "/wallet/deposits"
}

func (req *RequestForDepositHistories) Method() string {
	return http.MethodGet
}

func (req *RequestForDepositHistories) Query() string {
	values, _ := query.Values(queryInput{StartTime: req.StartTime.Unix(), EndTime: req.EndTime.Unix()})
	return values.Encode()
}

func (req *RequestForDepositHistories) Payload() []byte {
	return nil
}
