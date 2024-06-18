package telegram_events

import (
	"context"

	tgClient "go-tg-bot/clients/telegram-client"
	"go-tg-bot/events"
	"go-tg-bot/lib/wrappers"
)

type TgEvProcessor struct {
	tg *tgClient.Client
	// offset int
}

func NewTgEvProcessor(client *tgClient.Client) *TgEvProcessor {
	return &TgEvProcessor{
		tg: client,
	}
}

func (tep *TgEvProcessor) Process(ctx context.Context, event events.Event) error {
	switch event.Type {
	case events.Message:
		return tep.processMessage(ctx, event)
	default:
		return wrappers.WrapErr("can't process message", ErrUnknownEventType)
	}
}

func (tep *TgEvProcessor) processMessage(ctx context.Context, event events.Event) (err error) {
	defer func() { err = wrappers.WrapIfErr("can't process message", err) }()

	meta, err := meta(event)
	if err != nil {
		return err
	}

	if err := tep.doCmd(
		ctx,
		event.Text,
		meta.ChatID,
		meta.UserID,
		meta.UserName,
	); err != nil {
		return err
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, wrappers.WrapErr("can't get meta", ErrUnknownMetaType)
	}

	return res, nil
}
