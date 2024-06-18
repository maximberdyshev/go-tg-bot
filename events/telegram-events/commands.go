package telegram_events

import (
	"context"
	"log"
	"strings"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (tep *TgEvProcessor) doCmd(ctx context.Context, text string, chatID, userID int, userName string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from user '%s(%d)' in chat '%d'", text, userName, userID, chatID)

	switch text {
	case HelpCmd:
		return tep.sendHelp(ctx, chatID)
	case StartCmd:
		return tep.sendHello(ctx, chatID)
	default:
		// return tep.tg.SendMessage(ctx, chatID, msgUnknownCommand)
		return tep.tg.SendMessage(ctx, chatID, text)
	}
}

func (tep *TgEvProcessor) sendHelp(ctx context.Context, chatID int) error {
	return tep.tg.SendMessage(ctx, chatID, msgHelp)
}

func (tep *TgEvProcessor) sendHello(ctx context.Context, chatID int) error {
	return tep.tg.SendMessage(ctx, chatID, msgHello)
}
