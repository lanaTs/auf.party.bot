package api

import (
	"fmt"
	"log/slog"

	"gopkg.in/telebot.v3"
)

type Handlers struct {
	logger *slog.Logger
}

func NewHandlers(logger *slog.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

// start
func (h *Handlers) StartHandler(c telebot.Context) error {
	user := c.Sender()
	chat := c.Chat()

	h.logger.Info("start command received",
		"user_id", user.ID,
		"username", user.Username,
		"chat_id", chat.ID,
		"chat_type", chat.Type)

	message := fmt.Sprintf("привет, %s", user.Username)

	return c.Send(message)
}
