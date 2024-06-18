package main

import (
	"context"
	"log"

	tgClient "go-tg-bot/clients/telegram-client"
	eventConsumer "go-tg-bot/consumer/event-consumer"
	tgEvents "go-tg-bot/events/telegram-events"
	"go-tg-bot/lib/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	tgClient := tgClient.New(cfg.Telegram.Host, cfg.Telegram.BotToken)

	eventsFetcher := tgEvents.NewTgEvFetcher(tgClient)
	eventsProcessor := tgEvents.NewTgEvProcessor(tgClient)

	var batchSize int = 100
	consumer := eventConsumer.New(eventsFetcher, eventsProcessor, batchSize)

	log.Print("service started")

	if err := consumer.Start(ctx); err != nil {
		log.Fatal("service is stopped", err)
	}
}
