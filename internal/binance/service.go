package binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"strconv"
)

type ClientWrapper struct {
	api *binance.Client
}

func NewClientWrapper(api *binance.Client) *ClientWrapper {
	return &ClientWrapper{api: api}
}

func (c *ClientWrapper) GetBalance(ctx context.Context, asset string) (float64, error) {
	acc, err := c.api.NewGetAccountService().Do(ctx)
	if err != nil {
		return 0, err
	}

	for _, b := range acc.Balances {
		if b.Asset == asset {
			return strconv.ParseFloat(b.Free, 64)
		}
	}

	return 0, fmt.Errorf("asset %s not found", asset)
}

func (c *ClientWrapper) CreateMarketOrder(ctx context.Context, symbol string, quantity float64) error {
	_, err := c.api.NewCreateOrderService().
		Symbol(symbol).
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeMarket).
		Quantity(fmt.Sprintf("%.6f", quantity)).
		Do(ctx)

	return err
}
