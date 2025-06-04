package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
)

func initBinanceClient() *binance.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	return binance.NewClient(apiKey, secretKey)
}

package main

import (
"context"
"log"
"os"
"strconv"
"time"

"github.com/joho/godotenv"

"github.com/Zmey56/dca-bot/internal/binance"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Не удалось загрузить .env: %v", err)
	}

	// Initializing the client
	client := binance.NewClientWrapper(binance.NewBinanceClient())

	// The timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Trying to get a balance
	balance, err := client.GetBalance(ctx, "USDT")
	if binance.HandleBinanceError(err) {
		// can be repeated if necessary.
		balance, err = client.GetBalance(ctx, "USDT")
	}
	if err != nil {
		log.Fatalf("Ошибка при получении баланса: %v", err)
	}

	log.Printf("💰 Баланс USDT: %.2f", balance)

	// Trying to send a market order
	amount := os.Getenv("BUY_AMOUNT") // Amount to buy in USDT
	amountF, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Fatalf("❌ Неверная сумма покупки в BUY_AMOUNT: %v", err)
	}

	err = client.CreateMarketOrder(ctx, "BTCUSDT", amountF)
	if binance.HandleBinanceError(err) {
		// can be repeated if necessary.
		err = client.CreateMarketOrder(ctx, "BTCUSDT", amountF)
	}
	if err != nil {
		log.Fatalf("❌ Ошибка при создании ордера: %v", err)
	}

	log.Println("✅ Рыночный ордер отправлен успешно!")
}
