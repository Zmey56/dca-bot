package scheduler

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Zmey56/dca-bot/internal/binance"
	"github.com/robfig/cron/v3"
)

func Start(client binance.BinanceClient) {
	symbol := os.Getenv("SYMBOL")
	amountStr := os.Getenv("BUY_AMOUNT")
	cronExpr := os.Getenv("SCHEDULE") // example: "0 10 * * *"

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Fatalf("❌ BUY_AMOUNT parsing error: %v", err)
	}

	c := cron.New()
	_, err = c.AddFunc(cronExpr, func() {
		log.Println("🕒 It's time to buy!")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := client.CreateMarketOrder(ctx, symbol, amount)
		if binance.HandleBinanceError(err) {
			err = client.CreateMarketOrder(ctx, symbol, amount)
		}
		if err != nil {
			log.Printf("❌ Buy error: %v", err)
			return
		}

		log.Printf("✅ Buy %.6f %s", amount, symbol)
	})
	if err != nil {
		log.Fatalf("Error when adding a cron task: %v", err)
	}

	log.Println("📅 The scheduler is running")
	c.Start()
	select {} //blocking main so that cron works
}
