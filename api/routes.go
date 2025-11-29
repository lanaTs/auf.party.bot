package api

import "gopkg.in/telebot.v3"

type Route struct {
	Endpoint interface{}
	Handler  telebot.HandlerFunc
	Name     string
}

func (h *Handlers) GetRoutes() []Route {
	return []Route{
		{
			Endpoint: "/start",
			Handler:  h.StartHandler,
			Name:     "start",
		},
	}
}

func (h *Handlers) SetupHandlers(bot *telebot.Bot) {
	routes := h.GetRoutes()

	for _, route := range routes {
		bot.Handle(route.Endpoint, route.Handler)
		h.logger.Info("handler registered",
			"route", route.Name,
			"endpoint", route.Endpoint)
	}
}
