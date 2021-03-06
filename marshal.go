package coincross

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// String returns "BTC/USD" for {USD, BTC}.
func (p Pair) String() string {
	return string(p.Target + "/" + p.Base)
}

// Set sets the pair to {USD, BTC} from "BTC/USD".
func (p *Pair) Set(s string) error {
	parts := strings.Split(strings.ToUpper(s), "/")
	*p = Pair{Symbol(parts[1]), Symbol(parts[0])}
	return nil
}

// LowerString returns "btc_usd" for {USD, BTC}.
func (p Pair) LowerString() string {
	return strings.ToLower(string(p.Target + "_" + p.Base))
}

// MarshalJSON to "btc_usd" from {USD, BTC}.
func (p *Pair) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.LowerString())
}

// UnmarshalJSON from "btc_usd" to {USD, BTC}.
func (p *Pair) UnmarshalJSON(b []byte) (err error) {
	var s string
	err = json.Unmarshal(b, &s)
	if err == nil {
		parts := strings.Split(strings.ToUpper(s), "_")
		*p = Pair{Symbol(parts[1]), Symbol(parts[0])}
	}
	return
}

func (o Order) String() string {
	return fmt.Sprintf("%s\t%d\t%s\t%s\t%f\t%f(%f)", time.Unix(o.Timestamp, 0).Format("20060102 15:04:05"), o.Id, o.Type, o.Pair, o.Price, o.Remain, o.Amount)
}

func (t Trade) String() string {
	return fmt.Sprintf("%s %d\t%s\t%8.3f@%-8.6g\t!%s", t.Pair, t.Id, t.Type, t.Amount, t.Price, time.Unix(t.Timestamp, 0).Format("15:04:05"))
}

func (t Transaction) String() string {
	amounts := ""
	for k, v := range t.Amounts {
		amounts = amounts + fmt.Sprintf("\t%s:%f", k, v)
	}
	return fmt.Sprintf("%s\t%d%s\t%s", time.Unix(t.Timestamp, 0).Format("20060102 15:04:05"), t.Id, amounts, t.Descritpion)
}

func (t *TradeType) MarshalJSON() ([]byte, error) {
	var s string
	switch *t {
	case Buy:
		s = "buy"
	case Sell:
		s = "sell"
	}
	return json.Marshal(s)
}
func (t *TradeType) UnmarshalJSON(b []byte) (err error) {
	var s string
	err = json.Unmarshal(b, &s)
	if err == nil {
		switch strings.ToLower(s) {
		case "buy", "bid":
			*t = Buy
		case "sell", "ask":
			*t = Sell
		default:
			return fmt.Errorf("Unknown TradeType: %v", *t)
		}
	}
	return
}

func (t TradeType) String() string {
	switch t {
	case Sell:
		return "Sell"
	case Buy:
		return "Buy"
	default:
		return ""
	}
}

func (t *TradeType) Set(s string) error {
	return t.UnmarshalJSON([]byte(s))
}
