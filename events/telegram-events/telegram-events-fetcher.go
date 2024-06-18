package telegram_events

import (
	"context"

	tgClient "go-tg-bot/clients/telegram-client"
	"go-tg-bot/events"
	"go-tg-bot/lib/wrappers"
)

type TgEvFetcher struct {
	tg     *tgClient.Client
	offset int
}

func NewTgEvFetcher(client *tgClient.Client) *TgEvFetcher {
	return &TgEvFetcher{
		tg: client,
	}
}

func (tef *TgEvFetcher) Fetch(ctx context.Context, limit int) ([]events.Event, error) {
	updates, err := tef.tg.Updates(ctx, tef.offset, limit)
	if err != nil {
		return nil, wrappers.WrapErr("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	tef.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func event(upd tgClient.Update) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			UserID:   upd.Message.From.ID,
			UserName: upd.Message.From.UserName,
		}
	}

	return res
}

func fetchText(upd tgClient.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}

func fetchType(upd tgClient.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}

	return events.Message
}
