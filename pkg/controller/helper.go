package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api-market-data/pkg/model"
)

func (cmdc *CurrentMarketDataController) getCurrentMarketData(vs_currency string) (model.CurrentMarketData, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&order=market_cap_desc&per_page=10", vs_currency)
	response, err := http.Get(url)
	if err != nil {
		return model.CurrentMarketData{}, err
	}
	defer response.Body.Close()
	var currentMarketData model.CurrentMarketData
	if err := json.NewDecoder(response.Body).Decode(&currentMarketData.CurrentMarketData); err != nil {
		return model.CurrentMarketData{}, err
	}
	return currentMarketData, nil
}

func (cmdc *CurrentMarketDataController) getCurrentCoinData(id string, currentMarketData model.CurrentMarketData) (*model.CurrentCoinData, error) {
	for _, coin := range currentMarketData.CurrentMarketData {
		if coin.ID == id {
			return &coin, nil
		}
	}
	return nil, fmt.Errorf("coin with id '%s' not found", id)
}

func (hcd *HistoricalCoinDataController) getHistoricalCoinData(vs_currency string, dateFrom float64, dateTo float64, id string) (model.HistoricalCoinData, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart/range?vs_currency=%s&from=%f&to=%f", id, vs_currency, dateFrom, dateTo)
	response, err := http.Get(url)
	if err != nil {
		return model.HistoricalCoinData{}, err
	}
	defer response.Body.Close()
	var historicalCoinData model.HistoricalCoinData
	if err := json.NewDecoder(response.Body).Decode(&historicalCoinData); err != nil {
		return model.HistoricalCoinData{}, err
	}
	return historicalCoinData, nil
}
