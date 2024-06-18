package telegram_events

import "go-tg-bot/lib/wrappers"

type Meta struct {
	ChatID   int
	UserID   int
	UserName string
}

var (
	ErrUnknownEventType = wrappers.NewErr("unknown event type")
	ErrUnknownMetaType  = wrappers.NewErr("unknown meta type")
)
