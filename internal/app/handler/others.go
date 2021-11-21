package handler

import (
	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/pcherednichenko/spam_fighter_bot/internal/app/data"
)

func Other(l *zap.SugaredLogger, b *tb.Bot, s data.Storage) func(m *tb.Message) {
	return func(m *tb.Message) {
		writeBotStatistic(l, m.Chat.ID, m.Chat.Title)
		if _, ok := s.Exist(m.Chat, m.Sender); !ok {
			return
		}
		err := b.Delete(m)
		if err != nil {
			l.Errorf("error while deleting spam message: %v", err)
		}
	}
}
