package exchange

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/koo04/fnar-go/fnar"
	"github.com/pkg/errors"
)

type Exchange struct {
	BuyingOrders        []BuyingOrder  `json:"BuyingOrders"`
	SellingOrders       []SellingOrder `json:"SellingOrders"`
	CXDataModelID       string         `json:"CXDataModelId"`
	MaterialName        string         `json:"MaterialName"`
	MaterialTicker      string         `json:"MaterialTicker"`
	MaterialID          string         `json:"MaterialId"`
	ExchangeName        string         `json:"ExchangeName"`
	ExchangeCode        string         `json:"ExchangeCode"`
	Currency            string         `json:"Currency"`
	Previous            interface{}    `json:"Previous"`
	Price               float32        `json:"Price"`
	PriceTimeEpochMs    int64          `json:"PriceTimeEpochMs"`
	High                float32        `json:"High"`
	AllTimeHigh         float32        `json:"AllTimeHigh"`
	Low                 float32        `json:"Low"`
	AllTimeLow          float32        `json:"AllTimeLow"`
	Ask                 float32        `json:"Ask"`
	AskCount            int            `json:"AskCount"`
	Bid                 float32        `json:"Bid"`
	BidCount            int            `json:"BidCount"`
	Supply              int            `json:"Supply"`
	Demand              int            `json:"Demand"`
	Traded              int            `json:"Traded"`
	VolumeAmount        float32        `json:"VolumeAmount"`
	PriceAverage        float32        `json:"PriceAverage"`
	NarrowPriceBandLow  float32        `json:"NarrowPriceBandLow"`
	NarrowPriceBandHigh float32        `json:"NarrowPriceBandHigh"`
	WidePriceBandLow    float32        `json:"WidePriceBandLow"`
	WidePriceBandHigh   float32        `json:"WidePriceBandHigh"`
	MMBuy               int            `json:"MMBuy"`
	MMSell              int            `json:"MMSell"`
	UserNameSubmitted   string         `json:"UserNameSubmitted"`
	Timestamp           time.Time      `json:"Timestamp"`
}

type exchangeCache struct {
	exchanges map[string]*Exchange

	mu sync.Mutex
}

const endpoint = "/exchange"

var cache = &exchangeCache{
	exchanges: map[string]*Exchange{},
}

func GetAll(ctx context.Context, full bool) ([]*Exchange, error) {
	exchanges := []*Exchange{}
	cache.mu.Lock()
	if len(cache.exchanges) != 0 {
		for _, ce := range cache.exchanges {
			exchanges = append(exchanges, ce)
		}
		cache.mu.Unlock()
		return exchanges, nil
	}
	cache.mu.Unlock()

	endpointExtra := "/all"
	if full {
		endpointExtra = "/full"
	}
	resp, err := http.Get(fnar.BaseUrl + endpoint + endpointExtra)
	if err != nil {
		return nil, errors.Wrap(err, "getting all exchanges")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 204 {
			return nil, fnar.Err_NOT_FOUND
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.Wrap(errors.New(string(b)), "not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading the response body")
	}

	sBody := string(body)
	_ = sBody

	m := []map[string]interface{}{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	for _, exchangeMap := range m {
		exchange := &Exchange{}
		if err := exchange.parse(exchangeMap); err != nil {
			return nil, err
		}
		exchanges = append(exchanges, exchange)

		cache.mu.Lock()
		cache.exchanges[exchange.MaterialTicker+"."+exchange.ExchangeCode] = exchange
		cache.mu.Unlock()
	}

	return exchanges, nil
}

func Get(ctx context.Context, exchangeCode string) (*Exchange, error) {
	cache.mu.Lock()
	if cachedExchange, ok := cache.exchanges[exchangeCode]; ok {
		cache.mu.Unlock()
		return cachedExchange, nil
	}
	cache.mu.Unlock()

	resp, err := http.Get(fnar.BaseUrl + endpoint + "/" + exchangeCode)
	if err != nil {
		return nil, errors.Wrap(err, "getting exchange")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return nil, fnar.Err_NOT_FOUND
	}

	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.Wrap(errors.New(string(b)), "not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading the response body")
	}

	sBody := string(body)
	_ = sBody

	m := map[string]interface{}{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	e := &Exchange{}
	if err := e.parse(m); err != nil {
		return nil, err
	}

	cache.mu.Lock()
	cache.exchanges[e.MaterialTicker+"."+e.ExchangeCode] = e
	cache.mu.Unlock()

	return e, nil
}

func (e *Exchange) parse(m map[string]interface{}) error {
	// correct timestamp to time.Time
	var timestamp time.Time
	var err error
	ts, ok := m["Timestamp"]
	if ok {
		timestamp, err = time.Parse("2006-01-02T15:04:05.999999", ts.(string))
		if err != nil {
			return err
		}
	}
	m["Timestamp"] = timestamp

	// parse the buying orders
	if buyingOrderInterfaces, ok := m["BuyingOrders"].([]interface{}); ok {
		buyingOrders := []*BuyingOrder{}
		for _, bo := range buyingOrderInterfaces {
			buyingOrder := &BuyingOrder{}
			if err := buyingOrder.parse(bo.(map[string]interface{})); err != nil {
				return err
			}

			buyingOrders = append(buyingOrders, buyingOrder)
		}
		m["BuyingOrders"] = buyingOrders
	}

	// parse the selling orders
	if sellingOrderInterfaces, ok := m["SellingOrders"].([]interface{}); ok {
		sellingOrders := []*SellingOrder{}
		for _, so := range sellingOrderInterfaces {
			sellingOrder := &SellingOrder{}
			if err := sellingOrder.Parse(so.(map[string]interface{})); err != nil {
				return err
			}

			sellingOrders = append(sellingOrders, sellingOrder)
		}
		m["SellingOrders"] = sellingOrders
	}

	// convert to Exchange
	em, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "marshaling exchange map[string]interface")
	}
	if err := json.Unmarshal(em, e); err != nil {
		return errors.Wrap(err, "unmarshaling to Exchange")
	}

	return nil
}
