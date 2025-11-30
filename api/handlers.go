package api

import (
	"fmt"
	"log/slog"
	"strings"

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

func (h *Handlers) CreateListHandler(c telebot.Context) error {

}

func (h *Handlers) HelpHandler(c telebot.Context) error {
	message := `🎮 *Управление ботом-органайзером* 🤪

*Просто нажми на кнопку ↓*`

	keyboard := &telebot.ReplyMarkup{}

	btnCreate := keyboard.Data("🎉 Создать список", "btn_create")
	btnShow := keyboard.Data("📜 Показать списки", "btn_lists")
	btnAdd := keyboard.Data("➕ Добавить пункт", "btn_add")

	btnDelete := keyboard.Data("🗑️ Удалить список", "btn_delete")
	btnHelp := keyboard.Data("❓ Помощь", "btn_help")

	keyboard.Inline(
		keyboard.Row(btnCreate),
		keyboard.Row(btnShow, btnAdd),
		keyboard.Row(btnDelete, btnHelp),
	)

	return c.Send(message, &telebot.SendOptions{
		ParseMode:   "Markdown",
		ReplyMarkup: keyboard,
	})
}

func (h *Handlers) CallbackHandler(c telebot.Context) error {
	callback := c.Callback()

	defer c.Respond()

	cleanData := strings.TrimSpace(callback.Data)
	cleanData = strings.ReplaceAll(cleanData, "\n", "")
	cleanData = strings.ReplaceAll(cleanData, "\u000c", "")
	cleanData = strings.TrimSpace(cleanData)

	h.logger.Info("callback processed",
		"original", fmt.Sprintf("%q", callback.Data),
		"cleaned", cleanData)

	switch cleanData {
	case "btn_create":
		return c.Send("🎯 *Создание нового списка!*\n\nКак назовем наш список безумия? 💥",
			&telebot.SendOptions{ParseMode: "Markdown"})
	case "btn_lists":
		return c.Send("📋 *Вот твои списки!*\n\nСкоро здесь появятся твои списки...",
			&telebot.SendOptions{ParseMode: "Markdown"})
	case "btn_add":
		return c.Send("➕ *Добавляем пункт!*\n\nКакой пункт добавляем в список? 📝",
			&telebot.SendOptions{ParseMode: "Markdown"})
	case "btn_delete":
		return c.Send("🗑️ *Удаляем список!*\n\nКакой список отправляем в небытие? 💀",
			&telebot.SendOptions{ParseMode: "Markdown"})
	case "btn_help":
		return h.HelpHandler(c)
	default:
		h.logger.Warn("unknown callback", "data", cleanData)
		return c.Send("❌ Неизвестная команда")
	}
}
