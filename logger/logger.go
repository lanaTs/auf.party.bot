package logger

import (
	"log"
	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {

	file, err := os.OpenFile("bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //todo убрать текст в переменные
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return slog.New(handler)
}
